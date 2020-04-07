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

func (s *Server) handleText(m *linebot.TextMessage, rt string, _ *linebot.EventSource) error {
    if m.Text != "mall" {
        log.Print("not a mall message")
        return nil
    }
    paragon := Mall{
        Id:             "5e8ca4189f5eaf9ada7c3c47",
        Brand:          "Paragon",
        Description:    "Paul, BreadTalk, Flavor Field, Starbuck, and many more!",
    }
    bubble := paragon.createBubble()
    message := linebot.NewFlexMessage("Mall", bubble)
    _, err := s.bot.ReplyMessage(rt, message).Do()
    if err != nil {
      return err
    }
    return nil
}
