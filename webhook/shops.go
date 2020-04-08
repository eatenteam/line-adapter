package main

import "github.com/line/line-bot-sdk-go/linebot"

type Shop struct {
    Id          string      `json:"id"`
    Brand       string      `json:"brand"`
    Promo       string      `json:"promo"`
    Description string      `json:"description"`
    Stock       []Stock     `json:"stock"`
    Open        string      `json:"open"`
    Close       string      `json:"close"`
}

type Stock struct {
    Product     string      `json:"product"`
    Quantity    int         `json:"quantity"`
}

func (s *Shop) createBubble(txt string) *linebot.BubbleContainer {
    postback := &linebot.PostbackAction{
        Label:          "View Menu",
        Data:           "shop=" + s.Id,
        DisplayText:    s.Brand,
    }
    header := []linebot.FlexComponent{
        bubbleTextHeader(txt),
    }
    body := []linebot.FlexComponent{
        bubbleTextTitle("🍰 " + s.Brand),
        bubbleTextPromo(s.Promo, 0),
        bubbleSeparator(),
        bubbleBoxBaseline([]linebot.FlexComponent{
            bubbleTextBodyBold("Popular", 2),
            bubbleTextBodyRegular(s.Description, 4),
        }...),
        bubbleBoxBaseline([]linebot.FlexComponent{
            bubbleTextBodyBold("Open", 2),
            bubbleTextBodyRegular(formatTime(s.Open), 4),
        }...),
        bubbleBoxBaseline([]linebot.FlexComponent{
            bubbleTextBodyBold("Close", 2),
            bubbleTextBodyRegular(formatTime(s.Close), 4),
        }...),
    }
    footer := []linebot.FlexComponent{
        bubbleSpacer(),
        bubbleButton(postback),
    }
    return &linebot.BubbleContainer{
        Type:   linebot.FlexContainerTypeBubble,
        Size:   linebot.FlexBubbleSizeTypeGiga,
        Header: bubbleBoxHeader(header...),
        Body:   bubbleBoxBody(body...),
        Footer: bubbleBoxFooter(footer...),
        Styles: bubbleStyleDefault(),
    }
}

func formatTime(time string) string {
    runes := []rune(time)
    hh := string(runes[:2])
    mm := string(runes[2:])
    return hh + ":" + mm
}
