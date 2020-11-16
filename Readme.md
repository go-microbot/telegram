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
- [TODO](#todo)

## Getting started

### Installation
Download the latest version of the Bot API.

```bash
go get -u github.com/go-microbot/telegram
```

### Bot token
Create your own bot. Follow the [Official guide](https://core.telegram.org/bots/api#authorizing-your-bot).

---

## Update Strategies
There are two mutually exclusive ways of receiving updates for your bot — the `Long Polling` on one hand and `Webhooks` on the other. Incoming updates are stored on the server until the bot receives them either way, but they will not be kept longer than 24 hours.

### Long Polling
```go
package main
```

### Webhook
```go
package main
```

> Ports currently supported for Webhooks: **443**, **80**, **88**, **8443**.

---

## Example
See the [examples](./examples) folder to get all available examples.

---

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

---

## TODO
The bot is **isn't finished yet**. The main goal is to implement all available methods from [Telegram documentation](https://core.telegram.org/bots/api#available-methods).


Implementation     | Test coverage  | Method                          | Docs                                                               |
-----------------  | -------------  | ------------------------------- | -----------------------------------------------------------------  |
:x:                | :x:            | logout                          | https://core.telegram.org/bots/api#logout                          |
:x:                | :x:            | close                           | https://core.telegram.org/bots/api#close                           |
:x:                | :x:            | sendMessage                     | https://core.telegram.org/bots/api#sendMessage                     |
:x:                | :x:            | forwardMessage                  | https://core.telegram.org/bots/api#forwardMessage                  |
:x:                | :x:            | copyMessage                     | https://core.telegram.org/bots/api#copyMessage                     |
:x:                | :x:            | sendPhoto                       | https://core.telegram.org/bots/api#sendPhoto                       |
:x:                | :x:            | sendAudio                       | https://core.telegram.org/bots/api#sendAudio                       |
:x:                | :x:            | sendDocument                    | https://core.telegram.org/bots/api#sendDocument                    |
:x:                | :x:            | sendVideo                       | https://core.telegram.org/bots/api#sendVideo                       |
:x:                | :x:            | sendAnimation                   | https://core.telegram.org/bots/api#sendAnimation                   |
:x:                | :x:            | sendVoice                       | https://core.telegram.org/bots/api#sendVoice                       |
:x:                | :x:            | sendVideoNote                   | https://core.telegram.org/bots/api#sendVideoNote                   |
:x:                | :x:            | sendMediaGroup                  | https://core.telegram.org/bots/api#sendMediaGroup                  |
:x:                | :x:            | sendLocation                    | https://core.telegram.org/bots/api#sendLocation                    |
:x:                | :x:            | editMessageLiveLocation         | https://core.telegram.org/bots/api#editMessageLiveLocation         |
:x:                | :x:            | stopMessageLiveLocation         | https://core.telegram.org/bots/api#stopMessageLiveLocation         |
:x:                | :x:            | sendVenue                       | https://core.telegram.org/bots/api#sendVenue                       |
:x:                | :x:            | sendContact                     | https://core.telegram.org/bots/api#sendContact                     |
:x:                | :x:            | sendPoll                        | https://core.telegram.org/bots/api#sendPoll                        |
:x:                | :x:            | sendDice                        | https://core.telegram.org/bots/api#sendDice                        |
:x:                | :x:            | sendChatAction                  | https://core.telegram.org/bots/api#sendChatAction                  |
:x:                | :x:            | getUserProfilePhotos            | https://core.telegram.org/bots/api#getUserProfilePhotos            |
:x:                | :x:            | getFile                         | https://core.telegram.org/bots/api#getFile                         |
:x:                | :x:            | kickChatMember                  | https://core.telegram.org/bots/api#kickChatMember                  |
:x:                | :x:            | unbanChatMember                 | https://core.telegram.org/bots/api#unbanChatMember                 |
:x:                | :x:            | restrictChatMember              | https://core.telegram.org/bots/api#restrictChatMember              |
:x:                | :x:            | promoteChatMember               | https://core.telegram.org/bots/api#promoteChatMember               |
:x:                | :x:            | setChatAdministratorCustomTitle | https://core.telegram.org/bots/api#setChatAdministratorCustomTitle |
:x:                | :x:            | setChatPermissions              | https://core.telegram.org/bots/api#setChatPermissions              |
:x:                | :x:            | exportChatInviteLink            | https://core.telegram.org/bots/api#exportChatInviteLink            |
:x:                | :x:            | setChatPhoto                    | https://core.telegram.org/bots/api#setChatPhoto                    |
:x:                | :x:            | deleteChatPhoto                 | https://core.telegram.org/bots/api#deleteChatPhoto                 |
:x:                | :x:            | setChatTitle                    | https://core.telegram.org/bots/api#setChatTitle                    |
:x:                | :x:            | setChatDescription              | https://core.telegram.org/bots/api#setChatDescription              |
:x:                | :x:            | pinChatMessage                  | https://core.telegram.org/bots/api#pinChatMessage                  |
:x:                | :x:            | unpinChatMessage                | https://core.telegram.org/bots/api#unpinChatMessage                |
:x:                | :x:            | unpinAllChatMessages            | https://core.telegram.org/bots/api#unpinAllChatMessages            |
:x:                | :x:            | leaveChat                       | https://core.telegram.org/bots/api#leaveChat                       |
:x:                | :x:            | getChat                         | https://core.telegram.org/bots/api#getChat                         |
:x:                | :x:            | getChatAdministrators           | https://core.telegram.org/bots/api#getChatAdministrators           |
:x:                | :x:            | getChatMembersCount             | https://core.telegram.org/bots/api#getChatMembersCount             |
:x:                | :x:            | getChatMember                   | https://core.telegram.org/bots/api#getChatMember                   |
:x:                | :x:            | setChatStickerSet               | https://core.telegram.org/bots/api#setChatStickerSet               |
:x:                | :x:            | deleteChatStickerSet            | https://core.telegram.org/bots/api#deleteChatStickerSet            |
:x:                | :x:            | answerCallbackQuery             | https://core.telegram.org/bots/api#answerCallbackQuery             |
:x:                | :x:            | setMyCommands                   | https://core.telegram.org/bots/api#setMyCommands                   |
:x:                | :x:            | getMyCommands                   | https://core.telegram.org/bots/api#getMyCommands                   |
:x:                | :x:            | editMessageText                 | https://core.telegram.org/bots/api#editMessageText                 |
:x:                | :x:            | editMessageCaption              | https://core.telegram.org/bots/api#editMessageCaption              |
:x:                | :x:            | editMessageMedia                | https://core.telegram.org/bots/api#editMessageMedia                |
:x:                | :x:            | editMessageReplyMarkup          | https://core.telegram.org/bots/api#editMessageReplyMarkup          |
:x:                | :x:            | stopPoll                        | https://core.telegram.org/bots/api#stopPoll                        |
:x:                | :x:            | deleteMessage                   | https://core.telegram.org/bots/api#deleteMessage                   |
:x:                | :x:            | sendSticker                     | https://core.telegram.org/bots/api#sendSticker                     |
:x:                | :x:            | getStickerSet                   | https://core.telegram.org/bots/api#getStickerSet                   |
:x:                | :x:            | uploadStickerFile               | https://core.telegram.org/bots/api#uploadStickerFile               |
:x:                | :x:            | createNewStickerSet             | https://core.telegram.org/bots/api#createNewStickerSet             |
:x:                | :x:            | addStickerToSet                 | https://core.telegram.org/bots/api#addStickerToSet                 |
:x:                | :x:            | setStickerPositionInSet         | https://core.telegram.org/bots/api#setStickerPositionInSet         |
:x:                | :x:            | deleteStickerFromSet            | https://core.telegram.org/bots/api#deleteStickerFromSet            |
:x:                | :x:            | setStickerSetThumb              | https://core.telegram.org/bots/api#setStickerSetThumb              |
:x:                | :x:            | answerInlineQuery               | https://core.telegram.org/bots/api#answerInlineQuery               |
:x:                | :x:            | sendInvoice                     | https://core.telegram.org/bots/api#sendInvoice                     |
:x:                | :x:            | answerShippingQuery             | https://core.telegram.org/bots/api#answerShippingQuery             |
:x:                | :x:            | answerPreCheckoutQuery          | https://core.telegram.org/bots/api#answerPreCheckoutQuery          |
:x:                | :x:            | setPassportDataErrors           | https://core.telegram.org/bots/api#setPassportDataErrors           |
:x:                | :x:            | sendGame                        | https://core.telegram.org/bots/api#sendGame                        |
:x:                | :x:            | setGameScore                    | https://core.telegram.org/bots/api#setGameScore                    |
:x:                | :x:            | getGameHighScores               | https://core.telegram.org/bots/api#getGameHighScores               |
:x:                | :x:            | getGameHighScores               | https://core.telegram.org/bots/api#getGameHighScores               |