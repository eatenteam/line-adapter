package main

import (
    "encoding/json"

    linebot "github.com/line/line-bot-sdk-go/linebot"
)

func (s *Server) handleText(m *linebot.TextMessage, rt string, _ *linebot.EventSource) error {
    switch m.Text {
    case "Malls":
        var (
            malls   []Mall
            bubbles []*linebot.BubbleContainer
        )
        resp, err := s.http.Get(s.storage + "/malls")
        if err != nil {
            return err
        }
        defer resp.Body.Close()
        err = json.NewDecoder(resp.Body).Decode(&malls)
        if err != nil {
            return err
        }
        for _, mall := range malls {
            bubbles = append(bubbles, mall.createBubble())
        }
        carousel := createCarousel(bubbles...)
        message := linebot.NewFlexMessage("Malls", carousel)
        _, err = s.bot.ReplyMessage(rt, message).Do()
        if err != nil {
            return err
        }
        return nil
    }
    return nil
}
