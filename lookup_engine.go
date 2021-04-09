package main

import (
	"github.com/oschwald/geoip2-golang"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

const FILE_NAME = "geodb.file"

type LookupEngine struct {
	db *geoip2.Reader
}

func NewLookupEngine() (LookupEngine, error) {
	_, e := os.Stat("/tmp/" + FILE_NAME)
	if e != nil {
		log.Printf("DB file is absent, downloading...")
		f, err := os.Create("/tmp/" + FILE_NAME)
		if err != nil {
			return LookupEngine{}, err
		}

		url, err := GetEnvOrError("GEOIP_URL")
		if err != nil {
			return LookupEngine{}, err
		}

		// file doesn't exist - download it
		resp, err := http.Get(url)
		if err != nil {
			return LookupEngine{}, err
		}

		defer resp.Body.Close()

		size, err := io.Copy(f, resp.Body)
		if err != nil {
			_ = f.Close()
			_ = os.Remove("/tmp/" + FILE_NAME)
			return LookupEngine{}, err
		}

		defer f.Close()

		log.Printf("DB file successfully downloaded. Got %v bytes.", size)
	}
	//
	db, err := geoip2.Open("/tmp/" + FILE_NAME)
	if err != nil {
		return LookupEngine{}, err
	}

	return LookupEngine{db: db}, nil
}

func (e LookupEngine) LookupCity(ip net.IP, lang string) (LookupResponse, error) {
	cr, err := e.db.City(ip)
	if err != nil {
		return LookupResponse{}, err
	}

	return LookupResponse{
		ResponseLanguage: lang,
		CountryCode:      cr.Country.IsoCode,
		Country:          cr.Country.Names[lang],
		City:             cr.City.Names[lang],
		TimeZoneName:     cr.Location.TimeZone,
		Latitude:         cr.Location.Latitude,
		Longitude:        cr.Location.Longitude,
	}, nil
}
