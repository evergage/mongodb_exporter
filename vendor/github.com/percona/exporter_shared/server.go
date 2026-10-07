// Copyright 2017 Percona LLC
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

// Package exporter_shared provides shared code for Percona Prometheus exporters.
package exporter_shared

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	sslCertFileF = kingpin.Flag("web.ssl-cert-file", "Path to SSL certificate file.").String()
	sslKeyFileF  = kingpin.Flag("web.ssl-key-file", "Path to SSL key file.").String()

	landingPage = template.Must(template.New("home").Parse(strings.TrimSpace(`
<html>
<head>
	<title>{{ .name }} exporter</title>
</head>
<body>
	<h1>{{ .name }} exporter</h1>
	<p><a href="{{ .path }}">Metrics</a></p>
</body>
</html>
`)))
)

// DefaultMetricsHandler returns metrics handler for default Prometheus gatherer/registerer
// with logging and continuing on error.
// Handler is not protected by HTTP basic authentication - it is done by RunServer.
func DefaultMetricsHandler() http.Handler {
	h := promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{
		ErrorLog:      log.NewErrorLogger(),
		ErrorHandling: promhttp.ContinueOnError,
	})
	return promhttp.InstrumentMetricHandler(prometheus.DefaultRegisterer, h)
}

const (
	landingPath             = "/"
	detailPathSuffix        = "/detail"
	strictTransportSecurity = "max-age=63072000; includeSubDomains"
)

type serverPaths struct {
	telemetry string
	detail    string
	landing   string
}

// RunServer runs a server for an exporter with one telemetry route. It retains
// the existing landing page, HTTP basic-authentication, and TLS behavior.
// Function never returns.
func RunServer(name, addr, path string, handler http.Handler) {
	runServer(name, addr, path, handler, nil, false)
}

// RunServerWithDetail runs a server for an exporter with fast metrics at base
// and detailed metrics at base + "/detail". Both metric routes use the same
// listener, TLS configuration, and HTTP basic-authentication policy.
// Function never returns.
func RunServerWithDetail(name, addr, base string, fast, detail http.Handler) {
	runServer(name, addr, base, fast, detail, true)
}

func runServer(name, addr, base string, fast, detail http.Handler, withDetail bool) {
	paths, err := newServerPaths(base, withDetail)
	if err != nil {
		log.Fatal(err)
	}

	ssl, err := tlsEnabled(*sslCertFileF, *sslKeyFileF)
	if err != nil {
		log.Fatal(err)
	}

	mux, err := newServerMux(name, paths, fast, detail, ssl)
	if err != nil {
		log.Fatal(err)
	}

	if ssl {
		runHTTPS(addr, paths.telemetry, mux)
	} else {
		runHTTP(addr, paths.telemetry, mux)
	}
}

func newServerPaths(base string, withDetail bool) (serverPaths, error) {
	paths := serverPaths{
		telemetry: base,
		landing:   landingPath,
	}
	if withDetail {
		paths.detail = base + detailPathSuffix
	}

	if err := paths.validate(); err != nil {
		return serverPaths{}, err
	}
	return paths, nil
}

func (paths serverPaths) validate() error {
	routes := []struct {
		name string
		path string
	}{
		{name: "telemetry", path: paths.telemetry},
	}
	if paths.detail != "" {
		routes = append(routes, struct {
			name string
			path string
		}{name: "detail", path: paths.detail})
	}
	routes = append(routes, struct {
		name string
		path string
	}{name: "landing", path: paths.landing})

	for i, route := range routes {
		for _, other := range routes[:i] {
			if route.path == other.path {
				return fmt.Errorf("%s path %q collides with %s path", route.name, route.path, other.name)
			}
		}
	}

	for _, route := range routes {
		if err := validateRoutePath(route.path); err != nil {
			return fmt.Errorf("invalid %s path %q: %v", route.name, route.path, err)
		}
	}
	return nil
}

func validateRoutePath(routePath string) error {
	if routePath == "" {
		return fmt.Errorf("must not be empty")
	}
	if !strings.HasPrefix(routePath, "/") {
		return fmt.Errorf("must begin with /")
	}
	if strings.ContainsAny(routePath, "?#") {
		return fmt.Errorf("must not contain a query or fragment")
	}
	if routePath != path.Clean(routePath) {
		return fmt.Errorf("must be a canonical path")
	}
	return nil
}

func newServerMux(name string, paths serverPaths, fast, detail http.Handler, ssl bool) (*http.ServeMux, error) {
	if fast == nil {
		return nil, fmt.Errorf("fast handler must not be nil")
	}
	if paths.detail != "" && detail == nil {
		return nil, fmt.Errorf("detail handler must not be nil")
	}

	var buf bytes.Buffer
	data := map[string]string{"name": name, "path": paths.telemetry}
	if err := landingPage.Execute(&buf, data); err != nil {
		return nil, err
	}
	landing := buf.Bytes()

	mux := http.NewServeMux()
	mux.Handle(paths.telemetry, authHandler(fast))
	if paths.detail != "" {
		mux.Handle(paths.detail, authHandler(detail))
	}
	mux.HandleFunc(paths.landing, func(w http.ResponseWriter, _ *http.Request) {
		if ssl {
			w.Header().Add("Strict-Transport-Security", strictTransportSecurity)
		}
		_, _ = w.Write(landing)
	})
	return mux, nil
}

func tlsEnabled(certFile, keyFile string) (bool, error) {
	if (certFile == "") != (keyFile == "") {
		return false, fmt.Errorf("one of the flags --web.ssl-cert-file or --web.ssl-key-file is missing to enable HTTPS")
	}
	if certFile == "" {
		return false, nil
	}
	if _, err := os.Stat(certFile); err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("SSL certificate file does not exist: %s", certFile)
		}
		return false, fmt.Errorf("cannot stat SSL certificate file %q: %v", certFile, err)
	}
	if _, err := os.Stat(keyFile); err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("SSL key file does not exist: %s", keyFile)
		}
		return false, fmt.Errorf("cannot stat SSL key file %q: %v", keyFile, err)
	}
	return true, nil
}

func runHTTPS(addr, path string, handler http.Handler) {
	srv := newHTTPSServer(addr, handler)
	log.Infof("Starting HTTPS server for https://%s%s ...", addr, path)
	log.Fatal(srv.ListenAndServeTLS(*sslCertFileF, *sslKeyFileF))
}

func newHTTPSServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:      addr,
		Handler:   handler,
		TLSConfig: newTLSConfig(),
	}
}

func newTLSConfig() *tls.Config {
	// See internal security baseline.
	return &tls.Config{
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			// No SHA-1, ECDHE before plain RSA, GCM before CBC.
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
		},
	}
}

func runHTTP(addr, path string, handler http.Handler) {
	srv := newHTTPServer(addr, handler)
	log.Infof("Starting HTTP server for http://%s%s ...", addr, path)
	log.Fatal(srv.ListenAndServe())
}

func newHTTPServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
