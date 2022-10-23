package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	http.HandleFunc("/", greet)
	http.ListenAndServe(net.JoinHostPort("", port), nil)
}
