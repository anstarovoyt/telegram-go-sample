package main

import (
	"log"
	"os"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	bot, err := tg.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Failed to create new bot: %v", err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tg.NewMessage(update.Message.Chat.ID, "Hello, "+update.Message.From.FirstName+"!")
		msg.ReplyToMessageID = update.Message.MessageID

		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}
}
