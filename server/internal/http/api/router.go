package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRouter initializes the API routes for all the components using a subrouter to support subdomain
// The subrouter logic doesnot seem to work with newer versions of Gorilla Mux after v1.6.2
// An issue opened for this bug was not fixed by the developer: https://github.com/gorilla/mux/issues/522
func InitRouter(ctx context.Context) (serveMux *http.ServeMux) {

	r := mux.NewRouter().StrictSlash(true)

	// dynamic subdomain support
	// subdomain name is later available as a variable in the Gorilla mux as follows, just like any other route variable:
	//	vars := mux.Vars(r)
	//	subdomain := vars["subdomain"]
	// Change vibhormeshram to a variable
	s := r.Host("{subdomain:[a-zA-Z0-9\\-]*}" + "." + "vibhormeshram").Subrouter()

	serveMux = http.NewServeMux()

	serveMux.Handle("/", s)

	registerV1Routes(ctx, s, serveMux)

	return serveMux
}

// registerV1Routes registers all the version 1 API routes
func registerV1Routes(ctx context.Context, router *mux.Router, serveMux *http.ServeMux) {

	// registers the index and static asset routes
	registerIndexRoutes(ctx, router, serveMux)
}
