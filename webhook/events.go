package main

import (
    "log"

    linebot "github.com/line/line-bot-sdk-go/linebot"
)

func (s *Server) routeEvents(e *linebot.Event) {
    switch e.Type {
    case linebot.EventTypeMessage:
        switch message := e.Message.(type) {
        case *linebot.TextMessage:
            if err := s.handleText(message, e.ReplyToken, e.Source); err != nil {
                log.Print(err)
            }
        }
    }
}
