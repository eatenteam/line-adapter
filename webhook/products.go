package main

import (
    "strconv"

    "github.com/line/line-bot-sdk-go/linebot"
)

type Product struct {
    Id              string  `json:"id"`
    Name            string  `json:"name"`
    OriginalPrice   int     `json:"ori_price"`
    PromotionPrice  int     `json:"pro_price"`
}

func (p *Product) createBubble(txt string) *linebot.BubbleContainer {
    postback := &linebot.PostbackAction{
        Label:          "Add to Cart",
        Data:           "toCart=" + p.Id,
        DisplayText:    p.Name,
    }
    header := []linebot.FlexComponent{
        bubbleTextHeader(txt),
    }
    body := []linebot.FlexComponent{
        bubbleTextTitle(p.Name),
        bubbleSeparator(),
        bubbleBoxBaseline([]linebot.FlexComponent{
            bubbleTextBodyBold("Original:", 2),
            bubbleTextBodyRegular(strconv.Itoa(p.OriginalPrice), 4),
        }...),
        bubbleBoxBaseline([]linebot.FlexComponent{
            bubbleTextBodyBold("Now:", 2),
            bubbleTextBodyRegular(strconv.Itoa(p.PromotionPrice), 4),
        }...),
    }
    footer := []linebot.FlexComponent{
        bubbleButton(postback),
    }
    return &linebot.BubbleContainer{
        Type:   linebot.FlexContainerTypeBubble,
        Size:   linebot.FlexBubbleSizeTypeMega,
        Header: bubbleBoxHeader(header...),
        Body:   bubbleBoxBody(body...),
        Footer: bubbleBoxFooter(footer...),
        Styles: bubbleStyleDefault(),
    }
}
