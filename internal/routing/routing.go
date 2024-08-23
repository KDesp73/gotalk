package routing

import "net/http"

func Router() *http.ServeMux{
	router := http.NewServeMux()
	
	router.HandleFunc("GET /greet/{name}", Greeter)
	router.HandleFunc("/ping", Pong)
	
	return router
}

func AdminRouter() *http.ServeMux {
	adminRouter := http.NewServeMux()

	adminRouter.HandleFunc("GET /comments", )
}
