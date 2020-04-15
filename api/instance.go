package api

import "net/http"

// API defines server context
type API struct {
	Router *http.ServeMux
}

// Initialize function initializes API
func (api *API) Initialize() {
	api.routes()
}
