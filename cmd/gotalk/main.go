package main

import (
	"flag"
	"fmt"
	"gotalk/api/v1/middleware"
	"gotalk/api/v1/routing"
	"log"
	"net/http"
)

func main() {
	var port = 8080
	flag.IntVar(&port, "port", port, "Specify the port")
	flag.Parse()

	router := routing.Router()
	adminRouter := routing.AdminRouter()

	router.Handle("/", middleware.EnsureAdmin(adminRouter))

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.IsAuthenticated,
	)

	server := http.Server {
		Addr: fmt.Sprintf(":%d", port),
		Handler: stack(router),
	}

	log.Printf("Starting server on port %d", port)
	server.ListenAndServe()

}
