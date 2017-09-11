package main

import (
	"encoding/json"
	"testing"
)

var updatesData = `{"ok":true,"result":[{"update_id":78084782,"message":{"message_id":91,"from":{"id":284557985,"is_bot":false,"first_name":"Levey","last_name":"S","language_code":"en"},"chat":{"id":284557985,"first_name":"Levey","last_name":"S","type":"private"},"date":1505113509,"text":"121"}},{"update_id":78084783,"message":{"message_id":92,"from":{"id":284557985,"is_bot":false,"first_name":"Levey","last_name":"S","language_code":"en"},"chat":{"id":284557985,"first_name":"Levey","last_name":"S","type":"private"},"date":1505113559,"text":"/stop","entities":[{"offset":0,"length":5,"type":"bot_command"}]}}]}`

func TestUpdatesType(t *testing.T) {
	updatesBytes := []byte(updatesData)

	dat := Updates{}
	if err := json.Unmarshal(updatesBytes, &dat); err != nil {
		t.Error("updates types unmarshal failed")
	}

	secondText := dat.Result[1].Message.Text

	if secondText != "/stop" {
		t.Error("decode text of message failed")
	}

	firstChatID := dat.Result[0].Message.Chat.ID

	if firstChatID != 284557985 {
		t.Error("decode chat id failed")
	}
}
