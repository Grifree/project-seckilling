package pd

import sq "github.com/goclub/sql"

type IDOrder string
func (id IDOrder) String() string { return string(id) }
type TableOrder struct {
	sq.SoftDeleteDeletedAt
}
func (TableOrder) TableName() string { return "order" }
func (TableOrder) Column() (col struct{
	ID sq.Column
	ConsumerID sq.Column
	GoodsID sq.Column
	Status sq.Column
	CreatedAt sq.Column
	UpdatedAt sq.Column
}){
	col.ID = "id"
	col.ConsumerID = "consumer_id"
	col.GoodsID = "goods_id"
	col.Status = "status"
	col.CreatedAt = "created_at"
	col.UpdatedAt = "updated_at"
	return
}

type Order struct {
	ID IDOrder `db:"id"`
	ConsumerID IDConsumer `db:"consumer_id"`
	GoodsID IDGoods `db:"goods_id"`
	Status OrderStatus `db:"status"`
	sq.CreatedAtUpdatedAt
}
type OrderStatus string
