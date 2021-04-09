## GeoIP service

This is a sample REST service built with https://github.com/oschwald/geoip2-golang

## How to use

Service expects two environment variables:

**GEOIP_URL**: URL that contains MaxMind GeoLite2 and GeoIP2 databases
**REST_PORT**: Port for HTTP requests.

Service exports 3 REST endpoints:

/rest/v1/geo/{ip:[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}}/
/rest/v1/geo/{ip:[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}}/{lang:[a-z]{2}}
/rest/v1/geo/health


