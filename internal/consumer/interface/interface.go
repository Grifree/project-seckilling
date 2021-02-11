package IConsumerBiz

import (
	"context"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	vd "github.com/goclub/validator"
)

type Interface interface {
	ConsumerSignIn(ctx context.Context, data ConsumerSignIn) (consumerID pd.IDConsumer, reject error)
	VerifyConsumerID(ctx context.Context, consumerID pd.IDConsumer) (reject error)
}
type ConsumerSignIn struct {
	Name string
}

func (v ConsumerSignIn) VD(r *vd.Rule) {
	r.String(v.Name, vd.StringSpec{
		Name: "用户名",
	})
}