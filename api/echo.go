package api

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	
	w.Write([]byte(r.URL.Path))
}