package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ashish-amarnath/slack/pkg"
	"github.com/ashish-amarnath/slack/types"
	"github.com/ashish-amarnath/slack/utils"
	"github.com/golang/glog"
)

var (
	slackbotToken *string
)

// Simple Echo Bot to  demonstrating how to use this package.
func main() {
	slackbotToken = flag.String("slackbotToken", "", "Slack generated token for the bot")
	flag.Parse()

	slackConn := slackrtm.NewSlackServerConn(*slackbotToken)

	glog.V(1).Infoln("Slackbot listening for messages to process...")
	for {
		msg, err := slackConn.ReadMessage()
		if err != nil {
			glog.Errorf("Failed to read message sent to slackbot. err=%s\n", err.Error())
			continue
		}
		if msg.Type != types.MessageType || !strings.HasPrefix(msg.Text, "<@"+slackConn.UserID+">") {
			glog.V(9).Infof("Ignoring message %s\n", utils.StringifyMessage(msg))
			continue
		}

		resp := slackrtm.CreateRespForReq(msg)
		resp.Text = fmt.Sprintf("Bot says: Roger!!")
		slackConn.SendMessage(resp)
	}
}
