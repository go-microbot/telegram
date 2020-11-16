package api

const localAPIURL = "http://localhost:8081"

var localAPI TelegramAPI

func init() {
	localAPI = NewTelegramAPI( /*os.Getenv("TEST_BOT_TOKEN")*/ "1256583982:AAHoS3RanoCsXtKhNJCQSOftJXWSRJnLg2o")
	localAPI.url = localAPIURL
}
