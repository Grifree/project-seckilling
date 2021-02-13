package pd

import (
	sq "github.com/goclub/sql"
	"time"
)

type IDGoods string
func (id IDGoods) String() string {
	return string(id)
}

type TableGoods struct {
	sq.SoftDeleteDeletedAt
}
func (TableGoods)  TableName () string {
	return "goods"
}


func (TableGoods) Column() (col struct{
	ID sq.Column
	MerchantID sq.Column
	Title sq.Column
	PriceCent sq.Column
	Description sq.Column
	StartTime sq.Column
	EndTime sq.Column
	QuantityLimitPerPerson sq.Column

	CreatedAt sq.Column
	UpdatedAt sq.Column
}) {
	col.ID = "id"
	col.MerchantID  = "merchant_id"
	col.Title = "title"
	col.PriceCent = "price_cent"
	col.Description = "description"
	col.StartTime = "start_time"
	col.EndTime = "end_time"
	col.QuantityLimitPerPerson= "quantity_limit_per_person"
	col.CreatedAt = "created_at"
	col.UpdatedAt = "updated_at"
	return
}
type Goods struct {
	ID IDGoods `db:"id"`
	MerchantID IDMerchant `db:"merchant_id"`
	Title string `db:"title"`
	PriceCent uint `db:"price_cent"`
	Description string `db:"description"`
	StartTime time.Time `db:"start_time"`
	EndTime time.Time `db:"end_time"`
	QuantityLimitPerPerson uint `db:"quantity_limit_per_person"`
	sq.CreatedAtUpdatedAt
}

type TableGoodsInventory struct {
	sq.SoftDeleteDeletedAt
}
func (TableGoodsInventory) TableName() string {
	return "goods_inventory"
}
func (TableGoodsInventory) Column() (col struct{
	GoodsID sq.Column
	Inventory sq.Column

	CreatedAt sq.Column
	UpdatedAt sq.Column
}) {

	col.GoodsID = "goods_id"
	col.Inventory = "inventory"

	col.CreatedAt = "created_at"
	col.UpdatedAt = "updated_at"
	return
}


func (data GoodsInventory) WherePrimaryKey() []sq.Condition {
	return sq.And(data.Column().GoodsID, sq.Equal(data.GoodsID))
}
type GoodsInventory struct {
	TableGoodsInventory
	GoodsID IDGoods `db:"goods_id"`
	Inventory uint `db:"inventory"`
	sq.CreatedAtUpdatedAt
}
