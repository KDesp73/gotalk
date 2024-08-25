package routing

import (
	"gotalk/api/v1/middleware"
	"gotalk/api/v1/routing/handlers"
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()
	
	router.HandleFunc("/", handlers.ServeIndex)
	router.HandleFunc("GET /ping", handlers.Pong)
	router.HandleFunc("POST /register", handlers.Register)

	router.Handle("/admin/", http.StripPrefix("/admin", middleware.EnsureAdmin(AdminRouter())))
	router.Handle("/user/", http.StripPrefix("/user", middleware.EnsureAuthenticated(AuthRouter())))

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	return v1
}

func AuthRouter() *http.ServeMux {
	router := http.NewServeMux()
	
	router.HandleFunc("POST /comment/new", handlers.PostComment)
	
	return router
}

func AdminRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /sudo", handlers.Sudo)
	router.HandleFunc("POST /sudo/undo", handlers.UndoSudo)
	router.HandleFunc("POST /thread/new", handlers.NewThread)
	router.HandleFunc("DELETE /thread/delete", handlers.DeleteThread)
	router.HandleFunc("DELETE /comment/delete", handlers.DeleteComment)
	
	return router
}
