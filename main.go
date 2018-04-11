package main

import (
    "log"
    "os"

    "github.com/nlopes/slack"
    "strings"
)

func handle(input string) string {
    text := input
    var output string
    if strings.HasPrefix(text, "curl ") {
        log.Printf("it's curl\n")
        output = "calling curl"
    } else {
        log.Printf("it's NOT curl\n")
        output = "(nothing to do)"
    }
    return output
}

func run(api *slack.Client) int {
    rtm := api.NewRTM()
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
    token := os.Getenv("TOKEN")
    api := slack.New(token)
    os.Exit(run(api))
}
