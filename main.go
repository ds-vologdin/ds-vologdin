package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

const token = "xoxb-6832...."
const botName = "keel"

func main() {
	api := slack.New(token, slack.OptionDebug(true))
	channels, err := api.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, channel := range channels {
		fmt.Println(channel.Name, channel.ID)
	}
	channel, err := api.JoinChannel("k8s")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Join to channel %s\n", channel.ID)
	fmt.Println("----------------------------------------")
	fmt.Println("Send message")
	params := slack.NewPostMessageParameters()
	params.Username = botName

	var mgsOpts []slack.MsgOption
	mgsOpts = append(mgsOpts, slack.MsgOptionPostMessageParameters(params))
	mgsOpts = append(mgsOpts, slack.MsgOptionText("Some text", false))

	channelID, timestamp, err := api.PostMessage(channel.ID, mgsOpts...)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
	fmt.Println("----------------------------------------")

	users, err := api.GetUsers()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	for _, user := range users {
		fmt.Printf("%s - %s is bot %v\n", user.Name, user.ID, user.IsBot)
	}
}
