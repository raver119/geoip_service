package main

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"github.com/raver119/geoip_service/api"
	"net"
)

type LookupEngine struct {
	db *geoip2.Reader
}

func NewLookupEngine(fileName string) (LookupEngine, error) {
	db, err := geoip2.Open(fileName)
	if err != nil {
		// if file has bad format - report it, and delete, so it can be re-downloaded later again
		return LookupEngine{}, err
	}

	return LookupEngine{db: db}, nil
}

func (e LookupEngine) LookupCity(ip net.IP, lang string) (api.LookupResponse, error) {
	if ip == nil {
		return api.LookupResponse{}, fmt.Errorf("nil isn't a valid IP")
	}

	cr, err := e.db.City(ip)
	if err != nil {
		return api.LookupResponse{}, err
	}

	return api.LookupResponse{
		ResponseLanguage: lang,
		CountryCode:      cr.Country.IsoCode,
		Country:          cr.Country.Names[lang],
		City:             cr.City.Names[lang],
		TimeZoneName:     cr.Location.TimeZone,
		Latitude:         cr.Location.Latitude,
		Longitude:        cr.Location.Longitude,
	}, nil
}
