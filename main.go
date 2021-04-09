package main

import "strconv"

func main() {
	port, err := strconv.Atoi(GetEnvOrDefault("REST_PORT", "8080"))
	if err != nil {
		panic(err)
	}
	srv := RestServer(port)

	err = srv.Start()
	if err != nil {
		panic(err)
	}
}
