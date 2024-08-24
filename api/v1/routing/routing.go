package routing

import (
	"gotalk/api/v1/routing/handlers"
	"net/http"
)

func Router() *http.ServeMux{
	router := http.NewServeMux()
	
	router.HandleFunc("GET /ping", handlers.Pong)
	router.HandleFunc("POST /comment", handlers.PostComment)
	router.HandleFunc("POST /register", handlers.Register)
	
	return router
}

func AdminRouter() *http.ServeMux {
	adminRouter := http.NewServeMux()

	adminRouter.HandleFunc("POST /sudo/{user}", handlers.Sudo)
	adminRouter.HandleFunc("POST /thread/new", handlers.NewThread)
	adminRouter.HandleFunc("DELETE /thread/delete", handlers.DeleteThread)

	return adminRouter
}
