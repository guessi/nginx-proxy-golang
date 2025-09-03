package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var listenPort = flag.Int("p", 9000, "listen port")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// basic information
		fmt.Fprintf(w, "Host => %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr => %q\n\n", r.RemoteAddr)

		// print out all request headers
		for k, v := range r.Header {
			fmt.Fprintf(w, "%q => %q\n", k, v)
		}
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *listenPort), nil))
}
