package api

func (api *API) routes() {
	ipextractor := NewIPExtractorT()
	api.Router.HandleFunc("/newsletter/subscribers",
		api.ratelimit(ipextractor, 20, api.setCORS("*", api.subscribers())))
	api.Router.HandleFunc("/",
		api.ratelimit(ipextractor, 20, api.setCORS("*", api.base())))
}
