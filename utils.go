package main

import (
	"errors"
	"github.com/zhulik/margelet"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	catURL = "http://thecatapi.com/api/images/get?format=src&type=jpg"
)

func downloadFromUrl(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return []byte{}, errors.New("Image url respose code != 200")
	}

	return ioutil.ReadAll(response.Body)
}

func downloadCat() ([]byte, error) {
	return downloadFromUrl(catURL)
}

func sendCat(chatID int64, bot *margelet.Margelet) {
	if bot.ChatConfigRepository.Get(chatID) == "yes" {
		if bytes, err := downloadCat(); err == nil {

			msg := tgbotapi.NewPhotoUpload(chatID, tgbotapi.FileBytes{"cat.jpg", bytes})
			msg.ChatID = chatID

			bot.Send(msg)
		}
	}
}

func randomCatSender(bot *margelet.Margelet) {
	for {
		for _, chatID := range bot.ChatRepository.All() {
			go sendCat(chatID, bot)
		}
		time.Sleep(5 * time.Minute)
	}
}
