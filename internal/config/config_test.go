package conf

import (
	"log"
	"testing"
)

func TestNewConfig(t *testing.T) {
	log.Print(NewConfig())
}
