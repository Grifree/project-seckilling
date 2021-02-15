package connectRDS

import (
	_ "github.com/go-sql-driver/mysql"
	conf "github.com/goclub/project-seckilling/internal/config"
	sq "github.com/goclub/sql"
	"log"
	"testing"
)

type RDS struct {
	Main *sq.Database
}
func (rds RDS) Close() error {
	return rds.Main.Close()
}
func NewRDS(config conf.Config) (rds RDS, rdsClose func() error, err error) {
	mainDB, _, err := sq.Open("mysql", config.RDS.String()) ; if err != nil {
		return
	}
	rds.Main = mainDB
	rdsClose = rds.Close
	return
}

func TestRDS(t *testing.T) (rds RDS) {
	config := conf.TestConfig(t)
	var err error
	rds, _, err = NewRDS(config) ; if err != nil {
		panic(err)
	}
	log.Print(rds.Main.Ping())
	return rds
}