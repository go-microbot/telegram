package api

import "os"

var localAPI TelegramAPI

func init() {
	localAPI = NewTelegramAPI(os.Getenv("TEST_BOT_TOKEN"))
	localAPI.url = "http://localhost:8081"
}
