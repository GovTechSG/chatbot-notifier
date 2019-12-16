package telegram

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// SendMsg send message to telegram bot
func SendMsg(t string, c int64, m string) {
	bot, err := tgbotapi.NewBotAPI(t)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	msg := tgbotapi.NewMessage(c, m)
	bot.Send(msg)
}