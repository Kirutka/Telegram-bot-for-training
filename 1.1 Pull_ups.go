package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Pull_ups(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) { // ПОДТЯГИВАНИЯ
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Подтягивания. Введите максимальное количество за раз:")
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
		} else {
			count := Pull_ups_CAL(userCount)
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

func Pull_ups_CAL(userCount int) string {
	razdo4 := fmt.Sprintf("Каждая отдельная цифра это подход. Допустим запись (2 3 4) означает сделать 2 повторения, потом 3, потом 4. Отдыхайте 2-3 минуты между подходами. Программа состоит из опусканий. \nВот ваша программа тренировок: \n" + "\n" + "1 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 2 7 5 5 7;" + "\n" +
		"**Среда**" + "\n" + "- 3 8 6 6 8;" + "\n" +
		"**Пятница**" + "\n" + "- 4 9 6 6 8;" + "\n" + "\n" + "2 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 5 8 7 7 9;" + "\n" +
		"**Среда**" + "\n" + "- 5 9 8 8 9;" + "\n" +
		"**Пятница**" + "\n" + "- 6 10 8 8 10;" + "\n")
	raz4_5 := fmt.Sprintf("Каждая отдельная цифра это подход. Допустим запись (2 3 4) означает сделать 2 повторения, потом 3, потом 4. Отдыхайте 2-3 минуты между подходами. Программа состоит из опусканий. Программа состоит из опусканий. \nВот ваша программа тренировок: \n" + "\n" + "1 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 4 9 6 6 9;" + "\n" +
		"**Среда**" + "\n" + "- 5 9 7 7 9;" + "\n" +
		"**Пятница**" + "\n" + "- 6 10 8 8 10;" + "\n" + "\n" + "2 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 6 11 8 8 11;" + "\n" +
		"**Среда**" + "\n" + "- 7 12 10 10 12;" + "\n" +
		"**Пятница**" + "\n" + "- 8 14 11 11 14;" + "\n")
	raz6_8 := fmt.Sprintf("Каждая отдельная цифра это подход. Допустим запись (2 3 4) означает сделать 2 повторения, потом 3, потом 4. Отдыхайте 2-3 минуты между подходами. Программа состоит из опусканий. \nВот ваша программа тренировок: \n" + "\n" + "1 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 2 3 2 2 3;" + "\n" +
		"**Среда**" + "\n" + "- 2 3 2 2 4;" + "\n" +
		"**Пятница**" + "\n" + "- 3 4 2 2 4;" + "\n" + "\n" + "2 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 3 4 3 3 4;" + "\n" +
		"**Среда**" + "\n" + "- 3 4 3 3 5;" + "\n" +
		"**Пятница**" + "\n" + "- 4 5 4 4 6;" + "\n")
	raz9_11 := fmt.Sprintf("Каждая отдельная цифра это подход. Допустим запись (2 3 4) означает сделать 2 повторения, потом 3, потом 4. Отдыхайте 2-3 минуты между подходами. Программа состоит из опусканий. \nВот ваша программа тренировок: \n" + "\n" + "1 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 3 5 3 3 5;" + "\n" +
		"**Среда**" + "\n" + "- 4 6 4 4 6;" + "\n" +
		"**Пятница**" + "\n" + "- 5 7 5 5 6;" + "\n" + "\n" + "2 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 5 8 5 5 8;" + "\n" +
		"**Среда**" + "\n" + "- 6 9 6 6 8;" + "\n" +
		"**Пятница**" + "\n" + "- 6 9 6 6 10;" + "\n")
	raz12_15 := fmt.Sprintf("Каждая отдельная цифра это подход. Допустим запись (2 3 4) означает сделать 2 повторения, потом 3, потом 4. Отдыхайте 2-3 минуты между подходами. Программа состоит из опусканий. \nВот ваша программа тренировок: \n" + "\n" + "1 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 6 8 6 6 8;" + "\n" +
		"**Среда**" + "\n" + "- 6 9 6 6 9;" + "\n" +
		"**Пятница**" + "\n" + "- 7 10 6 6 9;" + "\n" + "\n" + "2 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 7 10 7 7 10;" + "\n" +
		"**Среда**" + "\n" + "- 8 11 8 8 10;" + "\n" +
		"**Пятница**" + "\n" + "- 9 11 9 9 11;" + "\n")
	raz16_20 := fmt.Sprintf("Каждая отдельная цифра это подход. Допустим запись (2 3 4) означает сделать 2 повторения, потом 3, потом 4. Отдыхайте 2-3 минуты между подходами. Программа состоит из опусканий. \nВот ваша программа тренировок: \n" + "\n" + "1 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 8 11 8 8 10;" + "\n" +
		"**Среда**" + "\n" + "- 9 12 9 9 11;" + "\n" +
		"**Пятница**" + "\n" + "- 9 13 9 9 12;" + "\n" + "\n" + "2 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 10 14 10 10 13;" + "\n" +
		"**Среда**" + "\n" + "- 11 15 10 10 13;" + "\n" +
		"**Пятница**" + "\n" + "- 11 15 11 11 13;" + "\n" + "\n" + "3 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 12 16 11 11 15;" + "\n" +
		"**Среда**" + "\n" + "- 12 16 12 12 16;" + "\n" +
		"**Пятница**" + "\n" + "- 13 17 13 13 16;" + "\n")
	raz21_25 := fmt.Sprintf("Каждая отдельная цифра это подход. Допустим запись (2 3 4) означает сделать 2 повторения, потом 3, потом 4. Отдыхайте 2-3 минуты между подходами. Программа состоит из опусканий. \nВот ваша программа тренировок: \n" + "\n" + "1 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 12 16 12 12 15;" + "\n" +
		"**Среда**" + "\n" + "- 13 16 12 12 16;" + "\n" +
		"**Пятница**" + "\n" + "- 13 17 12 12 16;" + "\n" + "\n" + "2 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 14 19 13 13 18;" + "\n" +
		"**Среда**" + "\n" + "- 14 19 14 14 19;" + "\n" +
		"**Пятница**" + "\n" + "- 15 20 14 14 20;" + "\n" + "\n" + "3 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 16 20 16 16 20;" + "\n" +
		"**Среда**" + "\n" + "- 16 21 16 16 20;" + "\n" +
		"**Пятница**" + "\n" + "- 17 22 16 16 21;" + "\n")
	raz26_30 := fmt.Sprintf("Каждая отдельная цифра это подход. Допустим запись (2 3 4) означает сделать 2 повторения, потом 3, потом 4. Отдыхайте 2-3 минуты между подходами. Программа состоит из опусканий. \nВот ваша программа тренировок: \n" + "\n" + "1 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 16 18 15 15 17;" + "\n" +
		"**Среда**" + "\n" + "- 16 20 16 16 19;" + "\n" +
		"**Пятница**" + "\n" + "- 17 21 16 16 20;" + "\n" + "\n" + "2 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 17 22 17 17 22;" + "\n" +
		"**Среда**" + "\n" + "- 18 23 18 18 22;" + "\n" +
		"**Пятница**" + "\n" + "- 19 25 18 18 24;" + "\n" + "\n" + "3 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 19 26 18 18 25;" + "\n" +
		"**Среда**" + "\n" + "- 19 27 19 19 26;" + "\n" +
		"**Пятница**" + "\n" + "- 20 28 20 20 28;" + "\n")
	raz31_35 := fmt.Sprintf("Каждая отдельная цифра это подход. Допустим запись (2 3 4) означает сделать 2 повторения, потом 3, потом 4. Отдыхайте 2-3 минуты между подходами. Программа состоит из опусканий. \nВот ваша программа тренировок: \n" + "\n" + "1 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 20 25 19 19 23;" + "\n" +
		"**Среда**" + "\n" + "- 22 25 21 21 25;" + "\n" +
		"**Пятница**" + "\n" + "- 23 26 23 23 25;" + "\n" + "\n" + "2 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 24 27 24 24 26;" + "\n" +
		"**Среда**" + "\n" + "- 25 28 24 24 27;" + "\n" +
		"**Пятница**" + "\n" + "- 25 29 25 25 28;" + "\n" + "\n" + "3 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 26 29 25 25 29;" + "\n" +
		"**Среда**" + "\n" + "- 26 30 26 26 30;" + "\n" +
		"**Пятница**" + "\n" + "- 26 32 26 26 32;" + "\n")
	raz36_40 := fmt.Sprintf("Каждая отдельная цифра это подход. Допустим запись (2 3 4) означает сделать 2 повторения, потом 3, потом 4. Отдыхайте 2-3 минуты между подходами. Программа состоит из опусканий. \nВот ваша программа тренировок: \n" + "\n" + "1 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 23 27 22 22 26;" + "\n" +
		"**Среда**" + "\n" + "- 24 28 24 24 28;" + "\n" +
		"**Пятница**" + "\n" + "- 25 29 24 24 29;" + "\n" + "\n" + "2 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 26 30 25 25 30;" + "\n" +
		"**Среда**" + "\n" + "- 26 31 25 25 31;" + "\n" +
		"**Пятница**" + "\n" + "- 26 31 26 26 31;" + "\n" + "\n" + "3 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 27 31 26 26 32;" + "\n" +
		"**Среда**" + "\n" + "- 28 32 26 26 32;" + "\n" +
		"**Пятница**" + "\n" + "- 28 34 27 27 34;" + "\n")
	raz41_50 := fmt.Sprintf("Каждая отдельная цифра это подход. Допустим запись (2 3 4) означает сделать 2 повторения, потом 3, потом 4. Отдыхайте 2-3 минуты между подходами. Программа состоит из опусканий. \nВот ваша программа тренировок: \n" + "\n" + "1 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 25 28 24 24 26;" + "\n" +
		"**Среда**" + "\n" + "- 25 29 25 25 28;" + "\n" +
		"**Пятница**" + "\n" + "- 25 30 25 25 29;" + "\n" + "\n" + "2 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 26 31 25 25 31;" + "\n" +
		"**Среда**" + "\n" + "- 26 32 26 26 32;" + "\n" +
		"**Пятница**" + "\n" + "- 27 32 26 26 32;" + "\n" + "\n" + "3 Неделя" + "\n" +
		"**Понедельник**" + "\n" + "- 27 34 26 26 33;" + "\n" +
		"**Среда**" + "\n" + "- 28 34 26 26 34;" + "\n" +
		"**Пятница**" + "\n" + "- 29 35 27 27 35;" + "\n")

	var itog string
	switch userCount {
	case 0:
		itog = string("Попробуйте висеть на турнике, а затем возвращайтесь)")
	case 1, 2, 3:
		itog = string(razdo4)
	case 4, 5:
		itog = string(raz4_5)
	case 6, 7, 8:
		itog = string(raz6_8)
	case 9, 10, 11:
		itog = string(raz9_11)
	case 12, 13, 14, 15:
		itog = string(raz12_15)
	case 16, 17, 18, 19, 20:
		itog = string(raz16_20)
	case 21, 22, 23, 24, 25:
		itog = string(raz21_25)
	case 26, 27, 28, 29, 30:
		itog = string(raz26_30)
	case 31, 32, 33, 34, 35:
		itog = string(raz31_35)
	case 36, 37, 38, 39, 40:
		itog = string(raz36_40)
	case 41, 42, 43, 44, 45, 46, 47, 48, 49, 50:
		itog = string(raz41_50)
	}
	if userCount >= 51 {
		itog = string("Ого, наша программа не рассчитана на таких силачей как вы)")
	}
	count := itog
	return count
}

func PodDop(userCount int) string {
	count75 := int(math.Round(float64(userCount) * 0.75))
	count80 := int(math.Round(float64(userCount) * 0.80))
	count85 := int(math.Round(float64(userCount) * 0.85))
	count90 := int(math.Round(float64(userCount) * 0.90))
	count75 = normalpod(count75)
	count80 = normalpod(count80)
	count85 = normalpod(count85)
	count90 = normalpod(count90)
	poddop := fmt.Sprintf("Запись (V x P x D) означает: \n(ВЕС х ПОВТОРЕНИЯ х ПОДХОДЫ) \nОтдыхайте 4-5 минут между подходами \nВот ваша программа тренировок: \n"+"\n"+"## 1 Неделя"+"\n"+
		"**Понедельник**"+"\n"+"- %d x 4 x 6;"+"\n"+
		"**Пятница**"+"\n"+"- %d x 4 x 6;"+"\n"+"\n"+"## 2 Неделя"+"\n"+
		"**Понедельник**"+"\n"+"- %d x 4 x 5;"+"\n"+
		"**Пятница**"+"\n"+"- %d x 4 x 5;"+"\n"+"\n"+"## 3 Неделя"+"\n"+
		"**Понедельник**"+"\n"+"- %d x 5 x 4;"+"\n"+
		"**Пятница**"+"\n"+"- %d x 5 x 4;"+"\n"+"\n"+"## 4 Неделя"+"\n"+
		"**Понедельник**"+"\n"+"- %d x 5 x 3;"+"\n"+
		"**Пятница**"+"\n"+"- %d x 5 x 3;"+"\n", count75, count75, count80, count80, count85, count85, count90, count90)
	itog := string(poddop)
	return itog
}

func normalpod(num int) int {
	switch {
	case num%10 <= 2:
		return num - (num % 10)
	case num%10 <= 7:
		return num - (num % 10) + 5
	default:
		return num - (num % 10) + 10
	}
}
