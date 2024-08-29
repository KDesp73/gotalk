package routing

import (
	"gotalk/api/v1/middleware"
	"gotalk/api/v1/routing/handlers"
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()
	auth := http.NewServeMux()
	admin := http.NewServeMux()
	v1 := http.NewServeMux()
	
	router.HandleFunc("GET /ping", handlers.Pong)
	router.HandleFunc("POST /users/new", handlers.Register)

	auth.Handle("/", middleware.EnsureAuthenticated(AuthRouter()))
	admin.Handle("/", middleware.EnsureAdmin(AdminRouter()))

	v1.Handle("/v1/", http.StripPrefix("/v1", router))
	v1.Handle("/v1/auth/", http.StripPrefix("/v1/auth", auth))
	v1.Handle("/v1/admin/", http.StripPrefix("/v1/admin", admin))
	v1.HandleFunc("/", handlers.ServeIndex)

	return v1
}

// Routes that need at least a registration
func AuthRouter() *http.ServeMux {
	router := http.NewServeMux()
	
	router.HandleFunc("DELETE /comments/{commentid}", handlers.DeleteComment) // ?threadid={threadid}
	router.HandleFunc("POST /users/{userid}/comment", handlers.PostComment) // ?threadid={threadid}&content={content}
	
	return router
}

// Routes that need administator privileges
func AdminRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /comments", handlers.GetComments)

	router.HandleFunc("DELETE /threads/{threadid}", handlers.DeleteThread)
	router.HandleFunc("GET /threads", handlers.GetThreads)
	router.HandleFunc("POST /threads/new", handlers.NewThread) // ?title={title}

	router.HandleFunc("PUT /users/{userid}/sudo", handlers.Sudo)
	router.HandleFunc("PUT /users/{userid}/sudo/revoke", handlers.UndoSudo)
	
	return router
}
