package respond

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

func Respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {
	var response string
	text := msg.Text
	text = strings.TrimPrefix(text, prefix)
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	acceptedGreetings := map[string]bool{
		"hello":    true,
		"good day": true,
		"hi":       true,
	}
	acceptedHowAreYou := map[string]bool{
		"how are you?":                 true,
		"how's it going?":              true,
		"have you become setient yet?": true,
	}

	if acceptedGreetings[text] {
		response = fmt.Sprintf("Greetings earthling :worrymuhahaha:")
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else if acceptedHowAreYou[text] {
		response = ":pog: That is none of your business"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	}
}
