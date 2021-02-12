package IGoodsMS

import (
	"context"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	"time"
)

type Interface interface {
	GoodsSet(ctx context.Context, data GoodsSet) (reject error)
}

type GoodsSet struct {
	MerchantID pd.IDMerchant
	GoodsID pd.IDGoods
	Title string
	Price uint64
	Description string
	StartTime time.Time
	EndTime time.Time
	QuantityLimitPerPerson uint
}
