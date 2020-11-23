package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const localAPIURL = "http://localhost:8081"

const (
	webhookURLCtxKey              = "webhook_url"
	chatIDCtxKey                  = "chat_id"
	existingPhotoIDCtxKey         = "existing_photo_id"
	existingPhotoUIDCtxKey        = "existing_photo_unique_id"
	photoToUploadURLCtxKey        = "photo_to_upload_url"
	videoToUploadURLCtxKey        = "video_to_upload_url"
	documentToUploadURLCtxKey     = "document_to_upload_url"
	audioToUploadURLCtxKey        = "audio_to_upload_url"
	groupChatIDCtxKey             = "group_chat_id"
	newGroupChatTitleCtxKey       = "new_group_chat_title"
	newGroupChatDescriptionCtxKey = "new_group_chat_description"
	sentMessageIDCtxKey           = "sent_message_id"
	botCommandsCtxKey             = "bot_commands"
	botUserIDCtxKey               = "bot_user_id"
	botUsernameCtxKey             = "bot_user_name"
	adminUserIDCtxKey             = "admin_user_id"
	stickerSetNameCtxKey          = "sticker_set_name"
	pollQuestionCtxKey            = "poll_question"
	pollMessageIDCtxKey           = "poll_message_id"
	gameShortNameCtxKey           = "game_short_name"
	gameScoreCtxKey               = "game_score"
	gameMessageIDCtxKey           = "game_message_id"
	uploadedStickerFileIDCtxKey   = "uploaded_sticker_file_id"
)

var (
	localAPI TelegramAPI
)

type Testable interface {
	Test(ctx context.Context, t *testing.T) context.Context
}

// TestDataKey
type TestDataKey interface{}

func TestMain(m *testing.M) {
	localAPI = NewTelegramAPI(os.Getenv("TEST_BOT_TOKEN"))
	localAPI.url = localAPIURL
	rand.Seed(time.Now().Unix())

	os.Exit(m.Run())
}

func Test_NewTelegramAPI(t *testing.T) {
	tAPI := NewTelegramAPI("123")
	require.NotNil(t, tAPI)
	require.Equal(t, "123", tAPI.token)
	require.NotNil(t, tAPI.client)
	require.Equal(t, baseURL, tAPI.url)
}

func TestTelegramAPI_Integration(t *testing.T) {
	// prepare context data.
	wd, err := os.Getwd()
	require.NoError(t, err)
	cfgMap, err := parseTestConfig(path.Join(wd, "./test_data/config.json"))
	require.NoError(t, err)
	ctx := context.Background()
	for k, v := range cfgMap {
		if val, ok := v.(float64); ok {
			v = int64(val)
		}
		ctx = context.WithValue(ctx, TestDataKey(k), v)
	}

	testCases := []struct {
		name        string
		testHandler Testable
	}{
		{
			name:        "close",
			testHandler: close{},
		},
		{
			name:        "logOut",
			testHandler: logOut{},
		},
		{
			name:        "leaveChat",
			testHandler: leaveChat{},
		},
		{
			name:        "getMe",
			testHandler: getMe{},
		},
		{
			name:        "setWebhook",
			testHandler: setWebhook{},
		},
		{
			name:        "getWebhookInfo",
			testHandler: getWebhookInfo{},
		},
		{
			name:        "deleteWebhook",
			testHandler: deleteWebhook{},
		},
		// should be always after webhook delete case.
		{
			name:        "getUpdates",
			testHandler: getUpdates{},
		},
		{
			name:        "sendMessage",
			testHandler: sendMessage{},
		},
		{
			name:        "editMessageLiveLocation",
			testHandler: editMessageLiveLocation{},
		},
		{
			name:        "stopMessageLiveLocation",
			testHandler: stopMessageLiveLocation{},
		},
		{
			name:        "sendPhoto",
			testHandler: sendPhoto{},
		},
		{
			name:        "setChatPermissions",
			testHandler: setChatPermissions{},
		},
		{
			name:        "setChatPhoto",
			testHandler: setChatPhoto{},
		},
		{
			name:        "setChatTitle",
			testHandler: setChatTitle{},
		},
		{
			name:        "setChatDescription",
			testHandler: setChatDescription{},
		},
		{
			name:        "getChat",
			testHandler: getChat{},
		},
		{
			name:        "forwardMessage",
			testHandler: forwardMessage{},
		},
		{
			name:        "setMyCommands",
			testHandler: setMyCommands{},
		},
		{
			name:        "getMyCommands",
			testHandler: getMyCommands{},
		},
		{
			name:        "deleteChatPhoto",
			testHandler: deleteChatPhoto{},
		},
		{
			name:        "sendLocation",
			testHandler: sendLocation{},
		},
		{
			name:        "getFile",
			testHandler: getFile{},
		},
		{
			name:        "pinChatMessage",
			testHandler: pinChatMessage{},
		},
		{
			name:        "unpinChatMessage",
			testHandler: unpinChatMessage{},
		},
		{
			name:        "unpinAllChatMessages",
			testHandler: unpinAllChatMessages{},
		},
		{
			name:        "getChatMember",
			testHandler: getChatMember{},
		},
		{
			name:        "exportChatInviteLink",
			testHandler: exportChatInviteLink{},
		},
		{
			name:        "getChatAdministrators",
			testHandler: getChatAdministrators{},
		},
		{
			name:        "setChatAdministratorCustomTitle",
			testHandler: setChatAdministratorCustomTitle{},
		},
		{
			name:        "getChatMembersCount",
			testHandler: getChatMembersCount{},
		},
		{
			name:        "sendAudio",
			testHandler: sendAudio{},
		},
		{
			name:        "sendDocument",
			testHandler: sendDocument{},
		},
		{
			name:        "sendVideo",
			testHandler: sendVideo{},
		},
		{
			name:        "sendAnimation",
			testHandler: sendAnimation{},
		},
		{
			name:        "sendVoice",
			testHandler: sendVoice{},
		},
		{
			name:        "copyMessage",
			testHandler: copyMessage{},
		},
		{
			name:        "deleteMessage",
			testHandler: deleteMessage{},
		},
		{
			name:        "sendChatAction",
			testHandler: sendChatAction{},
		},
		{
			name:        "getUserProfilePhotos",
			testHandler: getUserProfilePhotos{},
		},
		{
			name:        "createNewStickerSet",
			testHandler: createNewStickerSet{},
		},
		{
			name:        "setChatStickerSet",
			testHandler: setChatStickerSet{},
		},
		{
			name:        "getStickerSet",
			testHandler: getStickerSet{},
		},
		{
			name:        "sendVideoNote",
			testHandler: sendVideoNote{},
		},
		{
			name:        "deleteChatStickerSet",
			testHandler: deleteChatStickerSet{},
		},
		{
			name:        "sendMediaGroup",
			testHandler: sendMediaGroup{},
		},
		{
			name:        "sendDice",
			testHandler: sendDice{},
		},
		{
			name:        "sendVenue",
			testHandler: sendVenue{},
		},
		{
			name:        "sendContact",
			testHandler: sendContact{},
		},
		{
			name:        "sendPoll",
			testHandler: sendPoll{},
		},
		{
			name:        "stopPoll",
			testHandler: stopPoll{},
		},
		{
			name:        "kickChatMember",
			testHandler: kickChatMember{},
		},
		{
			name:        "unbanChatMember",
			testHandler: unbanChatMember{},
		},
		{
			name:        "restrictChatMember",
			testHandler: restrictChatMember{},
		},
		{
			name:        "promoteChatMember",
			testHandler: promoteChatMember{},
		},
		{
			name:        "sendGame",
			testHandler: sendGame{},
		},
		{
			name:        "setGameScore",
			testHandler: setGameScore{},
		},
		{
			name:        "getGameHighScores",
			testHandler: getGameHighScores{},
		},
		{
			name:        "uploadStickerFile",
			testHandler: uploadStickerFile{},
		},
		{
			name:        "addStickerToSet",
			testHandler: addStickerToSet{},
		},
	}
	for i := range testCases {
		tc := &testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctx = tc.testHandler.Test(ctx, t)
			// When sending messages inside a particular chat,
			// avoid sending more than one message per second.
			// We may allow short bursts that go over this limit,
			// but eventually you'll begin receiving 429 errors.
			// https://core.telegram.org/bots/faq#my-bot-is-hitting-limits-how-do-i-avoid-this.
			time.Sleep(time.Second)
		})
	}
}

func parseTestConfig(cfgPath string) (map[string]interface{}, error) {
	var cfgMap map[string]interface{}

	data, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &cfgMap); err != nil {
		return nil, err
	}

	return cfgMap, nil
}
