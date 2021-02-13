package goodsMS

import (
	"context"
	IGoodsMS "github.com/goclub/project-seckilling/internal/goods/memory_storage/interface"
	md "github.com/goclub/project-seckilling/internal/memory_data"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	red "github.com/goclub/redis"
)

func (dep MemoryStorage) GoodsSet(ctx context.Context, data IGoodsMS.GoodsSet) (reject error){
	goods := md.Goods{
		MerchantID: data.MerchantID,
		GoodsID: data.GoodsID,
		Title: data.Title,
		Price: data.Price,
		Description: data.Description,
		StartTime: data.StartTime,
		EndTime: data.EndTime,
		QuantityLimitPerPerson: data.QuantityLimitPerPerson,
	}
	goodsValues, reject := red.StructFieldValues(goods) ; if reject != nil {
		return
	}
	_, reject = red.HSET{
		Key: goods.RedisKey(),
		Multiple: goodsValues,
	}.Do(ctx, dep.client) ; if reject != nil {
		return
	}
}


func (dep MemoryStorage) GoodsGet(ctx context.Context, goodsID pd.IDGoods) (goods md.Goods, reject error) {
	fields, reject := red.StructFields(goods) ; if reject != nil {
		return
	}
	values, reject := red.HMGET{
		Key: goods.RedisKey(),
		Fields: fields,
	}.Do(ctx, dep.client) ; if reject != nil {
		return
	}
	reject = red.StructScan(&goods, values) ; if reject != nil {
		return
	}
	return
}
