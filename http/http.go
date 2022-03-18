package http

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.Write([]byte("Hello World"))
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Index)

	return mux
}
