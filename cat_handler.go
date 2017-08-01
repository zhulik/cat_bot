package main

import (
	"github.com/zhulik/margelet"
	"gopkg.in/telegram-bot-api.v4"
)

type CatHandler struct {
}

func (handler CatHandler) HandleCommand(message margelet.Message) error {

	message.Bot().Send(tgbotapi.NewChatAction(message.Message().Chat.ID, tgbotapi.ChatUploadPhoto))

	bytes, err := downloadCat()

	if err != nil {
		return err
	}

	msg := tgbotapi.NewPhotoUpload(message.Message().Chat.ID, tgbotapi.FileBytes{"cat.jpg", bytes})
	msg.ChatID = message.Message().Chat.ID
	//msg.ReplyToMessageID = message.Message().MessageID

	message.Bot().Send(msg)


	return nil
}

func (handler CatHandler) HelpMessage() string {
	return "Send image with cat"
}
