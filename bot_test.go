package main

import (
	"testing"
)

const (
	TESTTOKEN = "275517717:AAE6F5k4LDRB6yJr9uLityljNioDiapnbBI"
	CHATID    = 284557985
)

func TestGetUpdates(t *testing.T) {
	bot := NewBot(TESTTOKEN)

	updates, err := bot.GetUpdates(0)
	if err != nil {
		t.Errorf("GetUpdates get error %s.", err)
		return
	}

	if updates.Ok != true {
		t.Error("Ok of updates is false.")
	}
}

func TestSendMessage(t *testing.T) {
	bot := NewBot(TESTTOKEN)

	err := bot.SendMessage(CHATID, "Hello world.")
	if err != nil {
		t.Error("Send message error.")
	}
}
