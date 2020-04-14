package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/moki/api.moki.codes/dotenv"
)

func base(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "%v\n", req.Method)
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	r := dotenv.NewReaderT(".env")
	r.Read()

	http.HandleFunc("/", base)

	port := os.Getenv("BACKEND_CONTAINER_PORT")
	log.Fatal(
		http.ListenAndServeTLS(
			":"+port,
			"server.crt",
			"server.key",
			nil,
		))

}
