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
            bot.Send(msg)

            for update := range updates {
                if update.Message == nil {
                    continue
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
    case gender == "М" && age >= 6 && age <= 18:
        filePath = "files/male_young.pdf"
    case gender == "М" && age > 18 && age <= 60:
        filePath = "files/male_adult.pdf"
    case gender == "М" && age > 60:
        filePath = "files/male_senior.pdf"
    case gender == "Ж" && age >= 6 && age <= 18:
        filePath = "files/female_young.pdf"
    case gender == "Ж" && age > 18 && age <= 60:
        filePath = "files/female_adult.pdf"
    case gender == "Ж" && age > 60:
        filePath = "files/female_senior.pdf"
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