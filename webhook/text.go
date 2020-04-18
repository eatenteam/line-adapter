package main

import "github.com/line/line-bot-sdk-go/linebot"

const (
    partnerAltText   = "Become partner for various benefits!"
    shopImage = "https://sgp1.digitaloceanspaces.com/winnaries/eaten/bubble-shop.png"
    shopForm  = "https://forms.gle/Sj1irUkJAykef5sv9"
    deliImage = "https://sgp1.digitaloceanspaces.com/winnaries/eaten/bubble-delivery.png"
    deliForm  = "https://forms.gle/B5eyy38J7Gga7x4Q8"
)

func (s *Server) handleText(m *linebot.TextMessage, rt string, _ *linebot.EventSource) error {
    switch m.Text {
    case "สมัครเป็นพาร์ทเนอร์":
        carousel := createRegistrationCarousel()
        message := linebot.NewFlexMessage(partnerAltText, carousel)
        _, err := s.bot.ReplyMessage(rt, message).Do()
        if err != nil {
            return err
        }
        return nil
    }
    return nil
}

func createRegistrationCarousel() *linebot.CarouselContainer {
    var contents []*linebot.BubbleContainer
    contents[0] = &linebot.BubbleContainer{
        Type:       "bubble",
        Direction:  "ltr",
        Hero:       &linebot.ImageComponent{
            Type:           "image",
            URL:            shopImage,
            Size:           "full",
            AspectRatio:    "3:4",
            AspectMode:     "cover",
            Action:         &linebot.URIAction{URI: shopForm},
        },
    }
    contents[1] = &linebot.BubbleContainer{
        Type:       "bubble",
        Direction:  "ltr",
        Hero:       &linebot.ImageComponent{
            Type:           "image",
            URL:            deliImage,
            Size:           "full",
            AspectRatio:    "3:4",
            AspectMode:     "cover",
            Action:         &linebot.URIAction{URI: deliForm},
        },
    }
    return createCarousel(contents...)
}
