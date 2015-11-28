package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"github.com/zhulik/margelet"
	"errors"
)

type ConfigSession struct {
}

func (session ConfigSession) HandleResponse(bot margelet.MargeletAPI, message tgbotapi.Message, responses []string) (bool, error) {
	switch len(responses) {
	case 0:
		msg := tgbotapi.MessageConfig{Text: "Would you like to receive a cat images? (yes/no)"}
		msg.ReplyMarkup = tgbotapi.ForceReply{true, true}
		msg.ChatID = message.Chat.ID
		msg.ReplyToMessageID = message.MessageID
		bot.Send(msg)
		return false, nil
	case 1:
		if message.Text != "yes" && message.Text != "no" {
			msg := tgbotapi.MessageConfig{Text: "Sorry, i can't understand, type yes or no"}
			msg.ReplyMarkup = tgbotapi.ForceReply{true, true}
			msg.ChatID = message.Chat.ID
			msg.ReplyToMessageID = message.MessageID
			bot.Send(msg)
			return false, errors.New("Waiting for yes or no message")
		}

		if message.Text == "yes" {
			bot.GetConfigRepository().Set(message.Chat.ID, "yes")
			msg := tgbotapi.MessageConfig{Text: "Ok, i will send you a cat sometimes"}
			msg.ReplyMarkup = tgbotapi.ForceReply{true, true}
			msg.ChatID = message.Chat.ID
			msg.ReplyToMessageID = message.MessageID
			bot.Send(msg)
			return true, nil
		} else {
			bot.GetConfigRepository().Set(message.Chat.ID, "no")
			msg := tgbotapi.MessageConfig{Text: "Ok, i will not send you a cat sometimes"}
			msg.ReplyMarkup = tgbotapi.ForceReply{true, true}
			msg.ChatID = message.Chat.ID
			msg.ReplyToMessageID = message.MessageID
			bot.Send(msg)
			return true, nil
		}
	}
	return false, nil
}

func (responder ConfigSession) HelpMessage() string {
	return "Configure bot"
}
