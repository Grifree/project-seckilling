package migrate

import sq "github.com/goclub/sql"

func (Main) Migrate20210211123026_create_table_order(mi sq.Migrate) {
	mi.CreateTable(sq.CreateTableQB{
		TableName: "order",
		PrimaryKey: []string{"id"},
		Fields: append([]sq.MigrateField{
			mi.Field("id").Char(36).DefaultString(""),
			mi.Field("consumer_id").Char(36),
			mi.Field("goods_id").Char(36).DefaultString(""),
			mi.Field("status").Tinyint(0).DefaultInt(0),
		}, mi.CUDTimestamp()...),
		Key: map[string][]string{

		},
		Engine: mi.Engine().InnoDB,
		Charset: mi.Charset().Utf8mb4,
		Collate: mi.Utf8mb4_unicode_ci(),
	})
}
