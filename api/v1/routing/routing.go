package routing

import (
	"gotalk/api/v1/middleware"
	"gotalk/api/v1/routing/handlers"
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()
	v1 := http.NewServeMux()
	
	router.HandleFunc("GET /ping", handlers.Pong)
	router.HandleFunc("POST /users/new", handlers.Register) // ?name={name}&email={email}

	router.Handle("/auth/", http.StripPrefix("/auth", middleware.EnsureAuthenticated(AuthRouter())))
	router.Handle("/admin/", http.StripPrefix("/admin", middleware.EnsureAdmin(AdminRouter())))

	v1.Handle("/v1/", http.StripPrefix("/v1", router))
	v1.HandleFunc("/", handlers.ServeIndex)
	v1.HandleFunc("/dark", handlers.ServeIndex)

	return v1
}

// Routes that need at least a registration
func AuthRouter() *http.ServeMux {
	router := http.NewServeMux()
	
	router.HandleFunc("DELETE /comments/{commentid}", handlers.DeleteComment) // ?threadid={threadid}
	router.HandleFunc("POST /users/{userid}/comment", handlers.PostComment) // ?threadid={threadid}&content={content}
	router.HandleFunc("DELETE /users/{userid}", handlers.DeleteUser)
	
	return router
}

// Routes that need administator privileges
func AdminRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /comments", handlers.GetComments) // ?threadid={threadid}

	router.HandleFunc("DELETE /threads/{threadid}", handlers.DeleteThread)
	router.HandleFunc("GET /threads", handlers.GetThreads)
	router.HandleFunc("POST /threads/new", handlers.NewThread) // ?title={title}

	router.HandleFunc("GET /users", handlers.GetUsers)
	router.HandleFunc("PUT /users/{userid}/sudo", handlers.Sudo)
	router.HandleFunc("PUT /users/{userid}/sudo/revoke", handlers.UndoSudo)
	
	return router
}
