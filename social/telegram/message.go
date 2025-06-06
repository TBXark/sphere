package telegram

import (
	"bytes"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Button = models.InlineKeyboardButton

func NewButton[T any](text, route string, data T) Button {
	return Button{
		Text:         text,
		CallbackData: MarshalData(route, data),
	}
}

func NewURLButton(text, url string) Button {
	return Button{
		Text: text,
		URL:  url,
	}
}

func NewBytesInputFile(name string, data []byte) models.InputFile {
	return &models.InputFileUpload{
		Filename: name,
		Data:     bytes.NewReader(data),
	}
}

func NewStringInputFile(url string) models.InputFile {
	return &models.InputFileString{
		Data: url,
	}
}

type Message struct {
	Text      string
	Media     models.InputFile
	ParseMode models.ParseMode
	Button    [][]models.InlineKeyboardButton
}

func (m *Message) toSendMessageParams(chatID int64) *bot.SendMessageParams {
	params := &bot.SendMessageParams{
		ChatID:    chatID,
		Text:      m.Text,
		ParseMode: m.ParseMode,
	}
	if len(m.Button) > 0 {
		params.ReplyMarkup = &models.InlineKeyboardMarkup{
			InlineKeyboard: m.Button,
		}
	}
	return params
}

func (m *Message) toSendPhotoParams(chatID int64) *bot.SendPhotoParams {
	params := &bot.SendPhotoParams{
		ChatID:    chatID,
		Photo:     m.Media,
		Caption:   m.Text,
		ParseMode: m.ParseMode,
	}
	if len(m.Button) > 0 {
		params.ReplyMarkup = &models.InlineKeyboardMarkup{
			InlineKeyboard: m.Button,
		}
	}
	return params
}

func (m *Message) toEditMessageTextParams(chatID int64, messageID int) *bot.EditMessageTextParams {
	params := &bot.EditMessageTextParams{
		ChatID:    chatID,
		MessageID: messageID,
		Text:      m.Text,
		ParseMode: m.ParseMode,
	}
	if len(m.Button) > 0 {
		params.ReplyMarkup = &models.InlineKeyboardMarkup{
			InlineKeyboard: m.Button,
		}
	}
	return params
}

func (m *Message) toEditMessageCaptionParams(chatID int64, messageID int) *bot.EditMessageCaptionParams {
	params := &bot.EditMessageCaptionParams{
		ChatID:    chatID,
		MessageID: messageID,
		Caption:   m.Text,
		ParseMode: m.ParseMode,
	}
	if len(m.Button) > 0 {
		params.ReplyMarkup = &models.InlineKeyboardMarkup{
			InlineKeyboard: m.Button,
		}
	}
	return params
}

func (m *Message) toEditMessageMediaParams(chatID int64, messageID int) *bot.EditMessageMediaParams {
	params := &bot.EditMessageMediaParams{
		ChatID:    chatID,
		MessageID: messageID,
		Media:     nil,
	}
	photo := &models.InputMediaPhoto{
		Caption:   m.Text,
		ParseMode: m.ParseMode,
	}
	if upload, ok := m.Media.(*models.InputFileUpload); ok {
		photo.Media = "attach://" + upload.Filename
		photo.MediaAttachment = upload.Data
	}
	if url, ok := m.Media.(*models.InputFileString); ok {
		photo.Media = url.Data
	}
	params.Media = photo
	if len(m.Button) > 0 {
		params.ReplyMarkup = &models.InlineKeyboardMarkup{
			InlineKeyboard: m.Button,
		}
	}
	return params
}
