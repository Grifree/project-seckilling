package reqU

import (
	replyU "github.com/goclub/project-seckilling/internal/util_reply"
	vd "github.com/goclub/validator"
)

var checker vd.Checker
func init () {
	checker = vd.NewCN()
}

func Check(data vd.Data) (reject error) {
	report := checker.Check(data)
	if report.Fail {
		return replyU.RejectMessage(report.Message, false)
	}
	return
}