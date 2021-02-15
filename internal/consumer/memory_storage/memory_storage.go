package consumerMS

import connectKVDS "github.com/goclub/project-seckilling/internal/connect_kvds"

type MemoryStorage struct {
	kvds connectKVDS.KVDS
}
