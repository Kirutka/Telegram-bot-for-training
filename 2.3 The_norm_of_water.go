package main 

import (
     "fmt"
     "log"
     "strconv"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Get_ves(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) float64 {
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

func The_norm_of_water(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel){
	chatID := update.Message.Chat.ID 
	ves := Get_ves(bot, update, updates)
	norma := (ves*33)/1000
	svo := fmt.Sprintf("Ваша норма воды %.2f л", norma)
            msg := tgbotapi.NewMessage(chatID, svo)
            if _, err := bot.Send(msg); err != nil {
                log.Println("Error sending message:", err)
            }
			Calculators(bot, update, updates)

}
