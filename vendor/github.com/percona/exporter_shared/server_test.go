// Copyright 2026 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exporter_shared

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestServerMuxRoutesFastAndDetail(t *testing.T) {
	restoreAuth := setServerAuth(t, "")
	defer restoreAuth()

	mux := newTestDetailMux(t, "/metrics", false)
	for _, test := range []struct {
		path string
		body string
	}{
		{path: "/metrics", body: "fast"},
		{path: "/metrics/detail", body: "detail"},
	} {
		recorder := serveRequest(mux, test.path, "", "")
		if recorder.Code != http.StatusOK {
			t.Fatalf("GET %s status = %d, want %d", test.path, recorder.Code, http.StatusOK)
		}
		if got := recorder.Body.String(); got != test.body {
			t.Fatalf("GET %s body = %q, want %q", test.path, got, test.body)
		}
	}
}

func TestServerMuxLandingPage(t *testing.T) {
	restoreAuth := setServerAuth(t, "")
	defer restoreAuth()

	recorder := serveRequest(newTestDetailMux(t, "/metrics", false), "/", "", "")
	if recorder.Code != http.StatusOK {
		t.Fatalf("landing page status = %d, want %d", recorder.Code, http.StatusOK)
	}
	body := recorder.Body.String()
	if !strings.Contains(body, "<h1>mongodb exporter</h1>") {
		t.Fatalf("landing page body %q does not contain exporter name", body)
	}
	if !strings.Contains(body, "<a href=\"/metrics\">Metrics</a>") {
		t.Fatalf("landing page body %q does not link to telemetry path", body)
	}
}

func TestServerMuxCustomTelemetryPath(t *testing.T) {
	restoreAuth := setServerAuth(t, "")
	defer restoreAuth()

	mux := newTestDetailMux(t, "/telemetry", false)
	for _, test := range []struct {
		path string
		body string
	}{
		{path: "/telemetry", body: "fast"},
		{path: "/telemetry/detail", body: "detail"},
	} {
		recorder := serveRequest(mux, test.path, "", "")
		if recorder.Code != http.StatusOK {
			t.Fatalf("GET %s status = %d, want %d", test.path, recorder.Code, http.StatusOK)
		}
		if got := recorder.Body.String(); got != test.body {
			t.Fatalf("GET %s body = %q, want %q", test.path, got, test.body)
		}
	}
}

func TestServerMuxWithoutDetailRetainsFastRoute(t *testing.T) {
	restoreAuth := setServerAuth(t, "")
	defer restoreAuth()

	paths, err := newServerPaths("/metrics", false)
	if err != nil {
		t.Fatal(err)
	}
	mux, err := newServerMux("mongodb", paths, testHandler("fast"), nil, false)
	if err != nil {
		t.Fatal(err)
	}

	recorder := serveRequest(mux, "/metrics", "", "")
	if got := recorder.Body.String(); got != "fast" {
		t.Fatalf("fast route body = %q, want %q", got, "fast")
	}
}

func TestServerPathsRejectInvalidCombinations(t *testing.T) {
	if _, err := newServerPaths("/", true); err == nil || !strings.Contains(err.Error(), "collides") {
		t.Fatalf("base path / error = %v, want landing collision", err)
	}

	for _, base := range []string{"metrics", "/metrics/", "/metrics/../detail"} {
		if _, err := newServerPaths(base, true); err == nil {
			t.Errorf("base path %q was accepted", base)
		}
	}
}

func TestServerMuxAuthenticatesBothMetricRoutes(t *testing.T) {
	restoreAuth := setServerAuth(t, "user:password")
	defer restoreAuth()

	mux := newTestDetailMux(t, "/metrics", false)
	for _, test := range []struct {
		path string
		body string
	}{
		{path: "/metrics", body: "fast"},
		{path: "/metrics/detail", body: "detail"},
	} {
		unauthorized := serveRequest(mux, test.path, "", "")
		if unauthorized.Code != http.StatusUnauthorized {
			t.Fatalf("unauthorized GET %s status = %d, want %d", test.path, unauthorized.Code, http.StatusUnauthorized)
		}

		authorized := serveRequest(mux, test.path, "user", "password")
		if authorized.Code != http.StatusOK {
			t.Fatalf("authorized GET %s status = %d, want %d", test.path, authorized.Code, http.StatusOK)
		}
		if got := authorized.Body.String(); got != test.body {
			t.Fatalf("authorized GET %s body = %q, want %q", test.path, got, test.body)
		}
	}
}

func TestHTTPSServerUsesSharedRoutesAndTLSBaseline(t *testing.T) {
	restoreAuth := setServerAuth(t, "")
	defer restoreAuth()

	server := newHTTPSServer("127.0.0.1:0", newTestDetailMux(t, "/metrics", true))
	if server.TLSConfig == nil {
		t.Fatal("TLS configuration is nil")
	}
	if got := server.TLSConfig.MinVersion; got != tls.VersionTLS12 {
		t.Fatalf("TLS minimum version = %d, want %d", got, tls.VersionTLS12)
	}
	if !server.TLSConfig.PreferServerCipherSuites {
		t.Fatal("TLS server cipher-suite preference is disabled")
	}
	if !hasCipherSuite(server.TLSConfig.CipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256) {
		t.Fatal("TLS baseline cipher suite is missing")
	}

	detail := serveRequest(server.Handler, "/metrics/detail", "", "")
	if got := detail.Body.String(); got != "detail" {
		t.Fatalf("HTTPS detail route body = %q, want %q", got, "detail")
	}
	landing := serveRequest(server.Handler, "/", "", "")
	if got := landing.Header().Get("Strict-Transport-Security"); got != strictTransportSecurity {
		t.Fatalf("HTTPS landing Strict-Transport-Security = %q, want %q", got, strictTransportSecurity)
	}
}

func TestTLSEnabledRejectsPartialConfiguration(t *testing.T) {
	for _, test := range []struct {
		name string
		cert string
		key  string
	}{
		{name: "certificate only", cert: "cert"},
		{name: "key only", key: "key"},
	} {
		t.Run(test.name, func(t *testing.T) {
			enabled, err := tlsEnabled(test.cert, test.key)
			if err == nil {
				t.Fatal("partial TLS configuration was accepted")
			}
			if enabled {
				t.Fatal("partial TLS configuration enabled HTTPS")
			}
		})
	}
}

func newTestDetailMux(t *testing.T, base string, ssl bool) *http.ServeMux {
	t.Helper()

	paths, err := newServerPaths(base, true)
	if err != nil {
		t.Fatal(err)
	}
	mux, err := newServerMux("mongodb", paths, testHandler("fast"), testHandler("detail"), ssl)
	if err != nil {
		t.Fatal(err)
	}
	return mux
}

func testHandler(body string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(body))
	})
}

func serveRequest(handler http.Handler, target, username, password string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, target, nil)
	if username != "" || password != "" {
		request.SetBasicAuth(username, password)
	}
	handler.ServeHTTP(recorder, request)
	return recorder
}

func setServerAuth(t *testing.T, auth string) func() {
	t.Helper()

	oldAuthFile := *authFileF
	oldHTTPAuth, hadHTTPAuth := os.LookupEnv("HTTP_AUTH")
	*authFileF = ""
	if auth == "" {
		if err := os.Unsetenv("HTTP_AUTH"); err != nil {
			t.Fatal(err)
		}
	} else if err := os.Setenv("HTTP_AUTH", auth); err != nil {
		t.Fatal(err)
	}

	return func() {
		*authFileF = oldAuthFile
		if hadHTTPAuth {
			if err := os.Setenv("HTTP_AUTH", oldHTTPAuth); err != nil {
				t.Error(err)
			}
			return
		}
		if err := os.Unsetenv("HTTP_AUTH"); err != nil {
			t.Error(err)
		}
	}
}

func hasCipherSuite(cipherSuites []uint16, expected uint16) bool {
	for _, cipherSuite := range cipherSuites {
		if cipherSuite == expected {
			return true
		}
	}
	return false
}
