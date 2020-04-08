package main

import "github.com/line/line-bot-sdk-go/linebot"

const (
    colorPrimary    = "#392F5A"
    colorSecondary  = "#FFF8F0"
    colorSpecial    = "#ED2939"
)

func bubbleStyleDefault() *linebot.BubbleStyle {
    return &linebot.BubbleStyle{
        Header: &linebot.BlockStyle{BackgroundColor: colorSecondary},
        Body:   &linebot.BlockStyle{BackgroundColor: colorPrimary},
        Footer: &linebot.BlockStyle{BackgroundColor: colorPrimary},
    }
}

func bubbleBoxVertical(contents ...linebot.FlexComponent) *linebot.BoxComponent {
    return &linebot.BoxComponent{
        Type:       linebot.FlexComponentTypeBox,
        Layout:     linebot.FlexBoxLayoutTypeVertical,
        Contents:   contents,
    }
}

func bubbleBoxHorizontal(contents ...linebot.FlexComponent) *linebot.BoxComponent {
    return &linebot.BoxComponent{
        Type:       linebot.FlexComponentTypeBox,
        Layout:     linebot.FlexBoxLayoutTypeHorizontal,
        Contents:   contents,
    }
}

func bubbleBoxBody(contents ...linebot.FlexComponent) *linebot.BoxComponent {
    body := *bubbleBoxVertical()
    body.Spacing = linebot.FlexComponentSpacingTypeMd
    return &body
}

func bubbleButton(action linebot.TemplateAction) *linebot.ButtonComponent {
    return &linebot.ButtonComponent{
        Type:   linebot.FlexComponentTypeButton,
        Action: action,
        Color:  colorSecondary,
        Style:  linebot.FlexButtonStyleTypeSecondary,
    }
}

func bubbleTextHeader(txt string) *linebot.TextComponent {
    return &linebot.TextComponent{
        Type:   linebot.FlexComponentTypeText,
        Text:   txt,
        Weight: linebot.FlexTextWeightTypeBold,
        Color:  colorPrimary,
    }
}

func bubbleTextTitle(txt string) *linebot.TextComponent {
    return &linebot.TextComponent{
        Type:   linebot.FlexComponentTypeText,
        Text:   txt,
        Size:   linebot.FlexTextSizeTypeXl,
        Weight: linebot.FlexTextWeightTypeBold,
        Color:  colorSecondary,
    }
}

func bubbleTextBodyRegular(txt string, flex int) *linebot.TextComponent {
    return &linebot.TextComponent{
        Type:   linebot.FlexComponentTypeText,
        Text:   txt,
        Size:   linebot.FlexTextSizeTypeMd,
        Color:  colorSecondary,
        Wrap:   true,
        Flex:   linebot.IntPtr(flex),
    }
}

func bubbleTextBodyBold(txt string, flex int) *linebot.TextComponent {
    return &linebot.TextComponent{
        Type:   linebot.FlexComponentTypeText,
        Text:   txt,
        Size:   linebot.FlexTextSizeTypeMd,
        Weight: linebot.FlexTextWeightTypeBold,
        Color:  colorSecondary,
        Wrap:   true,
        Flex:   linebot.IntPtr(flex),
    }
}

func bubbleTextPrice(txt string, flex int) *linebot.TextComponent {
    return &linebot.TextComponent{
        Type:   linebot.FlexComponentTypeText,
        Text:   txt,
        Size:   linebot.FlexTextSizeTypeMd,
        Weight: linebot.FlexTextWeightTypeBold,
        Color:  colorSecondary,
        Wrap:   true,
        Flex:   linebot.IntPtr(flex),
        Align:  linebot.FlexComponentAlignTypeEnd,
    }
}

func bubbleSeparator() *linebot.SeparatorComponent {
    return &linebot.SeparatorComponent{
        Type:   linebot.FlexComponentTypeSeparator,
    }
}

func bubbleSpacer() *linebot.SpacerComponent {
    return &linebot.SpacerComponent{
        Type:   linebot.FlexComponentTypeSpacer,
        Size:   linebot.FlexSpacerSizeTypeXxl,
    }
}
