package traefik_plugin_replace_query_regex_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	traefik_plugin_replace_query_regex "github.com/NiklasPor/traefik-plugin-replace-query-regex"
)

func TestReplaceQueryRegex(t *testing.T) {
	cfg := &traefik_plugin_replace_query_regex.Config{
		Replacement: "BB",
		Regex:       "AA",
	}

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := traefik_plugin_replace_query_regex.New(ctx, next, cfg, "plugin-test")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost?testAAAwow=zZAAAZz&bbAAA=oAAo", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertQuery(t, req, "testBBAwow=zZBBAZz&bbBBA=oBBo")

}

func assertQuery(t *testing.T, req *http.Request, expected string) {
	t.Helper()

	if req.URL.RawQuery != expected {
		t.Errorf("invalid query value: %s", req.URL.RawQuery)
	}
}
