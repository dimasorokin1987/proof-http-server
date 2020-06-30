package main

import (
	"fmt"
	//"io"
	"net/http"
	//"html"
	//"log"
	"os"
)

type page {}
func (p page) serveHTTP (w http.ResponseWriter, _ *http.Request){
  fmt.Fprint(w,"<h1>hello world")
}

func main(){
  var p page
  err := http.ListenAndServe(":"+os.Getenv("PORT"),p)
  if(err!=nil)panic(err)
	//http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		//strReq := html.EscapeString(r.URL.Path)
		//fmt.Fprintf(w, "Hello, %q\n", strReq)
		//io.WriteString(w, "Hello from a HandleFunc #1!\n")
		//io.WriteString(w, "status success/n")
	//})

	//log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
