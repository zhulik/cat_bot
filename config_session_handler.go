package main

import (
	"errors"
	"github.com/zhulik/margelet"
)

type ConfigSessionHandler struct {
}

func (s ConfigSessionHandler) HandleSession(session margelet.Session) error {
	switch len(session.Responses()) {
	case 0:
		session.QuickForceReply("Would you like to receive a cat images sometimes? (yes/no)")
		return nil
	default:
		return s.handleResponse(session)
	}
}

func (responder ConfigSessionHandler) HelpMessage() string {
	return "Configure bot"
}

func (s ConfigSessionHandler) handleResponse(session margelet.Session) error {
	if session.Message().Text != "yes" && session.Message().Text != "no" {
		session.QuickForceReply("Sorry, i can't understand, type yes or no")
		return errors.New("Waiting for yes or no message")
	}

	if session.Message().Text == "yes" {
		session.GetConfigRepository().Set(session.Message().Chat.ID, "yes")
		session.QuickReply("Ok, i will send you a cat sometimes")
		session.Finish()
		return nil
	} else {
		session.GetConfigRepository().Set(session.Message().Chat.ID, "no")
		session.QuickReply("Ok, i will not send you a cat sometimes")
		session.Finish()
		return nil
	}
}

func (s ConfigSessionHandler) CancelSession(session margelet.Session) {
}
