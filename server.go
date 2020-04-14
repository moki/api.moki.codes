package main

import (
	"fmt"
	"log"
	"net/http"

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
	http.HandleFunc("/", base)
	r := dotenv.NewReaderT(".env")
	r.Read()
	log.Fatal(http.ListenAndServeTLS(":80", "server.crt", "server.key", nil))
}
