package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-")
	os.Setenv("SLACK_APP_TOKEN", "xapp-")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("I born in <year>", &slacker.CommandDefinition{
		Description: "Age Calculator",
		Examples:    []string{"I born in 1990"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			birthYear, err := strconv.Atoi(year)

			if err != nil || birthYear > time.Now().Year() {
				response.Reply("Please enter a valid year")
				return
			}

			age := time.Now().Year() - birthYear

			r := fmt.Sprintf("You are %d years old", age)
			response.Reply(r)
		},
	})

	definition := &slacker.CommandDefinition{
		Description: "Upload a text file containing sentence.",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			sentence := request.Param("sentence")
			client := botCtx.Client()
			ev := botCtx.Event()

			if ev.Channel != "" {
				client.PostMessage(ev.Channel, slack.MsgOptionText("Uploading file...", false))
				_, err := client.UploadFile(slack.FileUploadParameters{Title: "Text", Content: sentence, Channels: []string{ev.Channel}})
				if err != nil {
					fmt.Printf("Error uploading file: %s", err)
				}
			}
		},
	}
	bot.Command("upload <sentence>", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
