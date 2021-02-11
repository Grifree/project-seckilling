package IConsumerDataStorage

import (
	"context"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
)

type Interface interface {
	ConsumerHasConsumerByName(ctx context.Context, name string) (has bool, reject error)
	ConsumerCreateConsumer(ctx context.Context, data ConsumerCreateConsumer) (consumerID pd.IDConsumer, reject error)
	ConsumerHasConsumerByID(ctx context.Context, consumerID pd.IDConsumer) (has bool, reject error)
}
type ConsumerCreateConsumer struct {
	Name string
}