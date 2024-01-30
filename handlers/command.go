package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil { // Ignore non-message updates
		return
	}

	// Respond to a specific command (e.g., /start)
	if update.Message.IsCommand() {
		command := update.Message.Command()
		switch command {
		case "start":
			SendWelcomeMessage(bot, update.Message.Chat.ID)
		case "hello":
			SendHelloMessage(bot, update.Message.Chat.ID)
		}
	}
}

func SendWelcomeMessage(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Welcome to the Telegram Bot! You have started the bot. You can also say /hello to greet me.")
	log.Println("Welcoming :", bot.Self.FirstName)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func SendHelloMessage(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Hello, I'm your Telegram bot! You said hello.")
	log.Println("sent hello to", bot.Self.FirstName)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
