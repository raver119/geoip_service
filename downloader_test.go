package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	url := "https://github.com/maxmind/MaxMind-DB/raw/5a0be1c0320490b8e4379dbd5295a18a648ff156/test-data/GeoIP2-City-Test.mmdb"
	fileName := "/tmp/randomfile_geoip_test"
	require.NoError(t, DownloadFile(url, fileName))
	require.NoError(t, os.Remove(fileName))
}
