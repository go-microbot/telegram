package api

import "os"

const localAPIURL = "http://localhost:8081"

var localAPI TelegramAPI

func init() {
	localAPI = NewTelegramAPI(os.Getenv("TEST_BOT_TOKEN"))
	localAPI.url = localAPIURL
}
