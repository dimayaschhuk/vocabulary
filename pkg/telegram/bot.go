package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run() {
	bot, err := tgbotapi.NewBotAPI("5673296847:AAE2Z2Bz1uFt1SVL6O3r3khshGBbAnbbHeg")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

func Send() {
	bot, err := tgbotapi.NewBotAPI("5673296847:AAE2Z2Bz1uFt1SVL6O3r3khshGBbAnbbHeg")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	msg := tgbotapi.NewMessage(563738410, "test")
	bot.Send(msg)

	article := tgbotapi.NewInlineQueryResultArticle(563738410, "Echo", update.InlineQuery.Query)
	article.Description = update.InlineQuery.Query

	inlineConf := InlineConfig{
		InlineQueryID: update.InlineQuery.ID,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       []interface{}{article},
	}

	if _, err := bot.Request(inlineConf); err != nil {
		log.Println(err)
	}
}
