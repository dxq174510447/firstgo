package main

import "net/http"

func handleRequest(response http.ResponseWriter, request *http.Request) {

}

func main() {

	http.Handle("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
