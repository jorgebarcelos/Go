package main

import (
	"net/http"
	edp "endpoints/endpoints"
)

func main(){
	http.HandleFunc("/", edp.HomePage)
	http.HandleFunc("/contact", edp.ContactPage)
	http.ListenAndServe(":5000", nil)
}