package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respWithError(w, 400, "Something Went Wrong")
}
