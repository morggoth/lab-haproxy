package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServeTLS(":8000", "/etc/ssl/localhost/localhost.pem", "/etc/ssl/localhost/localhost.key", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Hello from %s with TLS\n", hostname)
}
