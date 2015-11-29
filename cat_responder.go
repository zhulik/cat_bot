package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"github.com/zhulik/margelet"
)

type CatResponder struct {
}

func (responder CatResponder) Response(margelet margelet.MargeletAPI, message tgbotapi.Message) error {
	margelet.Send(tgbotapi.NewChatAction(message.Chat.ID, tgbotapi.ChatUploadPhoto))

	bytes, err := downloadCat()

	if err != nil {
		return err
	}

	msg := tgbotapi.NewPhotoUpload(message.Chat.ID, tgbotapi.FileBytes{"cat.jpg", bytes})
	msg.ChatID = message.Chat.ID
	msg.ReplyToMessageID = message.MessageID

	margelet.Send(msg)

	return nil
}

func (responder CatResponder) HelpMessage() string {
	return "Send image with cat"
}
