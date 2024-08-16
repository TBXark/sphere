package bot

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Button struct {
	Text string
	Type string
	Data any
}

type Message struct {
	Text      string
	Button    [][]Button
	ParseMode models.ParseMode
}

func (m *Message) toInlineKeyboardMarkup() *models.InlineKeyboardMarkup {
	if len(m.Button) == 0 {
		return nil
	}
	keyboard := make([][]models.InlineKeyboardButton, 0, len(m.Button))
	for _, row := range m.Button {
		buttons := make([]models.InlineKeyboardButton, 0, len(row))
		for _, btn := range row {
			buttons = append(buttons, models.InlineKeyboardButton{
				Text:         btn.Text,
				CallbackData: marshalData(btn.Type, btn.Data),
			})
		}
		keyboard = append(keyboard, buttons)
	}
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}

func (m *Message) toSendMessageParams(chatID int64) *bot.SendMessageParams {
	return &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        m.Text,
		ParseMode:   m.ParseMode,
		ReplyMarkup: m.toInlineKeyboardMarkup(),
	}
}

func (m *Message) toEditMessageTextParams(chatID int64, messageID int) *bot.EditMessageTextParams {
	return &bot.EditMessageTextParams{
		ChatID:      chatID,
		MessageID:   messageID,
		Text:        m.Text,
		ParseMode:   m.ParseMode,
		ReplyMarkup: m.toInlineKeyboardMarkup(),
	}
}