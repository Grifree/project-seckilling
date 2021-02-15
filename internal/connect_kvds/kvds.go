package connectKVDS

import (
	"context"
	"errors"
	conf "github.com/goclub/project-seckilling/internal/config"
	red "github.com/goclub/redis"
	"github.com/mediocregopher/radix/v4"
	"testing"
)

type KVDS struct {
	Main red.DriverRadixClient4
}

func (kvds KVDS) Close() error {
	if kvds.Main.Core == nil {
		return errors.New("ksdb.Main is nil can not close")
	}
	return kvds.Main.Core.Close()
}

func NewKVDS(config conf.Config) (kvds KVDS, closeKVDS func() error, err error)  {
	closeKVDS = func() error { return nil }
	ctx := context.Background()
	addr := config.KVDS.Host + ":" + config.KVDS.Port
	client , err := (radix.PoolConfig{}).New(ctx, config.KVDS.Network, addr) ; if err != nil {
		panic(err)
	}
	closeKVDS = kvds.Close
	kvds.Main = red.DriverRadixClient4{Core: client}
	return
}

func TestKVDS(t *testing.T, mark string) (kvds KVDS) {
	config := conf.TestConfig(t)
	var err error
	kvds, _, err = NewKVDS(config);if err != nil {
		panic(err)
	}
	return
}