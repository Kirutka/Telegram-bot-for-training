package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Calculators(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) {
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "Выберите, что вы хотите сделать:")
	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("КБЖУ"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Индекс массы тела"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Норма воды"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Калории за пробежку"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Процент жира"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Спортивные разряды"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("ГТО"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
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
		case "КБЖУ":
			Main_menu(bot, update, updates)
		case "Индекс массы тела":
			Main_menu(bot, update, updates)
		case "Норма воды":
			Main_menu(bot, update, updates)
		case "Калории за пробежку":
			Main_menu(bot, update, updates)
		case "Процент жира":
			Main_menu(bot, update, updates)
		case "Спортивные разряды":
			Main_menu(bot, update, updates)
		case "ГТО":
			Main_menu(bot, update, updates)
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
