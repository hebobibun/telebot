package main

import (
	"log"
	"telebot/bot"
)

func main() {
	log.Println("Bot is getting started...")

	// Initialize the Telegram bot
	b := bot.InitializeBot()

	// Start listening for updates
	b.StartListening()
}
