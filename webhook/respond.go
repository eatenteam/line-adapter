package main

import "net/http"

func respondWithCode(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
}
