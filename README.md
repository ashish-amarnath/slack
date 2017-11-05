# slack
Package implements a framework for writing slack bots using web sockets and Real Time Messaging (RTM) APIs. More info on RTM API's for slack is available at https://api.slack.com/rtm.

Example illustrating the usage of this package is in `examples/main.go`
The example is a bot that replies "Bot says: Roger!!" for any message that is at-ed at the bot.


# Building and running the example bot.
`make buildbot` build the executable, `roger-bot` under the `build` directory.
The bot can be run as `./build/roger-bot -slackbotToken=<SLACK GENERATED PRIVATE TOKEN FOR YOUR BOT>`
Refer https://api.slack.com/bot-users for information on how to add new bot users and get a token for your bot. 

Go on... Bot-ify your slack world!

