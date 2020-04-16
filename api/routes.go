package api

func (api *API) routes() {
	api.Router.HandleFunc("/newsletter/subscribers", api.setCORS(api.subscribers(), "*"))
	api.Router.HandleFunc("/", api.setCORS(api.base(), "*"))
}
