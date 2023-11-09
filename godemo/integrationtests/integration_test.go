//go:build integration

package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const baseUrl = "http://127.0.0.1:8080"

func TestApp(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Readiness",
			path:         "/ready",
			expectedCode: 200,
			expectedBody: `"ready"`,
		},
		{
			name:         "Hello API",
			path:         "/api/hello",
			expectedCode: 200,
			expectedBody: `{"msg":"hello"}`,
		},
		{
			name:         "Doc",
			path:         "/api/doc",
			expectedCode: 200,
			expectedBody: mustReadFile(t, "../docs/swagger.json"),
		},
		{
			name:         "Not existing endpoint",
			path:         "/api/not-here",
			expectedCode: 404,
			expectedBody: `"page not found"`,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(baseUrl + tc.path)
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tc.expectedCode, resp.StatusCode)
			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			assert.JSONEq(t, tc.expectedBody, string(body))
		})
	}
}

func mustReadFile(t testing.TB, filename string) string {
	data, err := ioutil.ReadFile(filename)
	require.NoError(t, err)
	return string(data)
}
