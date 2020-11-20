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
	groupChatIDCtxKey             = "group_chat_id"
	newGroupChatTitleCtxKey       = "new_group_chat_title"
	newGroupChatDescriptionCtxKey = "new_group_chat_description"
	sentMessageIDCtxKey           = "sent_message_id"
	botCommandsCtxKey             = "bot_commands"
	botUserIDCtxKey               = "bot_user_id"
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
		sleep       int
	}{
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
			sleep:       1,
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
	}
	for i := range testCases {
		tc := &testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctx = tc.testHandler.Test(ctx, t)
			time.Sleep(time.Second * time.Duration(tc.sleep))
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
