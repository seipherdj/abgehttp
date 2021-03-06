package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	fs := http.FileServer(http.Dir(path))

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprint(res, "<h1>ABGEHTTPTransfer</h1>")
	})

	http.Handle("/static", fs)

	log.Fatal(http.ListenAndServe(":6969", nil))
}
