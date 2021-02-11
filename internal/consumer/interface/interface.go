package IConsumerBiz

import (
	"context"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
)

type Interface interface {
	ConsumerSignIn(ctx context.Context, data ConsumerSignIn) (consumerID pd.IDConsumer, reject error)
}
type ConsumerSignIn struct {
	Name string
}
