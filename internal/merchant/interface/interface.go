package IMerchantBiz

import (
	"context"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	vd "github.com/goclub/validator"
)

type Interface interface {
	MerchantSignIn(ctx context.Context, data MerchantSignIn) (merchantID pd.IDMerchant, reject error)
	VerifyMerchantID(ctx context.Context, merchantID pd.IDMerchant) (reject error)
}
type MerchantSignIn struct {
	Name string
}

func (v MerchantSignIn) VD(r *vd.Rule) {
	r.String(v.Name, vd.StringSpec{
		Name: "用户名",
	})
}