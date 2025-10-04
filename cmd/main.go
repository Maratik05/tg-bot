package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Maratik05/tg-bot/internal/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Not found file .env")
	}
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}
		service := product.NewService()
		fmt.Print("nan")
		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		case "list":
			listCommand(bot, update.Message, service)
		default:
			defaultResponse(bot, update.Message)
		}

	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help")
	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService *product.Service) {
	products := productService.List()
	outputMsg := "List products\n\n"

	for _, product := range products {
		outputMsg += product.Title
		outputMsg += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
	bot.Send(msg)
}
func defaultResponse(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID
	bot.Send(msg)
}
