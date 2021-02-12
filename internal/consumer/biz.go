package consumerBiz

import (
	consumerDS "github.com/goclub/project-seckilling/internal/consumer/data_storage"
	IConsumerDS "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	IConsumerBiz "github.com/goclub/project-seckilling/internal/consumer/interface"
	"testing"
)

type Biz struct {
	ds IConsumerDS.Interface
}
func NewBiz(ds IConsumerDS.Interface) IConsumerBiz.Interface {
	return Biz{
		ds: ds,
	}
}

func TestConsumer(t *testing.T) IConsumerBiz.Interface {
	return NewBiz(consumerDS.TestDS(t))
}