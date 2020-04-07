package main

import (
    "log"
    "net/http"
)

func (s *Server) onlyPost(fn http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            respondHTTPErr(w, r, http.StatusMethodNotAllowed)
            log.Println("⚠️  Method not allowed (" + r.Method, "/webhook/)")
            return
        }
        fn(w, r)
    }
}
