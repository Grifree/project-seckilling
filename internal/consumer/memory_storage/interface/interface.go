package IConsumerMS

import (
	"context"
	"time"
)

type Interface interface {
	LockConsumerCreateName(ctx context.Context, name string, expire time.Duration)(ok bool, unlock func(ctx context.Context) (err error)  ,reject error)
}

