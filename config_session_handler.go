package main

import (
	"errors"

	"github.com/zhulik/margelet"
	"gopkg.in/telegram-bot-api.v2"
)

type ConfigSessionHandler struct {
}

func (session ConfigSessionHandler) HandleSession(bot margelet.MargeletAPI, message tgbotapi.Message, responses []tgbotapi.Message) (bool, error) {
	switch len(responses) {
	case 0:
		msg := tgbotapi.MessageConfig{Text: "Would you like to receive a cat images sometimes? (yes/no)"}
		session.forceReply(bot, message, msg)
		return false, nil
	default:
		return session.handleResponse(bot, message)
	}
}

func (responder ConfigSessionHandler) HelpMessage() string {
	return "Configure bot"
}

func (session ConfigSessionHandler) handleResponse(bot margelet.MargeletAPI, message tgbotapi.Message) (bool, error) {
	if message.Text != "yes" && message.Text != "no" {
		msg := tgbotapi.MessageConfig{Text: "Sorry, i can't understand, type yes or no"}
		session.forceReply(bot, message, msg)
		return false, errors.New("Waiting for yes or no message")
	}

	if message.Text == "yes" {
		bot.GetConfigRepository().Set(message.Chat.ID, "yes")
		msg := tgbotapi.MessageConfig{Text: "Ok, i will send you a cat sometimes"}
		session.reply(bot, message, msg)
		return true, nil
	} else {
		bot.GetConfigRepository().Set(message.Chat.ID, "no")
		msg := tgbotapi.MessageConfig{Text: "Ok, i will not send you a cat sometimes"}
		session.reply(bot, message, msg)
		return true, nil
	}
}

func (session ConfigSessionHandler) forceReply(bot margelet.MargeletAPI, message tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.ReplyMarkup = tgbotapi.ForceReply{true, true}
	session.reply(bot, message, msg)
}

func (session ConfigSessionHandler) reply(bot margelet.MargeletAPI, message tgbotapi.Message, msg tgbotapi.MessageConfig) {
	msg.ChatID = message.Chat.ID
	msg.ReplyToMessageID = message.MessageID
	bot.Send(msg)
}

func (session ConfigSessionHandler) CancelSession(bot margelet.MargeletAPI, message tgbotapi.Message, responses []tgbotapi.Message) {
}
