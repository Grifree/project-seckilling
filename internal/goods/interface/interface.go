package IGoodsBiz

import (
	"context"
	"github.com/goclub/project-seckilling/internal/persistence_data"
	"time"
)

type Interface interface {
	MerchantGoodsCreate(ctx context.Context, data MerchantGoodsCreate) (id pd.IDGoods, reject error)
	MerchantGoodsUpdate(ctx context.Context, data MerchantGoodsUpdate) (reject error)
	MerchantGoodsList(ctx context.Context, data MerchantGoodsList) (goodsList MerchantGoodsListReply, reject error)
	MerchantGoods(ctx context.Context, MerchantGoodsID pd.IDGoods) (goods MerchantGoodsReply, reject error)
}

type MerchantGoodsCreate struct {
	Title string
	Price uint64
	Description string
	QuantityLimitPerPerson uint
}
type MerchantGoodsUpdate struct {
	GoodsID pd.IDGoods
	Title string
	Price uint64
	Description string
	QuantityLimitPerPerson uint
}
type MerchantGoodsList struct {
	Page uint
	PerPage uint
}
type MerchantGoodsListReply struct {
	Items []MerchantGoodsListReplyItem
	Total uint
}
type MerchantGoodsListReplyItem struct {
	GoodsID pd.IDGoods
	Title string
	Price uint64
	Description string
	QuantityLimitPerPerson uint
	CreateAt time.Time
	UpdateAt time.Time
}
type MerchantGoodsReply struct {
	GoodsID pd.IDGoods
	Title string
	Price uint64
	Description string
	QuantityLimitPerPerson uint
	CreateAt time.Time
	UpdateAt time.Time
}