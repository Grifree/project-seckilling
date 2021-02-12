package merchantBiz

import (
	merchantDS "github.com/goclub/project-seckilling/internal/merchant/data_storage"
	IMerchantDS "github.com/goclub/project-seckilling/internal/merchant/data_storage/interface"
	IMerchantBiz "github.com/goclub/project-seckilling/internal/merchant/interface"
	"testing"
)

type Biz struct {
	ds IMerchantDS.Interface
}
func NewBiz(ds IMerchantDS.Interface) IMerchantBiz.Interface {
	return Biz{
		ds: ds,
	}
}

func TestMerchant(t *testing.T) IMerchantBiz.Interface {
	return NewBiz(merchantDS.TestDS(t))
}