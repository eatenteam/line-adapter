package main

import (
    "net/http"
    "log"

    linebot     "github.com/line/line-bot-sdk-go/linebot"
)

type Server struct {
    http    *http.Client
    bot     *linebot.Client
    storage string
}

func main() {
    c := &Config{}
    if err := c.load(); err != nil {
        log.Fatalln("Loading Environemnt Error: ", err)
    }
    client := createHTTPClient()
    bot, err := createBot(c.LineToken, c.LineSecret)
    if err != nil {
        log.Fatalln("LINEBot Error: ", err)
    }
    s := &Server{client, bot, c.StorageUrl}
    mux := http.NewServeMux()
    mux.HandleFunc("/webhook", s.onlyPost(s.handleWebhooks()))
    mux.HandleFunc("/webhook/", s.onlyPost(s.handleWebhooks()))
    mux.HandleFunc("/health", s.handleHealthCheck())
    mux.HandleFunc("/health/", s.handleHealthCheck())
    log.Println("🌏  Server listening on ", ":" + c.Port)
    log.Fatalln(http.ListenAndServe(":" + c.Port, mux))
}

func createBot(token, secret string) (*linebot.Client, error) {
    client := &http.Client{}
    log.Println("🔥  Bootstrapping LINEBot...")
    bot, err := linebot.New(secret, token, linebot.WithHTTPClient(client))
    if err == nil {
        log.Println("✅  Successfully bootstrapping LINEBot")
    }
    return bot, err
}

func createHTTPClient() *http.Client {
    log.Println("🔥  Bootstrapping HTTP Client...")
    return &http.Client{}
}
