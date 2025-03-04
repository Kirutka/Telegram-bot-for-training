package main

import (
	"log"
    "strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Workout(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel){
    chatID := update.Message.Chat.ID
    msg := tgbotapi.NewMessage(chatID, "Выберите, что вы хотите сделать:")
    buttons := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Подтягивания"),),
        tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Отжимания на брусьях"),),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Классический жим лёжа"),),
        tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Отжимания от пола"),),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Пресс"),),
        tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Назад"),),
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
        // case "Подтягивания":
		// 	Pull_ups(bot, update, updates)
        // case "Отжимания на брусьях":
		// 	Push_ups_on_the_bars(bot, update, updates)
        // case "Классический жим лёжа":
		// 	Classic_bench_press(bot, update, updates)
        case "Отжимания от пола": // +++++++++++++++++++++++++++=
			Push_ups_from_the_floor(bot, update, updates)
        // case "Пресс":
		// 	Press(bot, update, updates)
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

func Push_ups_from_the_floor(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) { // отжимания, пресс
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите максимальное количество за раз:")
    buttons := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Назад"),
        ),
    )
    msg.ReplyMarkup = buttons
    _, err := bot.Send(msg)
    if err != nil {
        log.Panic(err)
    }
    for {
        update := <-updates
        if update.Message == nil {
            break
        }
        if update.Message.Text == "Назад"{
            Workout(bot, update, updates)
        }
        userCount, err := strconv.Atoi(update.Message.Text)
        if userCount < 0 || err != nil{
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите положительное, целое число")
            _, err = bot.Send(msg)
            if err != nil {
            log.Panic(err)
            } 
        } else {
    count := Push_ups_from_the_floor_CAL(userCount)
    msg = tgbotapi.NewMessage(update.Message.Chat.ID, count)
    buttons := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Назад"),
        ),
    )
    msg.ReplyMarkup = buttons
    _, err := bot.Send(msg)
    if err != nil {
        log.Panic(err)
    }
        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
        for update := range updates {
            if update.Message == nil {
                continue
            }
            user := update.Message.Text
            switch user {
            case "Назад":
                Workout(bot, update, updates)
            default:
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Нет такого варианта ответа")
                _, err := bot.Send(msg)
                if err != nil {
                log.Panic(err)
                }
                }
            }
        }
    }
}