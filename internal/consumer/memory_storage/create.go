package consumerMS

import (
	"context"
	md "github.com/goclub/project-seckilling/internal/memory_data"
	red "github.com/goclub/redis"
	"time"
)

func (dep MemoryStorage) LockConsumerCreateName(ctx context.Context, name string, expire time.Duration)(ok bool, unlock func(ctx context.Context) (err error), reject error) {
	consumerLock := md.LockConsumerCreateName{
		Name:name,
	}
	consumerLockKey := consumerLock.RedisKey()
	mutex := red.Mutex{
		Key:consumerLockKey,
		Expire:expire,
	}
	unlock = mutex.Unlock
	ok, reject = mutex.Lock(ctx,dep.kvds.Main); if reject != nil {
		return
	}
	return
}