package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jcgarciaram/demoPark/demoParkAPI"
	r "github.com/jcgarciaram/demoPark/demoParkAPI/routes"
)

func NewRouter() *mux.Router {

	routes := r.Routes{}

	// Append demoParkAPI routes
	routes.AppendRoutes(demoParkAPI.GetRoutes())

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		// Last thing that will be executed is the actual handler function
		handler = route.HandlerFunc

		// TODO: Build VerifyPerms. Idea: use routes struct and expand with modules and submodules
		// if route.VerifyPerms {
		// 	// Before that we will verify permissions
		// 	handler = VerifyPermissions(handler)
		// }

		if route.VerifyJWT {
			// Before THAT we will validate the token passed
			handler = jwtMiddleware.Handler(handler)
		}

		// And even before THAT we will start the logger to be able to calculate how everything takes.
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}

type MyServer struct {
	r *mux.Router
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, HEAD")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.Header().Set("Access-Control-Max-Age", "86400")
		// rw.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-XSRF-Token, X-HTTP-Method-Override, X-Requested-With, Mobile-Cookie")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Let Gorilla work
	s.r.ServeHTTP(rw, req)
}
