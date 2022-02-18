package protocol

import (
	"context"
)

type Store interface {
	Get(ctx context.Context, key []byte) ([]byte, error)
	Set(ctx context.Context, key, value []byte) error
	Update(ctx context.Context, key, value []byte) error
	Delete(ctx context.Context, key []byte) error
	Keys(ctx context.Context) <-chan []byte
}
