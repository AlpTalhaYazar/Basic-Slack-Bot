package main
import (
	"github.com/shomali11/slacker"
	"os"
)
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-")
	os.Setenv("SLACK_APP_TOKEN", "xapp-")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
}
