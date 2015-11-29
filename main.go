package main

import (
	"github.com/zhulik/margelet"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	token := kingpin.Flag("token", "Telegram Bot token").Required().Short('t').String()
	kingpin.Parse()

	margelet, err := margelet.NewMargelet("emergency_kittens", "127.0.0.1:6379", "", 3, *token, false)

	if err != nil {
		panic(err)
	}

	margelet.AddCommandHandler("/cat", CatResponder{})
	margelet.AddSessionHandler("/start", ConfigSession{})



	go randomCatSender(margelet)

	margelet.Run()
}
