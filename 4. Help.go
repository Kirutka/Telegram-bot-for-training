package main

import (
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Help(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) {
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
		case "Карта бота": // ГОТОВО
			kar, _ := ioutil.ReadFile("Map_bot.txt")
			msg := tgbotapi.NewMessage(chatID, string(kar))
			if _, err := bot.Send(msg); err != nil {
				log.Println("Error sending message:", err)
			}
			continue
		case "Связь с поддержкой": // ГОТОВО
			backButton := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Назад"),
				),
			)

			msg := tgbotapi.NewMessage(chatID, "Введите сообщение для поддержки или нажмите 'Назад':")
			msg.ReplyMarkup = backButton
			bot.Send(msg)

			update := <-updates
			if update.Message != nil {
				if update.Message.Text == "Назад" {
					Main_menu(bot, update, updates)
				} else {
					saveSupportMessage(update.Message)
					reply := tgbotapi.NewMessage(chatID, "Сообщение отправлено в поддержку!")
					bot.Send(reply)
					Main_menu(bot, update, updates)
				}
			}
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
