package main

import (
	"os"

	"github.com/go-microbot/telegram/api"
	"github.com/go-microbot/telegram/bot"
)

func main() {
	botAPI := api.NewTelegramAPI(os.Getenv("TOKEN"))
	tBot := bot.NewTelegramBot(&botAPI)
	tBot.Test()
}
