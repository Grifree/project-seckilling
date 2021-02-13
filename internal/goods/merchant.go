package goodsBiz

import (
	"context"
	IGoodsDS "github.com/goclub/project-seckilling/internal/goods/data_storage/interface"
	IGoodsBiz "github.com/goclub/project-seckilling/internal/goods/interface"
	IGoodsMS "github.com/goclub/project-seckilling/internal/goods/memory_storage/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	reqU "github.com/goclub/project-seckilling/internal/util_request"
)

func (dep Biz) MerchantGoodsCreate(ctx context.Context, data IGoodsBiz.MerchantGoodsCreate, merchantID pd.IDMerchant) (goodsID pd.IDGoods, reject error) {
	// 格式验证
	reject = reqU.Check(data) ; if reject != nil {
		return
	}
	// 合法性验证：暂无
	// 插入数据
	create := IGoodsDS.GoodsCreate{
		MerchantID: merchantID,
		Title: data.Title,
		Price: data.Price,
		Description: data.Description,
		StartTime: data.StartTime.Time,
		EndTime: data.EndTime.Time,
		QuantityLimitPerPerson: data.QuantityLimitPerPerson,
		Inventory: data.Inventory,
	}
	goodsID, reject = dep.ds.GoodsCreate(ctx, create) ; if reject != nil {
		return
	}
	reject = dep.ms.GoodsSet(ctx,IGoodsMS.GoodsSet{
		MerchantID: merchantID,
		GoodsID: goodsID,
		Title: create.Title,
		Price: create.Price,
		Description: create.Description,
		StartTime: create.StartTime,
		EndTime: create.EndTime,
		QuantityLimitPerPerson: create.QuantityLimitPerPerson,
	}) ; if reject != nil {
		return
	}
	return
}
func (dep Biz) MerchantGoodsUpdate(ctx context.Context, data IGoodsBiz.MerchantGoodsUpdate, merchantID pd.IDMerchant) (reject error) {
	// 格式校验
	reject = reqU.Check(data) ; if reject != nil {
		return
	}
	// 合法性验证
	reject = dep.OwnershipGoodsByMerchantID(ctx, data.GoodsID, merchantID) ;if reject != nil {
		return
	}
	// 更新数据
	update :=  IGoodsDS.GoodsUpdate{
		GoodsID: data.GoodsID,
		Title: data.Title,
		Price: data.Price,
		Description: data.Description,
		StartTime: data.StartTime.Time,
		EndTime: data.EndTime.Time,
		QuantityLimitPerPerson: data.QuantityLimitPerPerson,
		Inventory: data.Inventory,
	}
	reject = dep.ds.GoodsUpdate(ctx,update) ; if reject != nil {
		return
	}
	reject = dep.ms.GoodsSet(ctx,IGoodsMS.GoodsSet{
		MerchantID: merchantID,
		GoodsID: update.GoodsID,
		Title: update.Title,
		Price: update.Price,
		Description: update.Description,
		StartTime: update.StartTime,
		EndTime: update.EndTime,
		QuantityLimitPerPerson: update.QuantityLimitPerPerson,
		Inventory: update.Inventory,
	}) ; if reject != nil {
		return
	}
	// 将数据写入缓存
	return
}
// func (dep Biz) MerchantGoodsList(ctx context.Context, data IGoodsBiz.MerchantGoodsList, merchantID pd.IDMerchant) (goodsList IGoodsBiz.MerchantGoodsListReply, reject error) {
//
// }
