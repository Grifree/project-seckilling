package consumerDataStorage

import (
	connectRDS "github.com/goclub/project-seckilling/internal/connect_rds"
	IConsumerDataStorage "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	"testing"
)

type DataStorage struct {
	rds connectRDS.RDS
}

func NewDataStorage(rds connectRDS.RDS) IConsumerDataStorage.Interface {
	return DataStorage{
		rds:rds,
	}
}
func TestDataStorage(t *testing.T)  IConsumerDataStorage.Interface {
	rds := connectRDS.TestRDS(t)
	return NewDataStorage(rds)
}