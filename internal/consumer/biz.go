package consumerBiz

import (
	consumerDataStorage "github.com/goclub/project-seckilling/internal/consumer/data_storage"
	IConsumerDataStorage "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	IConsumerBiz "github.com/goclub/project-seckilling/internal/consumer/interface"
	"testing"
)

type Biz struct {
	ds IConsumerDataStorage.Interface
}
func NewBiz(ds IConsumerDataStorage.Interface) IConsumerBiz.Interface {
	return Biz{
		ds: ds,
	}
}

func TestConsumer(t *testing.T) IConsumerBiz.Interface {
	return NewBiz(consumerDataStorage.TestDataStorage(t))
}