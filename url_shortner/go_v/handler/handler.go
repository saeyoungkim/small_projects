package handler

import (
	"net/http"
)

func MapHandler(pathMap map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		currUrl := r.URL.Path
		if redirectedUrl, ok := pathMap[currUrl]; ok {
			http.Redirect(rw, r, redirectedUrl, http.StatusFound)
			return
		}
		fallback.ServeHTTP(rw, r)
	}
}
