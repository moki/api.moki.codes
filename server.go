package main

import (
	"fmt"
	"log"
	"net/http"
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

	log.Fatal(http.ListenAndServeTLS(":80", "server.crt", "server.key", nil))
}
