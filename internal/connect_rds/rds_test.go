package connectRDS_test

import (
	connectRDS "github.com/goclub/project-seckilling/internal/connect_rds"
	xtest "github.com/goclub/test"
	"log"
	"testing"
	"time"
)

func TestTestRDS(t *testing.T) {
	err := xtest.Run(10, func(i int) (op xtest.RunOp) {
		log.Print(i)
		connectRDS.TestRDS(t)
		time.Sleep(time.Second/2)
		//if i == 5 {
		//	panic("panic 5")
		//}
		return
	});if err != nil {
		panic(err)
	}
}
