package pd

import sq "github.com/goclub/sql"

type IDMerchant string
func (id IDMerchant) String() string { return string(id) }
type TableMerchant struct {
	sq.SoftDeleteDeletedAt
}
func (TableMerchant) TableName() string {
	return "merchant"
}

func (TableMerchant) Column() (col struct{
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
type Merchant struct {
	TableMerchant
	sq.DefaultLifeCycle

	ID IDMerchant `db:"id"`
	Name string `db:"name"`

	sq.CreatedAtUpdatedAt
}
func (row *Merchant) BeforeCreate() error {
	if len(row.ID) == 0 {
		row.ID = IDMerchant(sq.UUID())
	}
	return nil
}