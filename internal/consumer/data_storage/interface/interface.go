package IConsumerDS

import (
	"context"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
)

type Interface interface {
	HasConsumerByName(ctx context.Context, name string) (has bool, reject error)
	CreateConsumer(ctx context.Context, data CreateConsumer) (consumerID pd.IDConsumer, reject error)
	ConsumerHasConsumerByID(ctx context.Context, consumerID pd.IDConsumer) (has bool, reject error)
}
type CreateConsumer struct {
	Name string
}