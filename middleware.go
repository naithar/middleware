package middleware

import (
	"github.com/go-chi/chi"
	"net/http"
)

// FixSlashes is Middleware for Chi Router that removes repeated slashes
// and removes trailing slash if required
func FixSlashes(saveTrailing bool) func(next http.Handler) http.Handler {
	fn := func (next http.Handler) http.Handler {
		ifn := func(w http.ResponseWriter, r *http.Request) {
			var path string
			rctx := chi.RouteContext(r.Context())
			if len(rctx.RoutePath) > 0 {
				path = rctx.RoutePath
			} else {
				path = r.URL.Path
			}
			rctx.RoutePath = CleanPath(path, saveTrailing)

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(ifn)
	}
	return fn
}