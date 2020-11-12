package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-microbot/telegram/api"
	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/bot"
	"github.com/go-microbot/telegram/models"
	"github.com/go-microbot/telegram/query"
)

func main() {
	botAPI := api.NewTelegramAPI(os.Getenv("TOKEN"))
	tBot := bot.NewTelegramBot(&botAPI)

	///
	data, err := ioutil.ReadFile("telegram_test.key")
	if err != nil {
		panic(err)
	}
	certificate := models.InputFile(data)
	req := apiModels.SetWebhookRequest{
		Certificate: &certificate,
		URL:         query.NewParamString("https://1ad50cccab8b.ngrok.io"),
	}
	err = tBot.API().SetWebhook(context.Background(), req)
	if err != nil {
		panic(err)
	}
	status, err := tBot.API().GetWebhookInfo(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(status)

	/*err = tBot.API().DeleteWebhook(context.Background())
	if err != nil {
		panic(err)
	}

	status, err = tBot.API().GetWebhookInfo(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(status)*/
	///

	////////// LONG POLLING
	/*go tBot.WaitForUpdates(bot.NewUpdatesStrategyLongPolling(bot.LongPollingConfig{
		Timeout: 5,
		BotAPI:  &botAPI,
	}))*/

	////////// WEBHOOK
	go tBot.WaitForUpdates(bot.NewUpdatesStrategyWebhook(bot.WebhookConfig{
		ServeURL: "localhost:8443",
	}))

	updates, errs := tBot.Updates()
	for {
		select {
		case update, ok := <-updates:
			if !ok {
				fmt.Println("updates channel closed")
				return
			}
			///
			_, err := tBot.API().SendMessage(context.Background(), apiModels.SendMessageRequest{
				ChatID: update.Message.Chat.ID,
				Text:   fmt.Sprintf("Hello, %s!", update.Message.From.FirstName),
			})
			if err != nil {
				panic(err)
			}
			///
		case err, ok := <-errs:
			if !ok {
				fmt.Println("errors channel closed")
				return
			}
			fmt.Println(err)
		}
	}
}
