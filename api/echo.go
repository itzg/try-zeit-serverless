package api

import (
	"fmt"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	for k, vals := range r.URL.Query() {
		fmt.Fprintf(w, "%s = %s\n", k, strings.Join(vals, ","))
	}
}
