package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dlclark/regexp2"
	"github.com/gomodule/redigo/redis"
)

var emailChecker = regexp2.MustCompile("^[^\\s@]+@[^|\\s@]+\\.[^\\s@]+$", 0)
var nameChecker = regexp2.MustCompile("^(?!.*\\.\\.)(?!.*\\.$)[^\\W][\\w.]{0,29}$", 0)

func (api *API) subscribers() http.HandlerFunc {
	rgxmch := func(s string, r *regexp2.Regexp) bool {
		doesmatch, err := r.MatchString(s)
		if err != nil {
			log.Printf("Failed to perform regexp mathcing")
		}
		return doesmatch
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, http.StatusText(405), 405)
			return
		}
		email := r.PostFormValue("email")
		if email == "" || len(email) > 254 || !rgxmch(email, emailChecker) {
			http.Error(w, http.StatusText(400)+": invalid email", 400)
			return
		}
		name := r.PostFormValue("name")
		if name == "" || len(name) > 30 || !rgxmch(name, nameChecker) {
			http.Error(w, http.StatusText(400)+": invalid name", 400)
			return
		}

		conn := api.RedisPool.Get()
		defer (func() {
			err := conn.Close()
			if err != nil {
				log.Printf("Failed to close redis connection: %v", err)
			}
		})()
		succ, err := redis.Bool(conn.Do("HSETNX", "subscribers", email, name))
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			log.Printf("Failed to perform redis: hsetns command")
			return
		}
		if succ {
			w.WriteHeader(201)
			fmt.Fprintf(w, "%d: %s\n", 201, http.StatusText(201))
		} else {
			w.WriteHeader(409)
			fmt.Fprintf(w, "%d: %s\n", 409, http.StatusText(409))
		}
	}
}
