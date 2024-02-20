package handlers

import "net/http"

func Fileserver() http.Handler {
	return http.StripPrefix("/app", http.FileServer(http.Dir(".")))
}
