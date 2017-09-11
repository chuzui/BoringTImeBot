package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	BASEURL = "https://api.telegram.org/bot"
	TIMEOUT = "100"
)

type Bot struct {
	Token string
}

func NewBot(token string) *Bot {
	return &Bot{Token: token}
}

func (bot *Bot) GetUpdates(updateID int) (*Updates, error) {
	url := fmt.Sprintf(BASEURL+"%s/getUpdates?offset=%s&timeout=%s", bot.Token, strconv.Itoa(updateID), TIMEOUT)

	rep, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer rep.Body.Close()

	updates := Updates{}
	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &updates)
	if err != nil {
		return nil, err
	}

	return &updates, nil
}

func (bot *Bot) SendMessage(chatID int, message string) error {
	url := fmt.Sprintf(BASEURL+"%s/sendMessage?text=%s&chat_id=%s", bot.Token, message, strconv.Itoa(chatID))

	_, err := http.Get(url)
	if err != nil {
		return err
	}

	return nil
}
