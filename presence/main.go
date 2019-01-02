package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nlopes/slack"
)

var (
	slackToken = os.Getenv("SLACK_OAUTH_TOKEN")

	slackChannel = os.Getenv("SLACK_CHANNEL")

	api = slack.New(
		slackToken,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)))
)

// PresenceEvent - information from mqtt broker
type PresenceEvent struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Confidence   string `json:"confidence"`
	Manufacturer string `json:"manufacturer"`
}

func handler(event PresenceEvent) {

	confidence, _ := strconv.Atoi(event.Confidence)

	switch confidence {
	case 100:
		slackRelay(event.Name, "has arrived at the office", "#36a64f")
	case 0:
		slackRelay(event.Name, "has left the office", "#ff0000")
	}
}

func slackRelay(name string, msg string, color string) {
	attatchment := slack.Attachment{
		Text:  fmt.Sprintf("*%s* %s", name, msg),
		Color: color,
	}

	channelID, timestamp, err := api.PostMessage(slackChannel, slack.MsgOptionAttachments(attatchment))
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}

func main() {
	lambda.Start(handler)
}
