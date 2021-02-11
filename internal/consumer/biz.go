package consumerBiz

import (
	IConsumerDataStorage "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	IConsumerBiz "github.com/goclub/project-seckilling/internal/consumer/interface"
)

type Biz struct {
	ds IConsumerDataStorage.Interface
}
func NewBiz(ds IConsumerDataStorage.Interface) IConsumerBiz.Interface {
	return Biz{
		ds: ds,
	}
}
