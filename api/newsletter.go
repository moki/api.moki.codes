package api

import "net/http"

func (api *API) subscribe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, http.StatusText(405), 405)
		}
	}
}
