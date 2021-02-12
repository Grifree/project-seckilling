package consumerDS

import (
	connectRDS "github.com/goclub/project-seckilling/internal/connect_rds"
	IConsumerDS "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	"testing"
)

type DS struct {
	rds connectRDS.RDS
}

func NewDS(rds connectRDS.RDS) IConsumerDS.Interface {
	return DS{
		rds:rds,
	}
}
func TestDS(t *testing.T)  IConsumerDS.Interface {
	rds := connectRDS.TestRDS(t)
	return NewDS(rds)
}