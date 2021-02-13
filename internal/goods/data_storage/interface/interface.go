package IGoodsDS

import (
	"context"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	sq "github.com/goclub/sql"
	"time"
)

type Interface interface {
	Transaction(ctx context.Context, handle func (tx *sq.Transaction) sq.TxResult) (isRollback bool, err error)
	GoodsCreate(ctx context.Context, data GoodsCreate) (goodsID pd.IDGoods, reject error)
	GoodsUpdate(ctx context.Context, data GoodsUpdate) (reject error)
	MerchantIDByGoodsID(ctx context.Context, goodsID pd.IDGoods) (merchantID pd.IDMerchant, reject error)

}

type GoodsCreate struct {
	MerchantID pd.IDMerchant
	Title string
	PriceCent uint64
	Description string
	StartTime time.Time
	EndTime time.Time
	QuantityLimitPerPerson uint
	Inventory uint
}
type GoodsUpdate struct {
	GoodsID pd.IDGoods
	Title string
	PriceCent uint64
	Description string
	StartTime time.Time
	EndTime time.Time
	QuantityLimitPerPerson uint
	Inventory uint
}