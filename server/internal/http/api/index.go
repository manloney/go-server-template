package api

import (
	"context"
	"net/http"
	"strings"
	"website/server/common"

	"github.com/gorilla/mux"
)

func registerIndexRoutes(ctx context.Context, router *mux.Router, serveMux *http.ServeMux) {
	router.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("./ui/dist"))))

	// function to serve the front-end
	indexFunc := func(w http.ResponseWriter, r *http.Request) {
		// handle www as a subdomain
		// redirect www to the application endpoint
		vars := mux.Vars(r)
		subdomain := vars["subdomain"]

		if strings.EqualFold(subdomain, "www") {
			http.Redirect(w, r, common.GetApplicationEndPoint(), http.StatusPermanentRedirect)
			return
		}

		w.Header().Set("cache-control", "no-cache")
		http.ServeFile(w, r, "./ui/dist/index.html")
	}

	router.HandleFunc("/", indexFunc).Methods("GET")
}
