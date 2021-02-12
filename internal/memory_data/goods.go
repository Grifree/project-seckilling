package md

import (
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	"strings"
)

type Goods struct {
	MerchantID pd.IDMerchant `redis:"merchant_id"`
	GoodsID pd.IDGoods `redis:"goods_id"`
	Title string `redis:"title"`
	Price uint64 `redis:"price"`
	Description string `redis:"description"`
	// StartTime time.Time
	// EndTime time.Time
	QuantityLimitPerPerson uint `redis:"quantity_limit_per_person"`
}
func (data Goods) RedisKey () string {
	return strings.Join([]string{"goods", data.GoodsID.String()}, ":")
}