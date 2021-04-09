package main

import (
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestLookupEngine_LookupCity(t *testing.T) {
	e, err := NewLookupEngine("./test-data/test-data/GeoIP2-City-Test.mmdb")
	require.NoError(t, err)

	tests := []struct {
		name     string
		ip       net.IP
		wantCity string
		wantErr  bool
	}{
		{"test_0", net.ParseIP("81.2.69.142"), "London", false},
		{"test_1", net.ParseIP("alpha.beta.gamma."), "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := e.LookupCity(tt.ip, "en")
			if (err != nil) != tt.wantErr {
				t.Errorf("LookupCity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.Equal(t, tt.wantCity, got.City)
		})
	}
}
