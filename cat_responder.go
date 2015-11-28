package main

import (
	"errors"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/zhulik/margelet"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type CatResponder struct {
}

func downloadFromUrl(url string, output *os.File) error {
	response, err := http.Get(url)

	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Image url respose code != 200")
	}

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func (responder CatResponder) Response(bot margelet.MargeletAPI, message tgbotapi.Message) error {
	bot.Send(tgbotapi.NewChatAction(message.Chat.ID, tgbotapi.ChatUploadPhoto))

	file, _ := ioutil.TempFile("", "recognizer-")
	defer file.Close()
	defer os.Remove(file.Name())

	if err := downloadFromUrl("http://thecatapi.com/api/images/get?format=src&type=jpg", file); err != nil {
		return err
	}

	file.Seek(0, 0)
	bytes, _ := ioutil.ReadAll(file)

	msg := tgbotapi.NewPhotoUpload(message.Chat.ID, tgbotapi.FileBytes{"cat.jpg", bytes})
	msg.ChatID = message.Chat.ID
	msg.ReplyToMessageID = message.MessageID

	bot.Send(msg)

	return nil
}

func (responder CatResponder) HelpMessage() string {
	return "Send image with cat"
}
