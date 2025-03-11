package main

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GTO(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) {
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "Введите ваш пол (М или Ж):")
	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("М"),
			tgbotapi.NewKeyboardButton("Ж"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
	msg.ReplyMarkup = buttons
	bot.Send(msg)

	var gender string
	var age int

	for update := range updates {
		if update.Message == nil {
			continue
		}

		userInput := update.Message.Text

		if userInput == "Назад" {
			Main_menu(bot, update, updates)
			return
		}

		if userInput == "М" || userInput == "Ж" {
			gender = userInput

			msg := tgbotapi.NewMessage(chatID, "Введите ваш возраст (от 6 до 100 лет):")
			chatID := update.Message.Chat.ID
			buttons := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Назад"),
				),
			)
			msg.ReplyMarkup = buttons
			bot.Send(msg)

			for update := range updates {
				if update.Message == nil {
					continue
				}
				userInput := update.Message.Text

				if userInput == "Назад" {
					Calculators(bot, update, updates)
					return
				}

				ageInput := update.Message.Text

				parsedAge, err := strconv.Atoi(ageInput)
				if err != nil || parsedAge < 6 || parsedAge > 100 {
					msg := tgbotapi.NewMessage(chatID, "Некорректный возраст. Введите число от 6 до 100:")
					bot.Send(msg)
					continue
				}

				age = parsedAge
				break
			}

			sendPDFFiles(bot, chatID, gender, age)
			return
		} else {
			msg := tgbotapi.NewMessage(chatID, "Некорректный пол. Введите 'М' или 'Ж':")
			bot.Send(msg)
		}
	}
}

func sendPDFFiles(bot *tgbotapi.BotAPI, chatID int64, gender string, age int) {
	var filePath string

	switch {
	case age >= 6 && age <= 7:
		filePath = "ГТО/Школьники (6-7).pdf"
	case age > 8 && age <= 9:
		filePath = "ГТО/Школьники (8-9).pdf"
	case age >= 10 && age <= 11:
		filePath = "ГТО/Школьники (10-11).pdf"
	case age > 12 && age <= 13:
		filePath = "ГТО/Школьники (12-13).pdf"
	case age >= 14 && age <= 15:
		filePath = "ГТО/Школьники (14-15).pdf"
	case age > 16 && age <= 17:
		filePath = "ГТО/Школьники (16-17).pdf"

	case gender == "М" && age >= 18 && age <= 19:
		filePath = "ГТО/Мужчины (18-19).pdf"
	case gender == "М" && age >= 20 && age <= 24:
		filePath = "ГТО/Мужчины (20-24).pdf"
	case gender == "М" && age >= 25 && age <= 29:
		filePath = "ГТО/Мужчины (25-29).pdf"
	case gender == "М" && age >= 30 && age <= 34:
		filePath = "ГТО/Мужчины (30-34).pdf"
	case gender == "М" && age >= 35 && age <= 39:
		filePath = "ГТО/Мужчины (35-39).pdf"
	case gender == "М" && age >= 40 && age <= 44:
		filePath = "ГТО/Мужчины (40-44).pdf"
	case gender == "М" && age >= 45 && age <= 49:
		filePath = "ГТО/Мужчины (45-49).pdf"
	case gender == "М" && age >= 50 && age <= 54:
		filePath = "ГТО/Мужчины (50-54).pdf"
	case gender == "М" && age >= 55 && age <= 59:
		filePath = "ГТО/Мужчины (55-59).pdf"
	case gender == "М" && age >= 60 && age <= 64:
		filePath = "ГТО/Мужчины (60-64).pdf"
	case gender == "М" && age >= 65 && age <= 69:
		filePath = "ГТО/Мужчины (65-69).pdf"
	case gender == "М" && age >= 70:
		filePath = "ГТО/Мужчины (70-70+).pdf"

	case gender == "Ж" && age >= 18 && age <= 19:
		filePath = "ГТО/Женщины (18-19).pdf"
	case gender == "Ж" && age >= 20 && age <= 24:
		filePath = "ГТО/Женщины (20-24).pdf"
	case gender == "Ж" && age >= 25 && age <= 29:
		filePath = "ГТО/Женщины (25-29).pdf"
	case gender == "Ж" && age >= 30 && age <= 34:
		filePath = "ГТО/Женщины (30-34).pdf"
	case gender == "Ж" && age >= 35 && age <= 39:
		filePath = "ГТО/Женщины (35-39).pdf"
	case gender == "Ж" && age >= 40 && age <= 44:
		filePath = "ГТО/Женщины (40-44).pdf"
	case gender == "Ж" && age >= 45 && age <= 49:
		filePath = "ГТО/Женщины (45-49).pdf"
	case gender == "Ж" && age >= 50 && age <= 54:
		filePath = "ГТО/Женщины (50-54).pdf"
	case gender == "Ж" && age >= 55 && age <= 59:
		filePath = "ГТО/Женщины (55-59).pdf"
	case gender == "Ж" && age >= 60 && age <= 64:
		filePath = "ГТО/Женщины (60-64).pdf"
	case gender == "Ж" && age >= 65 && age <= 69:
		filePath = "ГТО/Женщины (65-69).pdf"
	case gender == "Ж" && age >= 70:
		filePath = "ГТО/Женщины (70-70+).pdf"
	default:
		msg := tgbotapi.NewMessage(chatID, "Не удалось определить подходящий файл.")
		bot.Send(msg)
		return
	}

	document := tgbotapi.NewDocument(chatID, tgbotapi.FilePath(filePath))
	_, err := bot.Send(document)
	if err != nil {
		log.Printf("Ошибка при отправке файла: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Произошла ошибка при отправке файла.")
		bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(chatID, "Файл успешно отправлен!")
		bot.Send(msg)
	}
}
