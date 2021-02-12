package goodsBiz

import (
	IGoodsDS "github.com/goclub/project-seckilling/internal/goods/data_storage/interface"
	IGoodsBiz "github.com/goclub/project-seckilling/internal/goods/interface"
	IGodsMS "github.com/goclub/project-seckilling/internal/goods/memory_storage/interface"
)

type Biz struct {
	ds IGoodsDS.Interface
	ms IGodsMS.Interface
}
func NewBiz(ds IGoodsDS.Interface,ms IGodsMS.Interface) IGoodsBiz.Interface {
	return Biz{
		ds: ds,
		ms: ms,
	}
}