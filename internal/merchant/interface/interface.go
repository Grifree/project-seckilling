package IMerchantBiz

import (
	"context"
	vd "github.com/goclub/validator"
)

type Interface interface {
	ConsumerSignIn(ctx context.Context, data ConsumerSignIn) (reject error)
}
type ConsumerSignIn struct {
	Name string
}

func (v ConsumerSignIn) VD(r *vd.Rule) {
	r.String(v.Name, vd.StringSpec{
		Name: "商户名",
	})
}
