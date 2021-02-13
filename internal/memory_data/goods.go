package md

import (
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	"strings"
	"time"
)

type Goods struct {
	MerchantID pd.IDMerchant `red:"merchant_id"`
	GoodsID pd.IDGoods `red:"goods_id"`
	Title string `red:"title"`
	PriceCent uint64 `red:"price_cent"`
	Description string `red:"description"`
	StartTime time.Time `red:"start_time"`
	EndTime time.Time `red:"end_time"`
	QuantityLimitPerPerson uint `red:"quantity_limit_per_person"`
	Inventory uint `red:"inventory"`
}
func (data Goods) RedisKey () string {
	return strings.Join([]string{"goods", data.GoodsID.String()}, ":")
}