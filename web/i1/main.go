package main

import "net/http"

func main() {
	http.HandleFunc("/", someFunc)
	http.HandleFunc("/john", someOtherFunc)
	http.ListenAndServe(":8080", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello universe"))
	_=req
}
func someOtherFunc(w http.ReasponseWriter, req *http.Request) {
	w.Write([]byte("Hi,John"))
	_=req
}