package pd

import sq "github.com/goclub/sql"

type IDConsumer string
func (id IDConsumer) String() string { return string(id) }
type TableConsumer struct {
	sq.SoftDeleteDeletedAt
}
func (TableConsumer) TableName() string {
	return "consumer"
}
func (TableConsumer) Column() (col struct{
	ID sq.Column
	Name sq.Column
	CreatedAt sq.Column
	UpdatedAt sq.Column
}) {
	col.ID = "id"
	col.Name ="name"
	col.CreatedAt = "created_at"
	col.UpdatedAt = "updated_at"
	return
}
type Consumer struct {
	TableConsumer
	sq.DefaultLifeCycle

	ID IDConsumer `db:"id"`
	Name string `db:"name"`

	sq.CreatedAtUpdatedAt
}
func (row *Consumer) BeforeCreate() error {
	if len(row.ID) == 0 {
		row.ID = IDConsumer(sq.UUID())
	}
	return nil
}