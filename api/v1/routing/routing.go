package routing

import (
	"gotalk/internal/routing/handlers"
	"net/http"
)

func Router() *http.ServeMux{
	router := http.NewServeMux()
	
	router.HandleFunc("GET /greet/{name}", handlers.Greeter)
	router.HandleFunc("/ping", handlers.Pong)
	
	return router
}

func AdminRouter() *http.ServeMux {
	adminRouter := http.NewServeMux()

	adminRouter.HandleFunc("POST /sudo/{user}", handlers.Sudo)

	return adminRouter
}
