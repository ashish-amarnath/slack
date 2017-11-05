package slackrtm

import (
	"fmt"
	"testing"

	"github.com/ashish-amarnath/slack/types"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStartSlackRTM(t *testing.T) {
	Convey("startSlackRTM", t, func() {
		Convey("should return failure if token is nil", func() {
			_, _, actualErr := startSlackRTM("")
			expectedErr := fmt.Errorf("expected non-empty slackbot integration token, got [%s]", "")
			So(actualErr, ShouldResemble, expectedErr)
		})
	})
}

func TestGetSlackRTMURL(t *testing.T) {
	Convey("getSlackRTMURL should return the correct RTM URL for the supplied token", t, func() {
		expectedRTMURL := `https://slack.com/api/rtm.start?token=unitTestToken`
		actualRTMURL := getSlackRTMURL("unitTestToken")
		So(actualRTMURL, ShouldResemble, expectedRTMURL)
	})
}

func TestParseRtmStartResponse(t *testing.T) {
	Convey("parseRtmStartResponse", t, func() {
		Convey("should fail when called with invalid response bytes", func() {
			testString := `{"ok":true,"self":{"id":"U725Q50AY","name":"ashisha-bot"`
			_, err := parseRtmStartResponse([]byte(testString))
			So(err, ShouldNotBeNil)
		})
		Convey("Should parse a valid byte array into RTM message format: Case 1", func() {
			testString := `{"ok":true,"self":{"id":"U725Q50AY","name":"foo-bot"},"url":"wss:\/\/lbmulti-n5li.lb.slack-msgs.com\/websocket\/fZjKq0UKF_zjK3ZpAldHTjjW4ZuQKsalcPXv-UY42hqEWZqYngRRALTzVCcMs600xuTgKer0VYIY7XjBPYAhD0MnjYcSwEHVDg-fi5cHF-Wdc80xMOb1ExHTfjmwlqsfilagHkIk3F1f5qUt6nGGWaNFvqKbqEpRFJ8jA8dxn4c=","mpims":null}`
			actualRespMsg, _ := parseRtmStartResponse([]byte(testString))
			var expectedRtmRespMsg types.ResponseRtmStart
			expectedRtmRespMsg.Ok = true
			expectedRtmRespMsg.Bot.ID = "U725Q50AY"
			expectedRtmRespMsg.Bot.Name = "foo-bot"
			expectedRtmRespMsg.URL = `wss://lbmulti-n5li.lb.slack-msgs.com/websocket/fZjKq0UKF_zjK3ZpAldHTjjW4ZuQKsalcPXv-UY42hqEWZqYngRRALTzVCcMs600xuTgKer0VYIY7XjBPYAhD0MnjYcSwEHVDg-fi5cHF-Wdc80xMOb1ExHTfjmwlqsfilagHkIk3F1f5qUt6nGGWaNFvqKbqEpRFJ8jA8dxn4c=`
			So(actualRespMsg.Ok, ShouldEqual, expectedRtmRespMsg.Ok)
			So(actualRespMsg.Error, ShouldEqual, expectedRtmRespMsg.Error)
			So(actualRespMsg.URL, ShouldEqual, expectedRtmRespMsg.URL)
			So(actualRespMsg.Bot.ID, ShouldEqual, expectedRtmRespMsg.Bot.ID)
			So(actualRespMsg.Bot.Name, ShouldEqual, expectedRtmRespMsg.Bot.Name)
		})
		Convey("Should parse a valid byte array into RTM message format: Case 2", func() {
			testString := `{"ok":false,"self":{"id":"U725Q50AY","name":"foo-bot"},"url":"wss:\/\/lbmulti-n5li.lb.slack-msgs.com\/websocket\/fZjKq0UKF_zjK3ZpAldHTjjW4ZuQKsalcPXv-UY42hqEWZqYngRRALTzVCcMs600xuTgKer0VYIY7XjBPYAhD0MnjYcSwEHVDg-fi5cHF-Wdc80xMOb1ExHTfjmwlqsfilagHkIk3F1f5qUt6nGGWaNFvqKbqEpRFJ8jA8dxn4c=","mpims":null,"error":"unit test error"}`
			actualRespMsg, _ := parseRtmStartResponse([]byte(testString))
			var expectedRtmRespMsg types.ResponseRtmStart
			expectedRtmRespMsg.Ok = false
			expectedRtmRespMsg.Error = "unit test error"
			expectedRtmRespMsg.Bot.ID = "U725Q50AY"
			expectedRtmRespMsg.Bot.Name = "foo-bot"
			expectedRtmRespMsg.URL = `wss://lbmulti-n5li.lb.slack-msgs.com/websocket/fZjKq0UKF_zjK3ZpAldHTjjW4ZuQKsalcPXv-UY42hqEWZqYngRRALTzVCcMs600xuTgKer0VYIY7XjBPYAhD0MnjYcSwEHVDg-fi5cHF-Wdc80xMOb1ExHTfjmwlqsfilagHkIk3F1f5qUt6nGGWaNFvqKbqEpRFJ8jA8dxn4c=`
			So(actualRespMsg.Ok, ShouldEqual, expectedRtmRespMsg.Ok)
			So(actualRespMsg.Error, ShouldEqual, expectedRtmRespMsg.Error)
			So(actualRespMsg.URL, ShouldEqual, expectedRtmRespMsg.URL)
			So(actualRespMsg.Bot.ID, ShouldEqual, expectedRtmRespMsg.Bot.ID)
			So(actualRespMsg.Bot.Name, ShouldEqual, expectedRtmRespMsg.Bot.Name)
		})
	})
}
