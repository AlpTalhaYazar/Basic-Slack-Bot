package main

import (
	"fmt"
	"github.com/shomali11/slacker"
	"os"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-")
	os.Setenv("SLACK_APP_TOKEN", "xapp-")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

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
