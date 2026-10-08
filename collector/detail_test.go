package collector

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/percona/mongodb_exporter/collector/serverstatusv5"
)

func TestDetailedRawCollectorDescribeAndRegistrationDoNotQueryMongoDB(t *testing.T) {
	client := new(mongo.Client)
	clientCalls := 0
	rawCalls := 0
	raw := detailedRaw(t, "mongod", 1)
	handler := newDetailedHandler(
		func() *mongo.Client {
			clientCalls++
			return client
		},
		func(*mongo.Client) (bson.Raw, error) {
			rawCalls++
			return raw, nil
		},
		false,
	)

	collector := &detailedRawCollector{module: serverstatusv5.NewDetailed(), raw: raw}
	descriptions := make(chan *prometheus.Desc)
	go func() {
		collector.Describe(descriptions)
		close(descriptions)
	}()
	descriptionCount := 0
	for range descriptions {
		descriptionCount++
	}
	if descriptionCount == 0 {
		t.Fatal("direct detailed collector Describe returned no descriptors")
	}

	registry := prometheus.NewPedanticRegistry()
	if err := registry.Register(collector); err != nil {
		t.Fatalf("register direct detailed collector: %v", err)
	}
	if clientCalls != 0 || rawCalls != 0 {
		t.Fatalf("Describe and registration made %d client calls and %d raw calls", clientCalls, rawCalls)
	}

	response := serveDetailedRequest(handler)
	if response.Code != http.StatusOK {
		t.Fatalf("detail status = %d, want %d: %s", response.Code, http.StatusOK, response.Body.String())
	}
	if clientCalls != 1 || rawCalls != 1 {
		t.Fatalf("one request made %d client calls and %d raw calls, want one each", clientCalls, rawCalls)
	}
}

func TestDetailedHandlerSuppressionRejectsBeforeMongoDBAccess(t *testing.T) {
	clientCalls := 0
	rawCalls := 0
	handler := newDetailedHandler(
		func() *mongo.Client {
			clientCalls++
			return new(mongo.Client)
		},
		func(*mongo.Client) (bson.Raw, error) {
			rawCalls++
			return nil, nil
		},
		true,
	)

	response := serveDetailedRequest(handler)
	if response.Code < 400 {
		t.Fatalf("suppressed detail status = %d, want non-2xx", response.Code)
	}
	if clientCalls != 0 || rawCalls != 0 {
		t.Fatalf("suppressed request made %d client calls and %d raw calls", clientCalls, rawCalls)
	}
	assertNoPrometheusMetricSamples(t, response.Body.String())
}

func TestDetailedHandlerQueryFailureDoesNotPublishStaleSamples(t *testing.T) {
	client := new(mongo.Client)
	calls := 0
	handler := newDetailedHandler(
		func() *mongo.Client { return client },
		func(*mongo.Client) (bson.Raw, error) {
			calls++
			if calls == 1 {
				return detailedRaw(t, "mongod", 17), nil
			}
			return nil, errors.New("serverStatus failed")
		},
		false,
	)

	first := serveDetailedRequest(handler)
	if first.Code != http.StatusOK {
		t.Fatalf("successful detail status = %d, want %d: %s", first.Code, http.StatusOK, first.Body.String())
	}
	if !strings.Contains(first.Body.String(), "mongodb_server_status_asserts_tripwire_total 17") {
		t.Fatalf("successful detail body lacks request value: %s", first.Body.String())
	}

	failed := serveDetailedRequest(handler)
	if failed.Code < 400 {
		t.Fatalf("failed detail status = %d, want non-2xx", failed.Code)
	}
	assertNoPrometheusMetricSamples(t, failed.Body.String())
	if strings.Contains(failed.Body.String(), " 17") {
		t.Fatalf("failed detail response published the preceding request value: %s", failed.Body.String())
	}
	if calls != 2 {
		t.Fatalf("two requests made %d raw calls, want 2", calls)
	}
}

func TestDetailedHandlerRejectsMongosFromRawResponse(t *testing.T) {
	rawCalls := 0
	handler := newDetailedHandler(
		func() *mongo.Client { return new(mongo.Client) },
		func(*mongo.Client) (bson.Raw, error) {
			rawCalls++
			return detailedRaw(t, "mongos", 23), nil
		},
		false,
	)

	response := serveDetailedRequest(handler)
	if response.Code < 400 {
		t.Fatalf("mongos detail status = %d, want non-2xx", response.Code)
	}
	if rawCalls != 1 {
		t.Fatalf("mongos detail made %d raw calls, want 1", rawCalls)
	}
	if !strings.Contains(response.Body.String(), "unavailable on mongos") {
		t.Fatalf("mongos response does not explain unavailability: %s", response.Body.String())
	}
	assertNoPrometheusMetricSamples(t, response.Body.String())
}

func TestDetailedHandlerEmitsOnlyDetailedModernMetrics(t *testing.T) {
	raw := mustMarshalDetailedRaw(t, bson.D{
		{Key: "process", Value: "mongod"},
		{Key: "asserts", Value: bson.D{{Key: "tripwire", Value: int64(31)}}},
		{Key: "wiredTiger", Value: bson.D{{
			Key:   "cache",
			Value: bson.D{{Key: "bytes allocated for updates", Value: int64(37)}},
		}}},
	})

	response := serveDetailedRequest(newDetailedMetricsHandler(serverstatusv5.NewDetailed(), raw))
	if response.Code != http.StatusOK {
		t.Fatalf("detail status = %d, want %d: %s", response.Code, http.StatusOK, response.Body.String())
	}
	body := response.Body.String()
	if !strings.Contains(body, "mongodb_server_status_asserts_tripwire_total 31") {
		t.Fatalf("detail body lacks detailed modern sample: %s", body)
	}
	for _, forbidden := range []string{
		"mongodb_server_status_wired_tiger_",
		"mongodb_exporter_",
		"mongodb_mongod_",
		"mongodb_version_info",
		"mongodb_up",
		"go_",
		"process_",
		"mongodb_exporter_build_info",
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("detail body contains forbidden family prefix %q: %s", forbidden, body)
		}
	}
}

func TestDetailedHandlerAccumulatesDiagnosticsAcrossRequestLocalRawResponses(t *testing.T) {
	client := new(mongo.Client)
	raws := []bson.Raw{
		mustMarshalDetailedRaw(t, bson.D{
			{Key: "process", Value: "mongod"},
			{Key: "asserts", Value: bson.D{{Key: "tripwire", Value: int64(101)}}},
			{Key: "catalogStats", Value: bson.D{{Key: "capped", Value: "not numeric"}}},
			{Key: "metrics", Value: bson.D{{Key: "commands", Value: bson.D{{Key: "unreviewed", Value: bson.D{{Key: "total", Value: int64(1)}}}}}}},
		}),
		mustMarshalDetailedRaw(t, bson.D{
			{Key: "process", Value: "mongod"},
			{Key: "asserts", Value: bson.D{{Key: "tripwire", Value: int64(202)}}},
			{Key: "catalogStats", Value: bson.D{{Key: "capped", Value: "still not numeric"}}},
			{Key: "metrics", Value: bson.D{{Key: "commands", Value: bson.D{{Key: "unreviewed", Value: bson.D{{Key: "total", Value: int64(2)}}}}}}},
		}),
	}
	call := 0
	handler := newDetailedHandler(
		func() *mongo.Client { return client },
		func(*mongo.Client) (bson.Raw, error) {
			raw := raws[call]
			call++
			return raw, nil
		},
		false,
	)

	first := serveDetailedRequest(handler)
	second := serveDetailedRequest(handler)
	if first.Code != http.StatusOK || second.Code != http.StatusOK {
		t.Fatalf("detail statuses = %d, %d; want 200, 200", first.Code, second.Code)
	}
	firstBody := first.Body.String()
	secondBody := second.Body.String()
	for metric, want := range map[string]string{
		`mongodb_server_status_decode_errors_total{family="mongodb_server_status_catalog_stats_capped"}`: "1",
		`mongodb_server_status_unknown_keys_total{family="commands"}`:                                    "1",
	} {
		if !strings.Contains(firstBody, metric+" "+want) {
			t.Fatalf("first detail body lacks %s %s: %s", metric, want, firstBody)
		}
	}
	for metric, want := range map[string]string{
		`mongodb_server_status_decode_errors_total{family="mongodb_server_status_catalog_stats_capped"}`: "2",
		`mongodb_server_status_unknown_keys_total{family="commands"}`:                                    "2",
	} {
		if !strings.Contains(secondBody, metric+" "+want) {
			t.Fatalf("second detail body lacks %s %s: %s", metric, want, secondBody)
		}
	}
	if !strings.Contains(firstBody, "mongodb_server_status_asserts_tripwire_total 101") || strings.Contains(firstBody, "mongodb_server_status_asserts_tripwire_total 202") {
		t.Fatalf("first detail response did not retain only its request-local raw value: %s", firstBody)
	}
	if !strings.Contains(secondBody, "mongodb_server_status_asserts_tripwire_total 202") || strings.Contains(secondBody, "mongodb_server_status_asserts_tripwire_total 101") {
		t.Fatalf("second detail response did not retain only its request-local raw value: %s", secondBody)
	}
}

func TestDetailedHandlerConcurrentRequestsKeepRawResponsesLocalWhileFastRawOperationOverlaps(t *testing.T) {
	client := new(mongo.Client)
	firstRaw := detailedRaw(t, "mongod", 101)
	secondRaw := detailedRaw(t, "mongod", 202)
	fastRaw := detailedRaw(t, "mongod", 303)
	firstEntered := make(chan struct{})
	secondEntered := make(chan struct{})
	fastEntered := make(chan struct{})
	releaseFirst := make(chan struct{})
	releaseSecond := make(chan struct{})
	releaseFast := make(chan struct{})
	var mu sync.Mutex
	detailRawCalls := 0
	fastRawCalls := 0

	handler := newDetailedHandler(
		func() *mongo.Client { return client },
		func(*mongo.Client) (bson.Raw, error) {
			mu.Lock()
			detailRawCalls++
			call := detailRawCalls
			mu.Unlock()
			switch call {
			case 1:
				close(firstEntered)
				<-releaseFirst
				return firstRaw, nil
			case 2:
				close(secondEntered)
				<-releaseSecond
				return secondRaw, nil
			default:
				return nil, errors.New("unexpected extra detailed serverStatus request")
			}
		},
		false,
	)
	fastRawGetter := func(*mongo.Client) (bson.Raw, error) {
		mu.Lock()
		fastRawCalls++
		mu.Unlock()
		close(fastEntered)
		<-releaseFast
		return fastRaw, nil
	}

	firstResult := make(chan *httptest.ResponseRecorder, 1)
	secondResult := make(chan *httptest.ResponseRecorder, 1)
	fastResult := make(chan error, 1)
	go func() { firstResult <- serveDetailedRequest(handler) }()
	<-firstEntered
	go func() {
		_, err := fastRawGetter(client)
		fastResult <- err
	}()
	<-fastEntered
	go func() { secondResult <- serveDetailedRequest(handler) }()
	<-secondEntered

	// The fast operation completes while both detail requests are still blocked
	// in their own raw getters, proving that all three operations overlapped.
	close(releaseFast)
	if err := <-fastResult; err != nil {
		t.Fatalf("simulated fast raw operation: %v", err)
	}
	close(releaseSecond)
	second := <-secondResult
	close(releaseFirst)
	first := <-firstResult

	if first.Code != http.StatusOK || second.Code != http.StatusOK {
		t.Fatalf("concurrent detail statuses = %d, %d; want 200, 200", first.Code, second.Code)
	}
	firstBody := first.Body.String()
	secondBody := second.Body.String()
	if !strings.Contains(firstBody, "mongodb_server_status_asserts_tripwire_total 101") || strings.Contains(firstBody, " 202") || strings.Contains(firstBody, " 303") {
		t.Fatalf("first detail response did not retain only its request-local raw value: %s", firstBody)
	}
	if !strings.Contains(secondBody, "mongodb_server_status_asserts_tripwire_total 202") || strings.Contains(secondBody, " 101") || strings.Contains(secondBody, " 303") {
		t.Fatalf("second detail response did not retain only its request-local raw value: %s", secondBody)
	}
	mu.Lock()
	defer mu.Unlock()
	if detailRawCalls != 2 {
		t.Fatalf("two concurrent detail requests made %d raw calls, want exactly one each", detailRawCalls)
	}
	if fastRawCalls != 1 {
		t.Fatalf("simulated fast operation made %d raw calls, want 1", fastRawCalls)
	}
}

func detailedRaw(t *testing.T, process string, tripwire int64) bson.Raw {
	t.Helper()
	return mustMarshalDetailedRaw(t, bson.D{
		{Key: "process", Value: process},
		{Key: "asserts", Value: bson.D{{Key: "tripwire", Value: tripwire}}},
	})
}

func mustMarshalDetailedRaw(t *testing.T, document interface{}) bson.Raw {
	t.Helper()
	data, err := bson.Marshal(document)
	if err != nil {
		t.Fatal(err)
	}
	return bson.Raw(data)
}

func serveDetailedRequest(handler http.Handler) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/metrics/detail", nil)
	handler.ServeHTTP(response, request)
	return response
}

func assertNoPrometheusMetricSamples(t *testing.T, body string) {
	t.Helper()
	for _, line := range strings.Split(body, "\n") {
		parser := expfmt.TextParser{}
		families, err := parser.TextToMetricFamilies(strings.NewReader(line + "\n"))
		if err != nil {
			continue
		}
		for name, family := range families {
			if len(family.GetMetric()) != 0 {
				t.Fatalf("response unexpectedly contains Prometheus sample %q: %s", name, body)
			}
		}
	}
}
