package server

import (
	"context"
	"log"
	"net/http"
	"website/server/common"
	"website/server/internal/http/api"
	"website/server/internal/http/middleware"

	"github.com/NYTimes/gziphandler"
	"github.com/urfave/negroni"
)

func initRoutesWithMiddleware(ctx context.Context) (n *negroni.Negroni) {

	n = negroni.New()

	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	n.Use(recovery)

	// use HTTP -> HTTPS redirect middleware
	n.Use(middleware.HTTPSRedirectMiddleware())

	// initialize the HTTP router
	mux := api.InitRouter(ctx)

	// middleware to gzip HTTP response
	h := gziphandler.GzipHandler(mux)

	n.UseHandler(h)

	return n
}

func Start() {
	var (
		err error
	)

	var serverPort = "8080"
	ctx := context.Background()
	handler := initRoutesWithMiddleware(ctx)

	var s *http.Server

	s = &http.Server{
		Addr:         ":" + serverPort,
		Handler:      handler,
		ReadTimeout:  common.ReadTimeout,
		WriteTimeout: common.WriteTimeout,
		IdleTimeout:  common.IdleTimeout,
	}

	err = s.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start the web server : %s", err.Error())
	}
}
