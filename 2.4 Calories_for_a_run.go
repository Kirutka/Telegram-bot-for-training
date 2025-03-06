package main 

import (
	"fmt"
	"log"
	"strconv"

   tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Ves(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) float64 {
	chatID := update.Message.Chat.ID
    msg := tgbotapi.NewMessage(chatID, "Введите ваш вес (в кг): ")
    buttons := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Назад"),
        ),
    )
    msg.ReplyMarkup = buttons
    if _, err := bot.Send(msg); err != nil {
        log.Panic(err)
    }
    var err error
    var ves int 
    for {
        update := <-updates
        if update.Message == nil {
            continue
        }

        if update.Message.Text == "Назад" {
            Calculators(bot, update, updates)
            return 0
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        ves, err = strconv.Atoi(update.Message.Text)
        if err != nil || ves <= 0 || ves >= 645 {
            msg := tgbotapi.NewMessage(chatID, "Неправильный ввод веса. Пожалуйста, введите корректный вес.")
            if _, err := bot.Send(msg); err != nil {
                log.Panic(err)
            }
            continue
        }
        break
    }
    return float64(ves)

}

func Get_speed(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) float64 {
	chatID := update.Message.Chat.ID
    msg := tgbotapi.NewMessage(chatID, "Введите вашу скорость (в км/ч): ")
    buttons := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Назад"),
        ),
    )
    msg.ReplyMarkup = buttons
    if _, err := bot.Send(msg); err != nil {
        log.Panic(err)
    }
    var err error
    var speed int 
    for {
        update := <-updates
        if update.Message == nil {
            continue
        }

        if update.Message.Text == "Назад" {
            Calculators(bot, update, updates)
            return 0
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        speed, err = strconv.Atoi(update.Message.Text)
        if err != nil || speed <= 0 || speed >= 645 {
            msg := tgbotapi.NewMessage(chatID, "Неправильный ввод скорости. Пожалуйста, введите корректную скорость.")
            if _, err := bot.Send(msg); err != nil {
                log.Panic(err)
            }
            continue
        }
        break
    }
    return float64(speed)

}

func Get_time(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) float64 {
	chatID := update.Message.Chat.ID
    msg := tgbotapi.NewMessage(chatID, "Сколько времени вы пробежали (в минутах): ")
    buttons := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Назад"),
        ),
    )
    msg.ReplyMarkup = buttons
    if _, err := bot.Send(msg); err != nil {
        log.Panic(err)
    }
    var err error
    var time int 
    for {
        update := <-updates
        if update.Message == nil {
            continue
        }

        if update.Message.Text == "Назад" {
            Calculators(bot, update, updates)
            return 0
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        time, err = strconv.Atoi(update.Message.Text)
        if err != nil || time <= 0 || time >= 645 {
            msg := tgbotapi.NewMessage(chatID, "Неправильный ввод времени. Пожалуйста, введите корректное время")
            if _, err := bot.Send(msg); err != nil {
                log.Panic(err)
            }
            continue
        }
        break
    }
    return float64(time)

}

func Calories_for_a_run(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) {
	chatID := update.Message.Chat.ID 
	speed := Get_speed(bot, update, updates) 
	ves := Ves(bot, update, updates)
	time := Get_time(bot, update, updates)
	
	call := (0.1 * speed + 3.5) * ves * time / 20
	svo := fmt.Sprintf("Вы потратили калорий: %0.f ", call)
            msg := tgbotapi.NewMessage(chatID, svo)
            if _, err := bot.Send(msg); err != nil {
                log.Println("Error sending message:", err)
            }
			Calculators(bot, update, updates)

}