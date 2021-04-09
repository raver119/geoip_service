package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type ServerEngine struct {
	port int

	e LookupEngine
	h GeoHandlers

	startedAsync bool
	r            *mux.Router
	srv          *http.Server
	wg           *sync.WaitGroup
}

func RestServer(port int) ServerEngine {
	e, err := NewLookupEngine()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	h := GeoHandlers{e: e}
	h.Register(router)

	return ServerEngine{
		port: port,
		e:    e,
		h:    h,
		srv:  &http.Server{Addr: ":" + strconv.Itoa(port), Handler: router},
		wg:   &sync.WaitGroup{},
	}
}

func (e ServerEngine) Start() (err error) {
	log.Printf("Starting REST server on port %v", e.port)
	return e.srv.ListenAndServe()
}

func (e ServerEngine) StartAsync() (err error) {
	e.startedAsync = true
	e.wg.Add(1)

	go func() {
		defer e.wg.Done()

		_ = e.srv.ListenAndServe()
	}()
	return
}

func (e ServerEngine) Stop() (err error) {
	if e.startedAsync {
		_ = e.srv.Shutdown(context.TODO())
		// do it
		e.wg.Wait()
	}
	return
}
