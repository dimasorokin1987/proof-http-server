package main

import (
	"fmt"
	"io"
	"net/http"
	"html"
	"log"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		strReq := html.EscapeString(r.URL.Path)
		fmt.Fprintf(w, "Hello, %q\n", strReq)
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
