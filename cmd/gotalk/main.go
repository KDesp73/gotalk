package main

import (
	"flag"
	"fmt"
	"gotalk/api/state"
	"gotalk/api/v1/middleware"
	"gotalk/api/v1/routing"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)


	var port = 8080
	flag.IntVar(&port, "port", port, "Specify the port")
	flag.Parse()
	
	// TODO: load state
	state.Instance = state.StateInit()

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
	go server.ListenAndServe()

	<-sigChan
	fmt.Println("\nTerminating...")

	for _, user := range state.Instance.Users.Users {
		user.Log()
		println()
	}

	// TODO: save state
}
