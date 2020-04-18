package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
)

// only supports same req limit for each endpoint atm
func (api *API) ratelimit(extractor IPExtractor, limit uint64, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ipaddr, err := extractor.Extract(r)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		conn := api.RedisPool.Get()
		closeconn := (func(conn redis.Conn) {
			err = conn.Close()
			if err != nil {
				log.Printf("Failed to close redis connection\nerr: %s", err)
			}
		})

		session := fmt.Sprintf("%s:%d", ipaddr, time.Now().Minute())
		nreqs, err := redis.Uint64(conn.Do("GET", session))
		if err != nil {
			if err == redis.ErrNil {
				nreqs = 0
			} else {
				http.Error(w, http.StatusText(500), 500)
				log.Printf("Failed to perform redis command: get\nerr: %s\n", err)
				closeconn(conn)
				return
			}
		}

		if nreqs >= limit {
			http.Error(w, http.StatusText(403), 403)
			closeconn(conn)
			return
		}

		err = conn.Send("MULTI")
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			closeconn(conn)
			return
		}
		err = conn.Send("INCR", session)
		if err != nil {
			closeconn(conn)
			http.Error(w, http.StatusText(500), 500)
			return

		}
		err = conn.Send("EXPIRE", session, 59)
		if err != nil {
			closeconn(conn)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		_, err = conn.Do("EXEC")
		if err != nil {
			closeconn(conn)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		closeconn(conn)
		h(w, r)
	}
}

func (api *API) setCORS(origin string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		h(w, r)
	}
}
