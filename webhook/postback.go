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
            bubbles = append(bubbles, shop.createBubble("@" + strings.ToUpper(mall.Brand)))
        }
        carousel := createCarousel(bubbles...)
        message := linebot.NewFlexMessage("Shops", carousel)
        _, err = s.bot.ReplyMessage(rt, message).Do()
        if err != nil {
            return err
        }
    case "shop":
        var (
            mall        Mall
            shop        Shop
            products    []Product
            bubbles     []*linebot.BubbleContainer
        )
        resp, err := s.http.Get(s.storage + "/shops/" + value)
        if err != nil {
            return err
        }
        defer resp.Body.Close()
        err = json.NewDecoder(resp.Body).Decode(&shop)
        if err != nil {
            return err
        }
        resp, err = s.http.Get(s.storage + "/malls/" + shop.Location.Mall)
        if err != nil {
            return err
        }
        defer resp.Body.Close()
        err = json.NewDecoder(resp.Body).Decode(&mall)
        if err != nil {
            return err
        }
        for _, stock := range shop.Stock {
            id := stock.Product
            var product Product
            resp, err = s.http.Get(s.storage + "/products/" + id)
            if err != nil {
                return err
            }
            defer resp.Body.Close()
            err = json.NewDecoder(resp.Body).Decode(&product)
            if err != nil {
                return err
            }
            products = append(products, product)
        }
        txt := strings.ToUpper(shop.Brand) + " @" + strings.ToUpper(mall.Brand)
        for _, product := range products {
            bubbles = append(bubbles, product.createBubble(txt))
        }
        carousel := createCarousel(bubbles...)
        message := linebot.NewFlexMessage("Products", carousel)
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
