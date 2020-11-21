
package main

import (
	"log"
	"net/http"
	"os"
)

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}


func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":"+os.Args[1], Log(http.FileServer(http.Dir(os.Getenv("PWD"))))))
}
