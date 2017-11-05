package utils

import (
	"fmt"

	"github.com/ashish-amarnath/slack/types"
)

// StringifyMessage returns a string representation of a message
func StringifyMessage(msg types.Message) string {
	return fmt.Sprintf("[ID=%d, Type=%s, Text=%s, Channel=%s, User=%s]",
		msg.ID, msg.Type, msg.Text, msg.Channel, msg.User)
}

// StringifySlackUser returns a string representation of a SlackUser
func StringifySlackUser(su types.SlackUser) string {
	return fmt.Sprintf("[ID=%s, FirstName=%s, LastName=%s, Email=%s]", su.ID, su.Profile.FirstName, su.Profile.LastName, su.Profile.Email)
}
