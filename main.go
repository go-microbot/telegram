package main

import (
	"os"

	"github.com/go-microbot/telegram/api"
	"github.com/go-microbot/telegram/bot"
)

func main() {
	botAPI := api.New(os.Getenv("TOKEN"))
	tBot := bot.New(&botAPI)
	tBot.Test()
}
