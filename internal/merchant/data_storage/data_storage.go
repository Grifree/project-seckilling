package merchantDS

import (
	connectRDS "github.com/goclub/project-seckilling/internal/connect_rds"
	IMerchantDS "github.com/goclub/project-seckilling/internal/merchant/data_storage/interface"
	"testing"
)

type DS struct {
	rds connectRDS.RDS
}

func NewDS(rds connectRDS.RDS) IMerchantDS.Interface {
	return DS{
		rds:rds,
	}
}
func TestDS(t *testing.T)  IMerchantDS.Interface {
	rds := connectRDS.TestRDS(t)
	return NewDS(rds)
}