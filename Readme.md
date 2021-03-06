# [GO-MICROBOT] Telegram

## This is fully documented and easy to use go-bindings for working with the [Telegram](https://telegram.org/) Bot API.

[![Coverage Status](https://coveralls.io/repos/github/go-microbot/telegram/badge.svg)](https://coveralls.io/github/go-microbot/telegram)

<p align="center">
  <img height="120" src="./.github/assets/gopher.png">
  <img height="80" src="./.github/assets/heart.png">
  <img height="120" src="./.github/assets/telegram.png">
</p>

> Please read the [Official Telegram Bot API Documentation](https://core.telegram.org/bots/api) before starting.

# Guides

- [Getting started](#getting-started)
  - [Installation](#installation)
  - [Create bot token](#bot-token)
- [Update Strategies](#update-strategies)
  - [Long Polling](#long-polling)
  - [Webhook](#webhook)
- [Example](#example)
- [Test](#test)
  - [Start Docker Containers](#start-docker-containers)
  - [Local testing](#run-tests-locally)
  - [Lint](#run-linter)

## Getting started

### Installation
Download the latest version of the Bot API.

```bash
go get -u github.com/go-microbot/telegram
```

### Bot token
Create your own bot. Follow the [Official guide](https://core.telegram.org/bots/api#authorizing-your-bot).

## Update Strategies
There are two mutually exclusive ways of receiving updates for your bot — the `Long Polling` on one hand and `Webhooks` on the other. Incoming updates are stored on the server until the bot receives them either way, but they will not be kept longer than 24 hours.

### Long Polling
```go
package main

import (
	"context"
	"fmt"

	"github.com/go-microbot/telegram/api"
	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/bot"
	"github.com/go-microbot/telegram/query"
)

const telegramBotToken = "<PASTE_YOUR_TOKEN_HERE>"

func main() {
	// init Bot API with token.
	botAPI := api.NewTelegramAPI(telegramBotToken)

	// create Bot instance.
	myBot := bot.NewTelegramBot(&botAPI)

	// delete webhook (if it was using before).
	if err := myBot.API().DeleteWebhook(context.Background()); err != nil {
		fmt.Printf("could not remove webhook: %v", err)
	}

	// start long polling.
	go myBot.WaitForUpdates(bot.NewUpdatesStrategyLongPolling(bot.LongPollingConfig{
		Timeout: 10,
		BotAPI:  &botAPI,
	}))

	// listen Bot's updates.
	updates, errs := myBot.Updates()
	for {
		select {
		case update, ok := <-updates:
			if !ok {
				fmt.Println("updates channel closed")
				return
			}

			// reply "hello" message.
			_, err := myBot.API().SendMessage(context.Background(), apiModels.SendMessageRequest{
				ChatID:           query.NewParamAny(update.Message.Chat.ID),
				Text:             fmt.Sprintf("Hello, %s!", update.Message.From.Username),
				ReplyToMessageID: &update.Message.ID,
			})
			if err != nil {
				panic(err)
			}
		case err, ok := <-errs:
			if !ok {
				fmt.Println("errors channel closed")
				return
			}
			fmt.Println(err)
		}
	}
}
```

### Webhook
To use a self-signed certificate, you need to create your [public key certificate](https://core.telegram.org/bots/self-signed). 

```go
package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/go-microbot/telegram/api"
	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/bot"
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
)

const telegramBotToken = "<PASTE_YOUR_TOKEN_HERE>"

func main() {
	// init Bot API with token.
	botAPI := api.NewTelegramAPI(telegramBotToken)

	// create Bot instance.
	myBot := bot.NewTelegramBot(&botAPI)

	// read certificate data.
	data, err := ioutil.ReadFile("telegram_test.key")
	if err != nil {
		panic(err)
	}

	// set webhook.
	req := apiModels.SetWebhookRequest{
		Certificate: form.NewPartText(string(data)),
		URL:         query.NewParamString("https://53ec7fc0c840.ngrok.io"), // ngrok, you need to use your server URL.
	}
	err = myBot.API().SetWebhook(context.Background(), req)
	if err != nil {
		panic(err)
	}

	// start listening.
	go myBot.WaitForUpdates(bot.NewUpdatesStrategyWebhook(bot.WebhookConfig{
		ServeURL: "localhost:8443", // server to catch Telegram requests.
	}))

	// listen Bot's updates.
	updates, errs := myBot.Updates()
	for {
		select {
		case update, ok := <-updates:
			if !ok {
				fmt.Println("updates channel closed")
				return
			}

			// reply "hello" message.
			_, err := myBot.API().SendMessage(context.Background(), apiModels.SendMessageRequest{
				ChatID:           query.NewParamAny(update.Message.Chat.ID),
				Text:             fmt.Sprintf("Hello, %s!", update.Message.From.Username),
				ReplyToMessageID: &update.Message.ID,
			})
			if err != nil {
				panic(err)
			}
		case err, ok := <-errs:
			if !ok {
				fmt.Println("errors channel closed")
				return
			}
			fmt.Println(err)
		}
	}
}
```

> Ports currently supported for Webhooks: **443**, **80**, **88**, **8443**.

## Example
See the [examples](./examples) folder to get all available examples.

## Test

**NOTE**
> Download and install [Docker](https://docs.docker.com/) before running tests locally.

### Start Docker Containers
Create your [Telegram Application](https://core.telegram.org/api/obtaining_api_id).

The mandatory options are `--api-id` and `--api-hash`. You must obtain your own `api_id` and `api_hash` as described in https://core.telegram.org/api/obtaining_api_id and specify them using the `--api-id` and `--api-hash` options or the `TELEGRAM_API_ID` and `TELEGRAM_API_HASH` environment variables.

Use the following command:
```bash
docker run --publish 8081:8081 -it --rm -d --name telegram-bot-api huntechio/telegram-bot-api:master-7cf91e4 --api-id=${TEST_API_ID} --api-hash=${TEST_API_HASH}
```

Or use [Makefile](./Makefile)'s `start_images` command:
```bash
make start_images
```

### Run tests locally
To run tests locally please specify the `TEST_BOT_TOKEN` env variable. It should contains your bot token.

Use the following command:
```bash
go test -p 1
```

Or use [Makefile](./Makefile)'s `test` command:
```bash
make test
```

### Run linter
Use the following commands:
```bash
golangci-lint cache clean
golangci-lint run --config .golangci.yml --timeout=5m
```

Or use [Makefile](./Makefile)'s `lint` command:
```bash
make lint
```
