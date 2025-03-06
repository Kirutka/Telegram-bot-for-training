package main

import (
	"io/ioutil"
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6611366577:AAGJBRUV0MmjkWfSWGPKhfchGw81uOVtdUk")
	if err != nil {
		log.Panic(err)
	}

	updates := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))

	for update := range updates {
		if update.Message == nil {
			continue
		}

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			var msg tgbotapi.MessageConfig
			switch update.Message.Text {
			case "/start":
				priv, err := ioutil.ReadFile("Greetings.txt")
				if err != nil {
					log.Println("Ошибка чтения файла:", err)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Не удалось прочитать файл.")
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, string(priv))
					Main_menu(bot, update, updates)
				}
				
			default:
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Для начала работы нажмите - /start")
			}

				if _, err := bot.Send(msg); err != nil {
					log.Print(err)
				}
		
	}
}

func Main_menu(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Это главное меню")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Тренировки"),
			tgbotapi.NewKeyboardButton("Калькуляторы"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Информация об упражнениях"),
			tgbotapi.NewKeyboardButton("Помощь"),
		),
	)

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
		for update := range updates {
			if update.Message == nil {
				continue
			}
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			var mess tgbotapi.MessageConfig
			switch update.Message.Text {
			case "Тренировки":
				Workout(bot, update, updates)
			case "Калькуляторы":
				Calculators(bot, update, updates)
			case "Информация об упражнениях":
				Information_about_exercises(bot, update, updates)
			case "Помощь":
				Help(bot, update, updates)
			default:
				mess = tgbotapi.NewMessage(update.Message.Chat.ID, "Нет такого варианта ответа")
			}

			if _, err := bot.Send(mess); err != nil {
				log.Print(err)
			}
		}
}

