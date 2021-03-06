package goodsBiz

import (
	"context"
	IGoodsBiz "github.com/goclub/project-seckilling/internal/goods/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	xtime "github.com/goclub/time"
)

func (dep Biz) ConsumerGoods(ctx context.Context, goodsID pd.IDGoods) (goods IGoodsBiz.ConsumerGoodsReply, reject error) {
	// 合法性验证:暂无
	// 读取数据
	// 为了降低压力 读取 redis 数据
	data, reject := dep.ms.GoodsGet(ctx, goodsID) ; if reject != nil {
		return
	}
	goods = IGoodsBiz.ConsumerGoodsReply{
		Title: data.Title,
		PriceCent: data.PriceCent,
		Description: data.Description,
		StartTime: xtime.NewChinaTime(data.StartTime),
		EndTime: xtime.NewChinaTime(data.EndTime),
		QuantityLimitPerPerson: data.QuantityLimitPerPerson,
		Inventory: data.Inventory,
	}
	return
}

