package main

import (
    "log"
    "net/http"
)

func (s *Server) handleHealthCheck() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        respond(w, r, http.StatusOK, nil)
        log.Println("🌡️  Healthcheck passed")
        return
    }
}
