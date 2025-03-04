package main 

import (
	"math"
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ClassGym(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) { // классический жим лёжа
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите свой максимальный жим на раз: ")
    buttons := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("⬅️ Назад"),
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
        if update.Message.Text == "Назад" {
            Workout(bot, update, updates)
        }

        userCount, err := strconv.Atoi(update.Message.Text)
        if userCount < 0 || err != nil {
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите положительное, целое число")
            _, err = bot.Send(msg)
            if err != nil {
                log.Panic(err)
            }
        } else if userCount == 0 {
            count := "Пробуйте держать на весу пустой гриф"
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, count)
            _, err = bot.Send(msg)
            if err != nil {
                log.Panic(err)
            }
            break
        } else {
            count := ClassicGym(userCount)
            msg = tgbotapi.NewMessage(update.Message.Chat.ID, count)
            buttons := tgbotapi.NewReplyKeyboard(
                tgbotapi.NewKeyboardButtonRow(
                    tgbotapi.NewKeyboardButton("Назад"),
                ),
                tgbotapi.NewKeyboardButtonRow(
                    tgbotapi.NewKeyboardButton("Главное меню"),
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
                case "Главное меню":
                    Main_menu(bot, update, updates)
                default:
                    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Нет такого варианта ответа 🤔")
                    _, err := bot.Send(msg)
                    if err != nil {
                        log.Panic(err)
                    }
                }
            }
        }
    }
}

func normal(num int) int {
    switch {
    case num%10 <= 2:
        return num - (num % 10)
    case num%10 <= 7:
        return num - (num % 10) + 5
    default:
        return num - (num % 10) + 10
    }
}

func ClassicGym(userCount int) string {
	count50 := int(math.Round(float64(userCount) * 0.50))
                count55 := int(math.Round(float64(userCount) * 0.55))
                count60 := int(math.Round(float64(userCount) * 0.60))
                count65 := int(math.Round(float64(userCount) * 0.65))
                count70 := int(math.Round(float64(userCount) * 0.70))
                count75 := int(math.Round(float64(userCount) * 0.75))
                count77 := int(math.Round(float64(userCount) * 0.77))
                count80 := int(math.Round(float64(userCount) * 0.80))
                count82 := int(math.Round(float64(userCount) * 0.82))
                count825 := int(math.Round(float64(userCount) * 0.825))
                count85 := int(math.Round(float64(userCount) * 0.85))
                count87 := int(math.Round(float64(userCount) * 0.87))
                count90 := int(math.Round(float64(userCount) * 0.90))
                count925 := int(math.Round(float64(userCount) * 0.925))
                count95 := int(math.Round(float64(userCount) * 0.95))
                count102 := int(math.Round(float64(userCount) * 1.02))
                
                count50 = normal(count50) 
                count55 = normal(count55) 
                count60 = normal(count60)
                count65 = normal(count65) 
                count70 = normal(count70) 
                count75 = normal(count75) 
                count77 = normal(count77) 
                count80 = normal(count80) 
                count82 = normal(count82) 
                count825 = normal(count825) 
                count85 = normal(count85) 
                count87 = normal(count87) 
                count90 = normal(count90) 
                count925 = normal(count925) 
                count95 = normal(count95) 
                count102 = normal(count102) 

                tren1 := fmt.Sprintf("\n" + "## 1 Неделя" + "\n" + 
                "**Понедельник**" + "\n" + "1) %d кг x 3 x 5;" + "\n" + 
                "**Пятница**" + "\n" + "1) %d кг x 5 x 2;" + "\n", int(count75), int(count60))
                tren2 := fmt.Sprintf("\n" + "## 2 Неделя" + "\n" + 
                "**Понедельник**" + "\n" + "1) %d кг x 4 x 2;" + "\n" + "2) %d кг x 3 x 5;" + "\n" + 
                "**Пятница**" + "\n" + "1) %d кг x 3 x 2;" + "\n" + "2) %d кг x 3 x 5;" + "\n", int(count70), int(count77), int(count70), int(count80))
                tren3 := fmt.Sprintf("\n" + "## 3 Неделя" + "\n" + 
                "**Понедельник**" + "\n" + "1) %d кг x 3 x 6; "+ "\n" + 
                "**Пятница**" + "\n" + "1) %d кг x 5 x 5;" + "\n", int(count80), int(count70))
                tren4 := fmt.Sprintf("\n" + "## 4 Неделя" + "\n" + 
                "**Понедельник**" + "\n" + "1) %d кг x 4 x 2; "+ "\n" +"2) %d кг x 4 x 2;" + "\n" +"3) %d кг x 3 x 2;" + "\n" + "4) %d кг x 2 x 2;" + "\n" + "5) %d кг x 2 x 3;" + "\n" + "**Пятница**" + "\n" + "1) %d кг x 3 x 2; "+ "\n" +"2) %d кг x 3 x 5;" + "\n", int(count70), int(count75), int(count77), int(count85), int(count82), int(count70), int(count75))
                tren5 := fmt.Sprintf("\n" + "## 5 Неделя" + "\n" + 
                "**Понедельник**" + "\n" + "1) %d кг x 3 x 2; "+ "\n" +"2) %d кг x 3 x 6;" + "\n" + 
                "**Пятница**" + "\n" + "1) %d кг x 6 x 1; "+ "\n" + "2) %d кг x 5 x 1;" + "\n" + "3) %d кг x 4 x 1;" + "\n" + "4) %d кг x 3 x 1;" + "\n" + "5) %d кг x 2 x 1;" + "\n" + "6) %d кг x 1 x 1;" + "\n" + "7) %d кг x 3 x 1;" + "\n" + "8) %d кг x 4 x 1;" + "\n" + "9) %d кг x 6 x 1;" + "\n" + "10) %d кг x 8 x 1;"+ "\n", int(count70), int(count80), int(count60), int(count65), int(count70), int(count80), int(count85), int(count87), int(count825), int(count75), int(count60), int(count55))
                tren6 := fmt.Sprintf("\n" + "## 6 Неделя" + "\n" + 
                "**Понедельник**" + "\n" + "1) %d кг x 4 x 2; "+ "\n" +"2) %d кг x 3 x 2;" + "\n" + "3) %d кг x 2 x 3;" + "\n" + 
                "**Пятница**" + "\n" + "1) %d x 4 x 4; " + "\n", int(count70), int(count80), int(count85), int(count80))
                tren7 := fmt.Sprintf("\n" + "## 7 Неделя" + "\n" + 
                "**Понедельник**" + "\n" + "1) %d кг x 3 x 2; "+ "\n" +"2) %d кг x 3 x 2;" + "\n" +"3) %d кг x 2 x 2; "+ "\n" +"4) %d кг x 3 x 2;" + "\n" + 
                "**Пятница**" + "\n" + "1) %d кг x 3 x 2; "+ "\n" +"2) %d кг x 3 x 1;" + "\n" + "3) %d кг x 2 x 3; "+ "\n" +"4) %d кг x 1 x 2;"+ "\n", int(count70), int(count80), int(count85), int(count80), int(count70), int(count80), int(count85), int(count87))
                tren8 := fmt.Sprintf("\n" + "## 8 Неделя" + "\n" + 
                "**Понедельник**" + "\n" + "1) %d кг x 3 x 2;" + "\n" + "2) %d кг x 3 x 2;" + "\n"+ "3) %d кг x 2 x 5;" + "\n" + 
                "**Пятница**" + "\n" + "1) %d кг x 3 x 2;" + "\n" + "2) %d кг x 4 x 4;"+ "\n", int(count70), int(count80), int(count85), int(count70), int(count80))
                tren9 := fmt.Sprintf("\n" + "## 9 Неделя" + "\n" + 
                "**Понедельник**" + "\n" + "1) %d кг x 8 x 1; "+ "\n" +"2) %d кг x 6 x 1;" + "\n" +"3) %d кг x 5 x 2; "+ "\n" +"4) %d кг x 4 x 2; "+ "\n" +"5) %d кг x 3 x 2;" + "\n" + "6) %d кг x 2 x 2;" + "\n" + 
                "**Пятница**" + "\n" + "1) %d кг x 3 x 2; "+ "\n" +"2) %d кг x 3 x 2;" + "\n" + "3) %d кг x 2 x 3; "+ "\n", int(count50), int(count60), int(count70), int(count80), int(count85), int(count87), int(count70), int(count80), int(count85))
                tren10 := fmt.Sprintf("\n" + "## 10 Неделя" + "\n" + 
                "**Понедельник**" + "\n" + "1) %d кг x 6 x 1;" + "\n" + "2) %d кг x 5 x 1;" + "\n" +"3) %d кг x 4 x 4;" + "\n" + 
                "**Пятница**" + "\n" + "1) %d кг x 2 x 4;" + "\n", int(count60), int(count70), int(count80), int(count80))
                tren11 := fmt.Sprintf("\n" + "## 11 Неделя" + "\n" + 
                "**Понедельник**" + "\n" + "1) %d кг x 6 x 1;" + "\n" + "2) %d кг x 5 x 1;" + "\n" + "3) %d кг x 4 x 1;" + "\n" + "4) %d кг x 3 x 1;" + "\n" + "5) %d кг x 3 x 1;" + "\n" + "6) %d кг x 3 x 1;" + "\n" + "7) %d кг x 3 x 1;" + "\n" + 
                "**Пятница**" + "\n" + "1) %d кг x 5 x 5;" + "\n" + "2) %d кг x 2 x 4;" + "\n" , int(count60), int(count70), int(count80), int(count85), int(count90), int(count925), int(count95), int(count75), int(count80))
                tren12 := fmt.Sprintf("\n" + "## 12 Неделя" + "\n" + 
                "**Понедельник**" + "\n" + "1) %d кг x 4 x 2;" + "\n" + "2) %d кг x 3 x 1;" + "\n" + "3) %d кг x 2 x 3;" + "\n" + 
                "**Пятница**" + "\n" + "1) %d кг x 4 x 4;" + "\n", int(count70), int(count80), int(count825), int(count70))
                tren13 := fmt.Sprintf("\n" + "## 13 Неделя" + "\n" + 
                "**Понедельник**" + "\n" + "1) %d кг x 6 x 1;" + "\n" + "2) %d кг x 5 x 1;" + "\n" + "3) %d кг x 3 x 1;" + "\n" + "4) %d кг x 2 x 1;" + "\n" + "5) %d кг x 1 x 1;" + "\n" + "6) %d кг x 1 x 1;" + "\n" + "7) %d кг x 1 x 1;" + "\n" + "И так до потолка, жать на 1 раз увеличивая вес", int(count50), int(count60), int(count70), int(count80), int(count90), userCount, int(count102))
				count := fmt.Sprintf("Запись (V x P x D) означает: \n(ВЕС х ПОВТОРЕНИЯ х ПОДХОДЫ). Отдыхайте 4-5 минут между подходами. \nВот ваша программа тренировок на 13 недель: " + "\n" + "%s %s %s %s %s %s %s %s %s %s %s %s %s" , tren1, tren2, tren3, tren4, tren5, tren6, tren7, tren8, tren9, tren10, tren11, tren12, tren13)
				return count
}
