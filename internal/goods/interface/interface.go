package IGoodsBiz

import (
	"context"
	"github.com/goclub/project-seckilling/internal/persistence_data"
	xtime "github.com/goclub/time"
	vd "github.com/goclub/validator"
	"time"
)

type Interface interface {
	MerchantGoodsCreate(ctx context.Context, data MerchantGoodsCreate, merchantID pd.IDMerchant) (id pd.IDGoods, reject error)
	MerchantGoodsUpdate(ctx context.Context, data MerchantGoodsUpdate, merchantID pd.IDMerchant) (reject error)
	// MerchantGoodsList(ctx context.Context, data MerchantGoodsList, merchantID pd.IDMerchant) (goodsList MerchantGoodsListReply, reject error)
	ConsumerGoods(ctx context.Context, goodsID pd.IDGoods) (goods ConsumerGoodsReply, reject error)

	OwnershipGoodsByMerchantID(ctx context.Context, goodsID pd.IDGoods, merchantID pd.IDMerchant) (reject error)
	ConsumerGoodsInventory(ctx context.Context, goodsID pd.IDGoods, consumer pd.IDConsumer) (inventory uint, reject error)

}

type MerchantGoodsCreate struct {
	Title string
	PriceCent uint64
	Description string
	StartTime xtime.ChinaTime
	EndTime xtime.ChinaTime
	QuantityLimitPerPerson uint
	Inventory uint
}
func (v MerchantGoodsCreate) VD(r *vd.Rule) {
	r.String(v.Title, vd.StringSpec{Name:"商品标题"})
	r.Uint64(v.PriceCent, vd.IntSpec{Name:"单价"})
	r.String(v.Description, vd.StringSpec{Name:"描述"})
	r.TimeRange(vd.TimeRange{
		"开始时间",v.StartTime.Time, "结束时间", v.EndTime.Time,
	})
	r.Uint(v.QuantityLimitPerPerson, vd.IntSpec{
		Name:"每人限购",
		Min: vd.Int(1),
	})
}
type MerchantGoodsUpdate struct {
	GoodsID pd.IDGoods
	Title string
	PriceCent uint64
	Description string
	StartTime xtime.ChinaTime
	EndTime xtime.ChinaTime
	QuantityLimitPerPerson uint
	Inventory uint
}

func (v MerchantGoodsUpdate) VD(r *vd.Rule) {
	r.String(v.GoodsID.String(), vd.StringSpec{
		Name:"商品ID",
		Ext: []vd.StringSpec{vd.UUID()},
	})
	r.String(v.Title, vd.StringSpec{Name:"商品标题"})
	r.Uint64(v.PriceCent, vd.IntSpec{Name:"单价"})
	r.String(v.Description, vd.StringSpec{Name:"描述"})
	r.TimeRange(vd.TimeRange{
		"开始时间",v.StartTime.Time, "结束时间", v.EndTime.Time,
	})
	r.Uint(v.QuantityLimitPerPerson, vd.IntSpec{
		Name:"每人限购",
		Min: vd.Int(1),
	})
}
type MerchantGoodsList struct {
	Search MerchantGoodsListSearch
	Page uint
	PerPage uint
}
type MerchantGoodsListSearch struct {
	Title string
	Description string
}
type MerchantGoodsListReply struct {
	Items []MerchantGoodsListReplyItem
	Total uint
}
type MerchantGoodsListReplyItem struct {
	Title string
	PriceCent uint64
	Description string
	StartTime xtime.ChinaTime
	EndTime xtime.ChinaTime
	QuantityLimitPerPerson uint
	Inventory uint
	CreateAt time.Time
	UpdateAt time.Time
}
type ConsumerGoodsReply struct {
	Title string
	PriceCent uint64
	Description string
	StartTime xtime.ChinaTime
	EndTime xtime.ChinaTime
	QuantityLimitPerPerson uint
	Inventory uint
}