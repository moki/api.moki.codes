package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func endpoints() []string {
	return []string{"/", "/newsletter/subscribers"}
}

func (api *API) base() http.HandlerFunc {
	type Response struct {
		Endpoints []string `json:"endpoints"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, http.StatusText(405), 405)
			return
		}
		resp := Response{endpoints()}
		encoded, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(encoded)
		if err != nil {
			log.Printf("Write failed: %v", err)
		}
	}
}
