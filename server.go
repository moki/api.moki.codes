package main

import (
	"log"
	"net/http"
	"os"

	"github.com/moki/api.moki.codes/api"
	"github.com/moki/api.moki.codes/dotenv"
)

func main() {
	r := dotenv.NewReaderT(".env")
	r.Read()

	port := os.Getenv("BACKEND_CONTAINER_PORT")

	api := api.API{}
	api.Router = http.NewServeMux()
	api.Initialize()

	log.Fatal(
		http.ListenAndServeTLS(
			":"+port,
			"certificates/server.crt",
			"certificates/server.key",
			api.Router,
		))
}
