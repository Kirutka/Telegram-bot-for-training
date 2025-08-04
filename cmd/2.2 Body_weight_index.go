package main

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// BodyWeightIndex - Главная функция для расчета ИМТ
func Body_Weight_Index(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) {
	chatID := update.Message.Chat.ID

	// Запрашиваем вес
	weight := requestWeight(bot, update, updates)
	if weight == 0 {
		return // Если вес не был введен корректно, завершаем выполнение
	}

	// Запрашиваем рост
	height := requestHeight(bot, update, updates)
	if height == 0 {
		return // Если рост не был введен корректно, завершаем выполнение
	}

	// Рассчитываем ИМТ
	bmi := calculateBMI(weight, height)
	diagnosis := interpretBMI(bmi)

	// Формируем ответ
	response := fmt.Sprintf("Индекс массы тела: %.1f\nПоказатель: %s", bmi, diagnosis)
	msg := tgbotapi.NewMessage(chatID, response)
	if _, err := bot.Send(msg); err != nil {
		log.Println("Ошибка отправки сообщения:", err)
	}
	Calculators(bot, update, updates)
}

// requestWeight - Запрашивает вес пользователя
func requestWeight(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) int {
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "Введите ваш вес в килограммах (например, 70):")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
	bot.Send(msg)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Обработка кнопки "Назад"
		if update.Message.Text == "Назад" {
			Calculators(bot, update, updates)
			return 0
		}

		weight, err := strconv.Atoi(update.Message.Text)
		if err != nil || weight <= 0 || weight >= 645 {
			msg := tgbotapi.NewMessage(chatID, "Неправильный ввод веса. Пожалуйста, введите корректное значение (кг):")
			bot.Send(msg)
			continue
		}

		// Завершаем обработку после получения корректного значения
		return weight
	}
	return 0
}

// requestHeight - Запрашивает рост пользователя
func requestHeight(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) float64 {
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "Введите ваш рост в сантиметрах (например, 175):")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
	bot.Send(msg)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Обработка кнопки "Назад"
		if update.Message.Text == "Назад" {
			Calculators(bot, update, updates)
			return 0
		}

		height, err := strconv.ParseFloat(update.Message.Text, 64)
		if err != nil || height <= 0 || height >= 249 {
			msg := tgbotapi.NewMessage(chatID, "Неправильный ввод роста. Пожалуйста, введите корректное значение (см):")
			bot.Send(msg)
			continue
		}

		// Завершаем обработку после получения корректного значения
		return height
	}
	return 0
}

// calculateBMI - Вычисляет индекс массы тела
func calculateBMI(weight int, height float64) float64 {
	heightInMeters := height / 100
	return float64(weight) / (heightInMeters * heightInMeters)
}

// interpretBMI - Интерпретирует значение ИМТ
func interpretBMI(bmi float64) string {
	switch {
	case bmi <= 16:
		return "Выраженный дефицит массы тела"
	case bmi > 16 && bmi <= 18.5:
		return "Недостаточная (дефицит) масса тела"
	case bmi > 18.5 && bmi <= 25:
		return "Норма"
	case bmi > 25 && bmi <= 30:
		return "Избыточная масса тела (предожирение)"
	case bmi > 30 && bmi <= 35:
		return "Ожирение первой степени"
	case bmi > 35 && bmi <= 40:
		return "Ожирение второй степени"
	default:
		return "Ожирение третьей степени"
	}
}
