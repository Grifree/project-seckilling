package consumerBiz

import (
	consumerDS "github.com/goclub/project-seckilling/internal/consumer/data_storage"
	IConsumerDS "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	IConsumerBiz "github.com/goclub/project-seckilling/internal/consumer/interface"
	IConsumerMS "github.com/goclub/project-seckilling/internal/consumer/memory_storage/interface"
	"testing"
)

type Biz struct {
	ds IConsumerDS.Interface
	ms IConsumerMS.Interface
}
func NewBiz(ds IConsumerDS.Interface, ms IConsumerMS.Interface) IConsumerBiz.Interface {
	return Biz{
		ds: ds,
		ms: ms,
	}
}

func TestConsumer(t *testing.T) IConsumerBiz.Interface {
	return NewBiz(consumerDS.TestDS(t))
}