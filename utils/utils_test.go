package utils

import (
	"testing"

	"github.com/ashish-amarnath/slack/types"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStringifyMessage(t *testing.T) {
	Convey("StringifyMessage", t, func() {
		Convey("Should return expected string version of a message object", func() {
			var testMsg types.Message
			testMsg.Channel = "utchannel"
			testMsg.ID = 12
			testMsg.Text = "unit test message"
			testMsg.Type = "message"
			testMsg.User = "unit-tester"
			expectedString := "[ID=12, Type=message, Text=unit test message, Channel=utchannel, User=unit-tester]"
			actualString := StringifyMessage(testMsg)
			So(actualString, ShouldEqual, expectedString)
		})
		Convey("Should use default values", func() {
			var testMsg types.Message
			expectedString := "[ID=0, Type=, Text=, Channel=, User=]"
			actualString := StringifyMessage(testMsg)
			So(actualString, ShouldEqual, expectedString)
		})
	})
}

func TestStringifySlackUser(t *testing.T) {
	Convey("StringifySlackUser", t, func() {
		Convey("Should return expected string version of a SlackUser object", func() {
			var to types.SlackUser
			to.ID = "UCRAY7US3R"
			to.Profile.FirstName = "John"
			to.Profile.LastName = "Doe"
			to.Profile.Email = "john.doe@johndoe.com"
			expected := "[ID=UCRAY7US3R, FirstName=John, LastName=Doe, Email=john.doe@johndoe.com]"
			actual := StringifySlackUser(to)
			So(actual, ShouldEqual, expected)
		})
		Convey("Should use default values", func() {
			var to types.SlackUser
			expected := "[ID=, FirstName=, LastName=, Email=]"
			actual := StringifySlackUser(to)
			So(actual, ShouldEqual, expected)
		})
	})
}
