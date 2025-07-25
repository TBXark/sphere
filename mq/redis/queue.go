package redis

import (
	"context"
)

type Queue[T any] struct {
	*options
}

func NewQueue[T any](opt ...Option) (*Queue[T], error) {
	opts := newOptions(opt...)
	err := opts.validate()
	if err != nil {
		return nil, err
	}
	return &Queue[T]{
		options: newOptions(opt...),
	}, nil
}

func (q *Queue[T]) Publish(ctx context.Context, topic string, data T) error {
	raw, err := q.codec.Marshal(data)
	if err != nil {
		return err
	}
	return q.client.RPush(ctx, topic, raw).Err()
}

func (q *Queue[T]) Consume(ctx context.Context, topic string) (T, error) {
	var zero T
	raw, err := q.client.LPop(ctx, topic).Bytes()
	if err != nil {
		return zero, err
	}
	var data T
	err = q.codec.Unmarshal(raw, &data)
	if err != nil {
		return zero, err
	}
	return data, nil
}

func (q *Queue[T]) PurgeQueue(ctx context.Context, topic string) error {
	return q.client.Del(ctx, topic).Err()
}

func (q *Queue[T]) Close() error {
	return q.client.Close()
}
