package main

import "github.com/line/line-bot-sdk-go/linebot"

type Mall struct {
    Id          string      `json:"id"`
    Brand       string      `json:"brand"`
    Description string      `json:"description"`
    Shops       []string    `json:"shops"`
}

func (m *Mall) createBubble() linebot.FlexContainer {
    postback := &linebot.PostbackAction{
        Label:          "See Restuarants",
        Data:           "mall=" + m.Id,
        DisplayText:    m.Brand,
    }
    body := []linebot.FlexComponent{
        bubbleTextTitle("🛍️ " + m.Brand),
        bubbleSeparator(),
        bubbleTextBodyRegular(m.Description, 0),
    }
    footer := []linebot.FlexComponent{
        bubbleButton(postback),
    }
    return &linebot.BubbleContainer{
        Type:   linebot.FlexContainerTypeBubble,
        Size:   linebot.FlexBubbleSizeTypeMega,
        Body:   bubbleBoxVertical(body...),
        Footer: bubbleBoxHorizontal(footer...),
        Styles: bubbleStyleDefault(),
    }
}
