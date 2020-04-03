package main

import (
    "context"
    "flag"
    "net/http"
    "log"
    "time"

    linebot     "github.com/line/line-bot-sdk-go/linebot"
    mongo       "go.mongodb.org/mongo-driver/mongo"
    options     "go.mongodb.org/mongo-driver/mongo/options"
    readpref    "go.mongodb.org/mongo-driver/mongo/readpref"
)

type Server struct {
    db  *mongo.Client
    bot *linebot.Client
}

func main() {
    var (
        addr        = flag.String("addr", ":80", "server address")
        mongo       = flag.String("mongo", "mongodb://localhost", "mongodb address")
        lineToken   = flag.String("token", "", "line channel access token")
        lineSecret  = flag.String("secret", "", "line channel secret")
    )
    flag.Parse()
    db, err := connectMongo(*mongo)
    if err != nil {
        log.Fatalln("Mongo Error: ", err)
    }
    bot, err := connectLine(*lineToken, *lineSecret)
    if err != nil {
        log.Fatalln("LINEBot Error: ", err)
    }
    s := &Server{db, bot}
    mux := http.NewServeMux()
    mux.HandleFunc("/webhook/", s.onlyPost(s.handleWebhooks()))
    mux.HandleFunc("/health/", s.handleHealthCheck())
    log.Println("🌏  Server listening on ", *addr)
    http.ListenAndServe(*addr, mux)
}

func connectMongo(uri string) (*mongo.Client, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    log.Println("🔥  Connecting to MongoDB Cluster...")
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }
    log.Println("🔥  Pinging to MongoDB Cluster...")
    ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        return nil, err
    }
    log.Println("✅  Successfully connect MongoDB Cluster")
    return client, nil
}

func connectLine(token, secret string) (*linebot.Client, error) {
    client := &http.Client{}
    log.Println("🔥  Bootstrapping LINEBot...")
    bot, err := linebot.New(secret, token, linebot.WithHTTPClient(client))
    if err == nil {
        log.Println("✅  Successfully bootstrapping LINEBot")
    }
    return bot, err
}
