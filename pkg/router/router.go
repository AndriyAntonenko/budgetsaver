package router

import (
	"fmt"
	"net/http"
)

type Router struct {
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)

	// TODO: Routing functionality
}
