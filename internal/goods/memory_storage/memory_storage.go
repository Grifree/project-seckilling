package goodsMS

import red "github.com/goclub/redis"

type MemoryStorage struct {
	client red.DriverRadixClient4
}