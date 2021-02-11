package main

import (
	conf "github.com/goclub/project-seckilling/internal/config"
	connectRDS "github.com/goclub/project-seckilling/internal/connect_rds"
	"github.com/goclub/project-seckilling/internal/migrate"
	sq "github.com/goclub/sql"
)

func checkAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}
func main () {
	config, err := conf.NewConfig() ; checkAndPanic(err)
	rds, rdsClose, err := connectRDS.NewRDS(config) ; checkAndPanic(err)
	defer rdsClose()
	sq.ExecMigrate(rds.Main, &migrate.Main{})
}