package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ReponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, %s, ¡This is a server!", r.URL.Path[1:])
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
  if err := http.ListenAndServe(":8080", nil) ; err != nil{
  log.Fatal(err)
  }
}
