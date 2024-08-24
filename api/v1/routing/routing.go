package routing

import (
	"gotalk/api/v1/middleware"
	"gotalk/api/v1/routing/handlers"
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()
	
	router.HandleFunc("GET /ping", handlers.Pong)
	router.HandleFunc("POST /register", handlers.Register)

	router.Handle("/admin/", middleware.EnsureAdmin(AdminRouter()))
	router.Handle("/", middleware.EnsureAuthenticated(AuthRouter()))

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	return v1
}

func AuthRouter() *http.ServeMux{
	router := http.NewServeMux()
	
	router.HandleFunc("POST /comment", handlers.PostComment)
	
	return router
}

func AdminRouter() *http.ServeMux {
	adminRouter := http.NewServeMux()

	adminRouter.HandleFunc("POST /sudo/{user}", handlers.Sudo)
	adminRouter.HandleFunc("POST /thread/new", handlers.NewThread)
	adminRouter.HandleFunc("DELETE /thread/delete", handlers.DeleteThread)
	
	admin := http.NewServeMux()
	admin.Handle("/admin/", http.StripPrefix("/admin", adminRouter))

	return admin
}
