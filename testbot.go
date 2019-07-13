package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/RDO34/go-slack-app/respond"
	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Printf("Event Received:\n")
			switch event := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", event.ConnectionCount)

			case *slack.MessageEvent:
				user, err := api.GetUserInfo(event.User)
				if err != nil {
					log.Fatal("Error getting user")
				} else {
					fmt.Printf("Message Received: <%v>: %v\n", user.RealName, event.Text)
				}
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				if event.User != info.User.ID && strings.HasPrefix(event.Text, prefix) {
					respond.Respond(rtm, event, prefix)
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", event.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid auth credentials")
				break Loop

			default:
				// Do nothing
			}
		}
	}
}
