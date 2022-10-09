package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type TelegramMessage struct {
	BotToken string
	ChatIDs  []string
	Message  string
	Mode     string
}

func NewTelegramMessage(recipient, message, mode string) *TelegramMessage {
	botToken := os.Getenv("BOT_TOKEN")
	tm := TelegramMessage{
		BotToken: botToken,
		Message:  message,
		Mode:     mode,
	}
	tm.GetChatIds(recipient)
	return &tm
}

func (tm *TelegramMessage) Send() {
	for _, chatId := range tm.ChatIDs {
		if err := tm.SendMessage(chatId); err != nil {
			log.Println("send message error:", err)
		}
	}
}

func (tm *TelegramMessage) SendMessage(ChatId string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tm.BotToken)
	body, _ := json.Marshal(map[string]string{
		"chat_id":    ChatId,
		"text":       tm.Message,
		"parce_mode": tm.Mode,
	})
	_, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	log.Println("send message to:", ChatId)
	return nil
}

func (tm *TelegramMessage) GetChatIds(recipient string) {
	// TODO: Get chat_ids from database

	// Test recipients
	recipients := map[string][]string{"naliway": {"345182391"}}
	chatIds := recipients[recipient]

	tm.ChatIDs = chatIds
	log.Println("chat ids:", tm.ChatIDs)
}
