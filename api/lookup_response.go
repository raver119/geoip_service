package api

import "encoding/json"

type LookupResponse struct {
	ResponseLanguage string
	CountryCode      string
	Country          string
	City             string
	TimeZoneName     string
	Latitude         float64
	Longitude        float64
}

func (r LookupResponse) ToJson() (b []byte) {
	b, _ = json.Marshal(r)
	return
}

func (r LookupResponse) ToJsonString() string {
	return string(r.ToJson())
}
