package api_test

import (
	"encoding/json"
	"godemo/api"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	r := api.NewRouter(&json.RawMessage{})
	req := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	resp, err := io.ReadAll(rr.Result().Body)
	require.NoError(t, err)
	require.JSONEq(t, `{"msg":"hello"}`, string(resp))
}
