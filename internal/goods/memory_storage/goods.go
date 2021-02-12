package goodsMS

import (
	"context"
	IGoodsMS "github.com/goclub/project-seckilling/internal/goods/memory_storage/interface"
	md "github.com/goclub/project-seckilling/internal/memory_data"
	red "github.com/goclub/redis"
)

func (dep MemoryStorage) GoodsSet(ctx context.Context, data IGoodsMS.GoodsSet) (reject error){
	goods := md.Goods{
		MerchantID: data.MerchantID,
		GoodsID: data.GoodsID,
		Title: data.Title,
		Price: data.Price,
		Description: data.Description,
		QuantityLimitPerPerson: data.QuantityLimitPerPerson,
	}
	goodsValues, reject := red.StructToFieldValue(goods) ; if reject != nil {
		return 
	}
	_, reject = red.HSET{
		Key: goods.RedisKey(),
		Multiple: goodsValues,
	}.Do(ctx, dep.client) ; if reject != nil {
		return
	}
}