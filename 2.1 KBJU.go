package main

import (
	"log"
	"strconv"
	"math"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SVO(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel, chatID int64, pol int, voz int, rost int, ves int, act float64) {
    msg := tgbotapi.NewMessage(chatID, "Для каких целей? ")
    buttons := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Похудение"),
        ),
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Поддержание веса"),
        ),
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Набор массы"),
        ),
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Выйти из опроса"),
        ),
    )
    msg.ReplyMarkup = buttons
    if _, err := bot.Send(msg); err != nil {
        log.Println("Error sending message:", err)
        return
    }

    for {
        update := <-updates
        if update.Message == nil || update.Message.Chat.ID != chatID {
            continue
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        switch update.Message.Text {
        case "Похудение":
            cal := math.Round(0.8*((10*float64(ves) + 6.25*float64(rost) - 5*float64(voz) + float64(pol)) * act))
            bel := math.Round((0.3 * cal) / 4)
            zir := math.Round((0.2 * cal) / 9)
            ugl := math.Round((0.5 * cal) / 4)
            svo := fmt.Sprintf("*Для похудения*\nКоличество КАЛОРИЙ: %.0f\nКоличество БЕЛКА (г): %.0f\nКоличество ЖИРОВ (г): %.0f\nКоличество УГЛЕВОДОВ (г): %.0f", cal, bel, zir, ugl)
            msg := tgbotapi.NewMessage(chatID, svo)
            if _, err := bot.Send(msg); err != nil {
                log.Println("Error sending message:", err)
            }
        case "Поддержание веса":
            cal := math.Round(((10*float64(ves) + 6.25*float64(rost) - 5*float64(voz) + float64(pol)) * act))
            bel := math.Round((0.2 * cal) / 4)
            zir := math.Round((0.3 * cal) / 9)
            ugl := math.Round((0.5 * cal) / 4)
            svo := fmt.Sprintf("*Для поддержания веса*\nКоличество КАЛОРИЙ: %.0f\nКоличество БЕЛКА (г): %.0f\nКоличество ЖИРОВ (г): %.0f\nКоличество УГЛЕВОДОВ (г): %.0f", cal, bel, zir, ugl)
            msg := tgbotapi.NewMessage(chatID, svo)
            if _, err := bot.Send(msg); err != nil {
                log.Println("Error sending message:", err)
            }
        case "Набор массы":
            cal := math.Round(1.15*((10*float64(ves) + 6.25*float64(rost) - 5*float64(voz) + float64(pol)) * act))
            bel := math.Round((0.3 * cal) / 4)
            zir := math.Round((0.2 * cal) / 9)
            ugl := math.Round((0.5 * cal) / 4)
            svo := fmt.Sprintf("*Для набора массы*\nКоличество КАЛОРИЙ: %.0f\nКоличество БЕЛКА (г): %.0f\nКоличество ЖИРОВ (г): %.0f\nКоличество УГЛЕВОДОВ (г): %.0f", cal, bel, zir, ugl)
            msg := tgbotapi.NewMessage(chatID, svo)
            if _, err := bot.Send(msg); err != nil {
                log.Println("Error sending message:", err)
            }
        case "Выйти из опроса":
            Calculators(bot, update, updates)
            return
        default:
            msg := tgbotapi.NewMessage(chatID, "Неправильный ввод. Пожалуйста, выберите одну из целей.")
            if _, err := bot.Send(msg); err != nil {
                log.Println("Error sending message:", err)
                return
            }
            continue
        }
    }
}


func RasKBZU(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) {
	chatID := update.Message.Chat.ID 
	pol := getpol(bot, update, updates) 
	voz := getvoz(bot, update, updates)
	rost := getrost(bot, update, updates)
	ves := getves(bot, update, updates)
	act := getact(bot, update, updates)
    log.Print(pol, voz, rost, ves, act)
	SVO(bot, update, updates, chatID, pol, voz, rost, ves, act)

}

func getpol(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) int {
    chatID := update.Message.Chat.ID
    msg := tgbotapi.NewMessage(chatID, "Введите ваш пол: 👤")
    buttons := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Мужской"),
            tgbotapi.NewKeyboardButton("Женский"),
        ),
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Выйти из опроса"),
        ),
    )
    msg.ReplyMarkup = buttons
    if _, err := bot.Send(msg); err != nil {
        log.Panic(err)
    }

    var pol int // Объявляем переменную пол
    for {
        update := <-updates
        if update.Message == nil {
            continue
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        switch update.Message.Text {
        case "Мужской":
            pol = 5 // Присваиваем значение переменной
        case "Женский":
            pol = -161 // Присваиваем значение переменной
        case "Выйти из опроса":
            Calculators(bot, update, updates)
            return 0
        default:
            msg := tgbotapi.NewMessage(chatID, "Неправильный ввод пола. Пожалуйста, выберите снова")
            if _, err := bot.Send(msg); err != nil {
                log.Panic(err)
            }
            continue
        }
        break
    }
    return pol
}

func getvoz(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) int {
    chatID := update.Message.Chat.ID
    msg := tgbotapi.NewMessage(chatID, "Введите ваш возраст: ")
    buttons := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Выйти из опроса"),
        ),
    )
    msg.ReplyMarkup = buttons
    if _, err := bot.Send(msg); err != nil {
        log.Panic(err)
    }
    var err error
    var voz int
    for {
        update := <-updates
        if update.Message == nil {
            continue
        }

        if update.Message.Text == "Выйти из опроса" {
            Calculators(bot, update, updates)
            return 0
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        voz, err = strconv.Atoi(update.Message.Text)
        if err != nil || voz <= 0 || voz >= 122 {
            msg := tgbotapi.NewMessage(chatID, "Неправильный ввод возраста. Пожалуйста, введите корректный возраст")
            if _, err := bot.Send(msg); err != nil {
                log.Panic(err)
            }
            continue
        }
        break
    }
    return voz
}

func getrost(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) int {
    chatID := update.Message.Chat.ID
    msg := tgbotapi.NewMessage(chatID, "Введите ваш рост (в см): ")
    buttons := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Выйти из опроса"),
        ),
    )
    msg.ReplyMarkup = buttons
    if _, err := bot.Send(msg); err != nil {
        log.Panic(err)
    }
    var err error
    var rost int 
    for {
        update := <-updates
        if update.Message == nil {
            continue
        }

        if update.Message.Text == "Выйти из опроса" {
            Calculators(bot, update, updates)
            return 0
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        rost, err = strconv.Atoi(update.Message.Text)
        if err != nil || rost <= 0 || rost >= 249 {
            msg := tgbotapi.NewMessage(chatID, "Неправильный ввод роста. Пожалуйста, введите корректный рост")
            if _, err := bot.Send(msg); err != nil {
                log.Panic(err)
            }
            continue
        }
        break
    }
    return rost
}

func getves(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) int {
    chatID := update.Message.Chat.ID
    msg := tgbotapi.NewMessage(chatID, "Введите ваш вес (в кг): ")
    buttons := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Выйти из опроса"),
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

        if update.Message.Text == "Выйти из опроса" {
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
    return ves
}

func getact(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) float64 {
    chatID := update.Message.Chat.ID
    msg := tgbotapi.NewMessage(chatID, "Введите ваш уровень физической активности: 🏋️‍♂️")
    buttons := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Малоподвижный образ жизни"),
        ),
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Низкая активность, легкие тренировки 1-3 раза в неделю"),
        ),
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Умеренная активность, тренировки 3-5 дней в неделю"),
        ),
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Высокая активность, интенсивные тренировки 6-7 раз в неделю"),
        ),
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Предельная активность, тяжелая физическая работа и спорт"),
        ),
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Выйти из опроса"),
        ),
    )
    msg.ReplyMarkup = buttons
    if _, err := bot.Send(msg); err != nil {
        log.Panic(err)
    }

    var act float64 // Объявляем переменную act
    for {
        update := <-updates
        if update.Message == nil {
            continue
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        switch update.Message.Text {
        case "Малоподвижный образ жизни":
            act = 1.2
        case "Низкая активность, легкие тренировки 1-3 раза в неделю":
            act = 1.375
        case "Умеренная активность, тренировки 3-5 дней в неделю":
            act = 1.55
        case "Высокая активность, интенсивные тренировки 6-7 раз в неделю":
            act = 1.725
        case "Предельная активность, тяжелая физическая работа и спорт":
            act = 1.9
        case "Выйти из опроса":
            Calculators(bot, update, updates)
            return 0
        default:
            msg := tgbotapi.NewMessage(chatID, "Неправильный ввод активности. Пожалуйста, выберите один из вариантов.")
            if _, err := bot.Send(msg); err != nil {
                log.Panic(err)
            }
            continue
        }
        break
    }
    return act
}
