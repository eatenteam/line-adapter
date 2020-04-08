package main

import "github.com/line/line-bot-sdk-go/linebot"

func createCarousel(contents ...*linebot.BubbleContainer) *linebot.CarouselContainer {
    return &linebot.CarouselContainer{
        Type:       linebot.FlexContainerTypeCarousel,
        Contents:   contents,
    }
}
