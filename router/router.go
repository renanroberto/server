package router

import (
	"log"
	"net/http"
	"time"
)

type Route struct {
	Name    string
	Path    string
	Methods []string
	Handler http.HandlerFunc
	Log     bool
}

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	for _, route := range routes {
		router.HandleFunc(route.Init())
	}

	return router
}

func (route Route) Init() (path string, handler http.HandlerFunc) {
	path = route.Path

	if route.Log {
		handler = route.Logger(route.Handler)
	} else {
		handler = route.Handler
	}

	handler = route.MethodAllowed(handler)

	return
}

func (route Route) MethodAllowed(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := r.Method

		hasMethod := false
		for _, value := range route.Methods {
			if method == value {
				hasMethod = true
				break
			}
		}

		if hasMethod {
			handler.ServeHTTP(w, r)
		} else {
			http.Error(w, "Method not allowed", 405)
		}
	}
}

func (route Route) Logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		duration := time.Since(start)

		log.Printf("%s \t %s \t %s \t %s\n", route.Name, r.Method, r.RequestURI, duration)
	}
}
