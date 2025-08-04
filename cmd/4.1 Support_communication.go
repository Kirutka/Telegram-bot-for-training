package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func saveSupportMessage(msg *tgbotapi.Message) {
	filename := "text/support_messages.txt"
	user := msg.From.UserName
	if user == "" {
		user = "NoUsername"
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	text := msg.Text

	entry := fmt.Sprintf("User: %s\nTime: %s\nMessage: %s\n\n", user, timestamp, text)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	if _, err := f.WriteString(entry); err != nil {
		log.Println("Error writing to file:", err)
	}

}
