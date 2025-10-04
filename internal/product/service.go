package product

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Service struct {
	bot *tgbotapi.BotAPI
}

func NewService(bot *tgbotapi.BotAPI) *Service {
	return &Service{
		bot: bot,
	}
}

func (s *Service) list() []Product {
	return AllProducts
}

func (s *Service) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help")
	s.bot.Send(msg)
}

func (s *Service) List(inputMessage *tgbotapi.Message) {
	products := s.list()
	outputMsg := "List products\n\n"

	for _, product := range products {
		outputMsg += product.Title
		outputMsg += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
	s.bot.Send(msg)
}
func (s *Service) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID
	s.bot.Send(msg)
}

func (s *Service) HandleUpdate(update tgbotapi.Update) {
	var registeredCommands = map[string]func(inputMessage *tgbotapi.Message){
		"help": s.Help,
		"list": s.List,
	}
	if update.Message == nil { // If we got a message
		return
	}
	command, ok := registeredCommands[update.Message.Command()]
	if ok {
		command(update.Message)
	} else {
		s.Default(update.Message)
	}
}

func Help(message *tgbotapi.Message) {
	panic("unimplemented")
}
