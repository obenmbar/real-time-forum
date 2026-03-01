package forum

import (
	"net/http"
	"strings"
)

func SafeFileServer() http.Handler {
	handler := http.FileServer(http.Dir("../frontend"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.HasSuffix(r.URL.Path, "/") {
			http.Error(w, "Access Forbidden", http.StatusForbidden)
			return
		}
		if strings.Contains(r.URL.Path, "/.") {
			http.Error(w, "Access Forbidden", http.StatusForbidden)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
