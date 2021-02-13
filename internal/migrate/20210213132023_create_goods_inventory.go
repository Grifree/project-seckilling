package migrate

import sq "github.com/goclub/sql"

func (Main) Migrate20210213132023_create_goods_inventory(mi sq.Migrate) {
	mi.CreateTable(sq.CreateTableQB{
		TableName: "goods_inventory",
		PrimaryKey: []string{"goods_id"},
		Fields: append([]sq.MigrateField{
			mi.Field("goods_id").Char(36),
			mi.Field("inventory").Int(11).Unsigned(),
		}, mi.CUDTimestamp()...),
		Key: map[string][]string{

		},
		Engine: mi.Engine().InnoDB,
		Charset: mi.Charset().Utf8mb4,
		Collate: mi.Utf8mb4_unicode_ci(),
	})
}
