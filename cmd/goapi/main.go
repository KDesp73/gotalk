package main

import (
	"flag"
	"fmt"
	"goapi/internal/middleware"
	"goapi/internal/routing"
	"log"
	"net/http"
)

func main() {
	var port = 8080
	flag.IntVar(&port, "port", port, "Specify the port")
	flag.Parse()

	router := routing.Router()
	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.IsAuthenticated,
		middleware.LoadUser,
	)

	server := http.Server {
		Addr: fmt.Sprintf(":%d", port),
		Handler: stack(router),
	}

	log.Printf("Starting server on port %d", port)
	server.ListenAndServe()

}
