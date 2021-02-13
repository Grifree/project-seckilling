package md

import (
	"strings"
)

type LockConsumerCreateName struct {
	Name string
}
func (data LockConsumerCreateName) RedisKey() string {
	return strings.Join([]string{"lock_consumer_create_name", data.Name}, ":")
}
