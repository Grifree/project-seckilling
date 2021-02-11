package migrate

import sq "github.com/goclub/sql"

func (Main) Migrate20210211123026_create_table_merchant(mi sq.Migrate) {
	mi.CreateTable(sq.CreateTableQB{
		TableName: "merchant",
		PrimaryKey: []string{"id"},
		Fields: append([]sq.MigrateField{
			mi.Field("id").Char(36).DefaultString(""),
			mi.Field("name").Varchar(255).DefaultString(""),
		}, mi.CUDTimestamp()...),
		Key: map[string][]string{

		},
		Engine: mi.Engine().InnoDB,
		Charset: mi.Charset().Utf8mb4,
		Collate: mi.Utf8mb4_unicode_ci(),
	})
}
