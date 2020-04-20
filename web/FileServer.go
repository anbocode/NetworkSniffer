package web

import (
	"net/http"
)

func FileServer(prefix string) http.Handler {
	var (
		fileServer http.Handler = http.FileServer(http.Dir("./static"))
	)

	return http.StripPrefix(prefix, fileServer)
}
