package api

import (
	"net/http"

	"github.com/gomodule/redigo/redis"
)

// API defines server context
type API struct {
	Router    *http.ServeMux
	RedisPool *redis.Pool
}

// Initialize function initializes API
func (api *API) Initialize() {
	api.routes()
}
