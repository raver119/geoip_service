package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/raver119/geoip_service/api"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGeoHandlers_LookupIp(t *testing.T) {
	e, err := NewLookupEngine("./test-data/test-data/GeoIP2-City-Test.mmdb")
	require.NoError(t, err)
	handlers := GeoHandlers{e: e}
	router := mux.NewRouter()
	handlers.Register(router)

	tests := []struct {
		name         string
		ip           string
		lang         string
		expectedCode int
		expectedCity string
	}{
		{"test_0", "81.2.69.142", "en", http.StatusOK, "London"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost/rest/v1/geo/%v/%v", tt.ip, tt.lang), bytes.NewReader([]byte{}))
			require.NoError(t, err)
			rw := httptest.NewRecorder()
			router.ServeHTTP(rw, r)

			require.Equal(t, tt.expectedCode, rw.Code)

			if rw.Code == http.StatusOK {
				var response api.LookupResponse
				_ = json.Unmarshal(rw.Body.Bytes(), &response)
				require.Equal(t, tt.expectedCity, response.City)
			}
		})
	}
}
