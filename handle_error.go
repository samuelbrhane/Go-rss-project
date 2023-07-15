package main

import "net/http"

func handleError(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 200, "Something went wrong")
}
