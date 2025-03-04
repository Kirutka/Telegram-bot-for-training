package main

import (
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Progress(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) {
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "Выберите, что вы хотите сделать:")
	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Карта бота"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Связь с поддержкой"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("⬅️ Назад"),
		),
	)
	msg.ReplyMarkup = buttons
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
	for {
		update := <-updates
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Text {
		case "Карта бота":
			kar, _ := ioutil.ReadFile("Map_bot.txt")
			msg := tgbotapi.NewMessage(chatID, string(kar))
			if _, err := bot.Send(msg); err != nil {
				log.Println("Error sending message:", err)
			}
			continue
			// case "Связь с поддержкой":
			// 	Support_communication(bot, update, updates)
		case "Назад":
			Main_menu(bot, update, updates)
		default:
			msg := tgbotapi.NewMessage(chatID, "Неправильный ввод")
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		}
		break
	}
}
