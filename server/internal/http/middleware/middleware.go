package middleware

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

// HTTPSRedirectMiddleware redirects all traffic to HTTPS if X-Forwarded-Proto header is HTTP
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-Proto
// https://cloud.google.com/appengine/docs/flexible/go/reference/request-headers
// https://github.com/golang/go/issues/28940#issuecomment-441749380
func HTTPSRedirectMiddleware() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		proto := r.Header.Get("x-forwarded-proto")

		if proto == "http" || proto == "HTTP" {
			http.Redirect(w, r, fmt.Sprintf("https://%s%s", r.Host, r.URL), http.StatusPermanentRedirect)
			return
		}

		// TODO: Get SSL Certificate
		// if common.IsSSLEnabled() && r.TLS == nil {
		// 	http.Redirect(w, r, fmt.Sprintf("https://%s%s", r.Host, r.URL), http.StatusPermanentRedirect)
		// 	return
		// }

		next(w, r)
	})
}
