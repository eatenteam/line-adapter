package main

import (
    "log"
    "strings"
    "encoding/json"

    linebot "github.com/line/line-bot-sdk-go/linebot"
)

func (s *Server) handlePostback(postback *linebot.Postback, rt string, _ *linebot.EventSource) error {
    key, value := extractData(postback)
    switch key {
    case "mall":
        var (
            mall    Mall
            shops   []Shop
            bubbles []*linebot.BubbleContainer
        )
        resp, err := s.http.Get(s.storage + "/malls/" + value)
        if err != nil {
            return err
        }
        defer resp.Body.Close()
        err = json.NewDecoder(resp.Body).Decode(&mall)
        if err != nil {
            return err
        }
        for _, id := range mall.Shops {
            var shop Shop
            resp, err = s.http.Get(s.storage + "/shops/" + id)
            if err != nil {
                return err
            }
            defer resp.Body.Close()
            err = json.NewDecoder(resp.Body).Decode(&shop)
            if err != nil {
                return err
            }
            shops = append(shops, shop)
        }
        for _, shop := range shops {
            bubbles = append(bubbles, shop.createBubble(strings.ToUpper(mall.Brand)))
        }
        carousel := createCarousel(bubbles...)
        message := linebot.NewFlexMessage("Malls", carousel)
        _, err = s.bot.ReplyMessage(rt, message).Do()
        if err != nil {
            return err
        }
    }
    return nil
}

func extractData(pb *linebot.Postback) (string, string) {
    result := strings.Split(pb.Data, "=")
    log.Println("🌏  " + result[0] + ": " + result[1])
    return result[0], result[1]
}
