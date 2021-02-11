package migrate

import sq "github.com/goclub/sql"

func (Main) Migrate20210211123026_create_table_goods(mi sq.Migrate) {
	mi.CreateTable(sq.CreateTableQB{
		TableName: "goods",
		PrimaryKey: []string{"id"},
		Fields: append([]sq.MigrateField{
			mi.Field("id").Char(36).DefaultString(""),
			mi.Field("merchant_id").Char(36).DefaultString(""),
			mi.Field("title").Varchar(255).DefaultString(""),
			mi.Field("price").Int(11).DefaultInt(0),
			mi.Field("description").Type("text", 0),
			mi.Field("start_time").Timestamp().DefaultCurrentTimeStamp(),
			mi.Field("end_time").Timestamp().DefaultCurrentTimeStamp(),
			mi.Field("quantity_limit_per_person").Tinyint(0),
		}, mi.CUDTimestamp()...),
		Key: map[string][]string{

		},
		Engine: mi.Engine().InnoDB,
		Charset: mi.Charset().Utf8mb4,
		Collate: mi.Utf8mb4_unicode_ci(),
	})
}
