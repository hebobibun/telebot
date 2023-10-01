package bot

import (
	"log"
	"os"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// "github.com/spf13/viper"
	"telebot/handlers"
)

type Bot struct {
	API *tgbotapi.BotAPI
}

func InitializeBot() *Bot {
	// // Configure Viper to read from the .env file
	// viper.SetConfigFile(".env")

	// // Check if the .env file exists
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalf("Error reading .env file: %v", err)
	// }

	// // Retrieve the API token using Viper
	// apiToken := viper.GetString("TELEGRAM_BOT_API_TOKEN")
	// if apiToken == "" {
	// 	log.Fatal("API token not found in .env file. Please make sure the .env file contains TELEGRAM_BOT_API_TOKEN.")
	// }

	// Retrieve the API token using os.Getenv
	apiToken := os.Getenv("TELEGRAM_BOT_API_TOKEN")
	if apiToken == "" {
		log.Fatal("API token not found in environment variables. Please make sure you have set TELEGRAM_BOT_API_TOKEN.")
	}

	// Initialize the bot with the retrieved API token
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		log.Fatal(err)
	}

	return &Bot{API: bot}
}

func (b *Bot) StartListening() {
	// Set up an update configuration
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Get updates from the Telegram API
	updates, err := b.API.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	// Process incoming updates
	for update := range updates {
		handlers.HandleUpdate(b.API, update)
	}
}
