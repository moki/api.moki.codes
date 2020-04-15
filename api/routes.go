package api

func (api *API) routes() {
	api.Router.HandleFunc("/newsletter/subscribe", api.subscribe())
	api.Router.HandleFunc("/", api.base())
}
