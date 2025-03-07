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
			RasKBZU(bot, update, updates) // ГОТОВО
		case "Индекс массы тела":
			Body_Weight_Index(bot, update, updates) // ГОТОВО
		case "Норма воды":
			The_norm_of_water(bot, update, updates) // ГОТОВО
		case "Калории за пробежку":
			Calories_for_a_run(bot, update, updates) // ГОТОВО
		case "ГТО":
			GTO(bot, update, updates)
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
