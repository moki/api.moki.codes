package api

import "net/http"

func (api *API) setCORS(h http.HandlerFunc, origin string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		h(w, r)
	}
}
