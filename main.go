package main

import (
    "log"
    "os"

    "github.com/nlopes/slack"
)


func run(sc *slack.Client) int {
    rtm := sc.NewRTM()
    go rtm.ManageConnection()

    for {
        select {
        case msg := <-rtm.IncomingEvents:
            switch ev := msg.Data.(type) {
            case *slack.HelloEvent:
                log.Print("Hello Event")

            case *slack.MessageEvent:
                log.Printf("Message: %v\n", ev)
                log.Printf("Message.Msg: %v\n", ev.Msg)
                log.Printf("Message.Msg.Text: %s\n", ev.Msg.Text)
                output := handle(ev.Msg.Text)
                rtm.SendMessage(rtm.NewOutgoingMessage(output, ev.Channel))

            case *slack.InvalidAuthEvent:
                log.Print("Invalid credentials")
                return 1

            }
        }
    }
}

func main() {
    token := os.Getenv("SLACK_CURL_BOT_TOKEN")
    sc := slack.New(token)
    os.Exit(run(sc))
}
