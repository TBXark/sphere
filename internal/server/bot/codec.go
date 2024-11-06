package bot

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot/models"
	botv1 "github.com/tbxark/sphere/api/bot/v1"
	"github.com/tbxark/sphere/pkg/telegram"
	"math/rand"
)

var _ botv1.CounterServiceCodec[models.Update, telegram.Message] = &CounterServiceCodec{}

const (
	CommandStart   = "/start"
	CommandCounter = "/counter"
)

const (
	QueryCounter = "counter"
)

type CounterServiceCodec struct{}

func (b *CounterServiceCodec) DecodeCounterRequest(ctx context.Context, update *models.Update) (*botv1.CounterRequest, error) {
	value := UnmarshalUpdateDataWithDefault(update, 0)
	return &botv1.CounterRequest{
		Count: int32(value),
	}, nil
}

func (b *CounterServiceCodec) EncodeCounterResponse(ctx context.Context, reply *botv1.CounterResponse) (*telegram.Message, error) {
	return &telegram.Message{
		Text: fmt.Sprintf("Counter: %d", reply.Count),
		Button: [][]telegram.Button{
			{
				NewButton("Increment", QueryCounter, reply.Count+1),
				NewButton("Decrement", QueryCounter, reply.Count-1),
			},
			{
				NewButton("Reset", QueryCounter, 0),
				NewButton("Random", QueryCounter, rand.Int()%100),
			},
		},
	}, nil
}

func (b *CounterServiceCodec) DecodeStartRequest(ctx context.Context, update *models.Update) (*botv1.StartRequest, error) {
	return &botv1.StartRequest{
		Name: update.Message.From.FirstName,
	}, nil
}

func (b *CounterServiceCodec) EncodeStartResponse(ctx context.Context, reply *botv1.StartResponse) (*telegram.Message, error) {
	return &telegram.Message{
		Text: reply.Message,
	}, nil
}
