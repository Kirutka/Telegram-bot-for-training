package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Workout(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) {
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "Выберите, что вы хотите сделать:")
	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Подтягивания")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Отжимания на брусьях")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Классический жим лёжа")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Отжимания от пола")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Пресс")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Назад")),
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
		case "Подтягивания":
			Pull_ups(bot, update, updates) // ГОТОВО
		case "Отжимания на брусьях":
			Push_ups_on_the_bars(bot, update, updates) // ГОТОВО
		case "Классический жим лёжа":
			Classic_bench_press(bot, update, updates) // ГОТОВО
		case "Отжимания от пола":
			Push_ups_from_the_floor(bot, update, updates) // ГОТОВО
		case "Пресс":
			Press(bot, update, updates) // ГОТОВО
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
