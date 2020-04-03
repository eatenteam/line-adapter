package main

import (
    "log"
    "net/http"
)

func (s *Server) handleHealthCheck() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        respondWithCode(w, r, http.StatusOK)
        log.Println("🌡️  Healthcheck passed")
        return
    }
}
