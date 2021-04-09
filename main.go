package main

import (
	"log"
	"strconv"
)

func main() {
	fileName := "/tmp/" + FILE_NAME
	url, err := GetEnvOrError("GEOIP_URL")
	if err != nil {
		log.Printf("GEOIP_URL must be defined")
		return
	}

	err = DownloadFile(url, fileName)
	if err != nil {
		log.Printf("Unable to download GeoIP file: %v", err)
		return
	}

	port, err := strconv.Atoi(GetEnvOrDefault("REST_PORT", "8080"))
	if err != nil {
		panic(err)
	}
	srv := RestServer(port, fileName)

	err = srv.Start()
	if err != nil {
		panic(err)
	}
}
