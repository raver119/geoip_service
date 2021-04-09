package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

type GeoHandlers struct {
	e LookupEngine
}

func (h GeoHandlers) Register(r *mux.Router) {
	r.HandleFunc("/rest/v1/geo/{ip:[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}}/", h.LookupIp).Methods(http.MethodGet)
	r.HandleFunc("/rest/v1/geo/{ip:[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}}/{lang:{a-zA-Z}{2}}", h.LookupIp).Methods(http.MethodGet)
}

func (h GeoHandlers) LookupIp(w http.ResponseWriter, r *http.Request) {
	ip := mux.Vars(r)["ip"]
	lang := mux.Vars(r)["lang"]

	// take care of default language
	if len(lang) == 0 {
		lang = "en"
	}

	nip := net.ParseIP(ip)
	if nip == nil {
		http.Error(w, fmt.Sprintf("Bad IP: %v", ip), http.StatusBadRequest)
		return
	}

	response, err := h.e.LookupCity(nip, lang)
	if err != nil {
		http.Error(w, fmt.Sprintf("Lookup failed: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(response.ToJson())
}
