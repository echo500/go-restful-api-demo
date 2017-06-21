package main

import "net/http"

func main() {

	router := NewRounter()

	http.ListenAndServe(":9090", router)
}
