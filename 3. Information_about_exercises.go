package main

import (
	"io/ioutil"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Information_about_exercises(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) {
	content, _ := ioutil.ReadFile("spisok.txt")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(content))
	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Посмотреть список"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
	msg.ReplyMarkup = buttons
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if update.Message == nil {
			continue
		}
		if number, err := strconv.Atoi(update.Message.Text); err == nil {
			if number >= 1 && number <= 42 {
				count := Info(number)
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, count)
				_, err := bot.Send(msg)
				if err != nil {
					log.Panic(err)
				}
				continue
			}
		}
		user := update.Message.Text
		switch user {
		case "Посмотреть список":
			Information_about_exercises(bot, update, updates) // ГОТОВО
		case "Назад":
			Main_menu(bot, update, updates)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Нет такого варианта ответа 🤔")
			_, err := bot.Send(msg)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}
