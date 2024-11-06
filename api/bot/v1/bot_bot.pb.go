// Code generated by protoc-gen-sphere. DO NOT EDIT.
// versions:
// source: bot/v1/bot.proto

package botv1

import (
	context "context"
	bot "github.com/go-telegram/bot"
	models "github.com/go-telegram/bot/models"
	telegram "github.com/tbxark/sphere/pkg/telegram"
)

var _ = new(context.Context)
var _ = new(telegram.Message)
var _ = new(bot.Bot)
var _ = new(models.Update)

const BotHandlerBotServiceCounter = "/bot.v1.BotService/Counter"
const BotHandlerBotServiceStart = "/bot.v1.BotService/Start"

type BotServiceServer interface {
	Counter(context.Context, *CounterRequest) (*CounterResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
}

type BotServiceCodec interface {
	DecodeCounterRequest(ctx context.Context, update *models.Update) (*CounterRequest, error)
	EncodeCounterResponse(ctx context.Context, reply *CounterResponse) (*telegram.Message, error)
	DecodeStartRequest(ctx context.Context, update *models.Update) (*StartRequest, error)
	EncodeStartResponse(ctx context.Context, reply *StartResponse) (*telegram.Message, error)
}

func _BotService_Start0_Bot_Handler(srv BotServiceServer, codec BotServiceCodec) telegram.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) error {
		req, err := codec.DecodeStartRequest(ctx, update)
		if err != nil {
			return err
		}
		reply, err := srv.Start(ctx, req)
		if err != nil {
			return err
		}
		msg, err := codec.EncodeStartResponse(ctx, reply)
		if err != nil {
			return err
		}
		return telegram.SendMessage(ctx, b, update, msg)
	}
}

func _BotService_Counter0_Bot_Handler(srv BotServiceServer, codec BotServiceCodec) telegram.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) error {
		req, err := codec.DecodeCounterRequest(ctx, update)
		if err != nil {
			return err
		}
		reply, err := srv.Counter(ctx, req)
		if err != nil {
			return err
		}
		msg, err := codec.EncodeCounterResponse(ctx, reply)
		if err != nil {
			return err
		}
		return telegram.SendMessage(ctx, b, update, msg)
	}
}

func RegisterBotServiceBotServer(srv BotServiceServer, codec BotServiceCodec) map[string]telegram.HandlerFunc {
	handlers := make(map[string]telegram.HandlerFunc)
	handlers[BotHandlerBotServiceStart] = _BotService_Start0_Bot_Handler(srv, codec)
	handlers[BotHandlerBotServiceCounter] = _BotService_Counter0_Bot_Handler(srv, codec)
	return handlers
}
