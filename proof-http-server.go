package main

import (
	//"fmt"
	"io"
	"net/http"
	//"html"
	"log"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		//strReq := html.EscapeString(r.URL.Path)
		//fmt.Fprintf(w, "Hello, %q\n", strReq)
		//io.WriteString(w, "Hello from a HandleFunc #1!\n")
		io.WriteString(w, "status success")
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
