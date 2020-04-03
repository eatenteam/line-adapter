package main

import (
    "net/http"
    "log"
)

func (s *Server) handleWebhooks() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        events, err := s.bot.ParseRequest(r)
        if err != nil {
            respondWithCode(w, r, http.StatusBadRequest)
            log.Println("⚠️  Bad webhook request")
            return
        }
        log.Println("🌏  Webhook activated")
        for _, event := range events {
            log.Println("🌏  Recieved -", event.Type)
        }
    }
}
