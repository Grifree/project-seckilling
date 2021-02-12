package merchantBiz

import (
	"context"
	IMerchantDS "github.com/goclub/project-seckilling/internal/merchant/data_storage/interface"
	IMerchantBiz "github.com/goclub/project-seckilling/internal/merchant/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	replyU "github.com/goclub/project-seckilling/internal/util_reply"
	reqU "github.com/goclub/project-seckilling/internal/util_request"
)

func (dep Biz) MerchantSignIn(ctx context.Context, data IMerchantBiz.MerchantSignIn) (merchantID pd.IDMerchant, reject error) {
	// 格式验证
	reject = reqU.Check(data) ; if reject != nil {
		return
	}
	has, reject := dep.ds.MerchantHasMerchantByName(ctx, data.Name) ; if reject != nil {
		return
	}
	if has {
		return "", replyU.RejectMessage("用户名已存在", false)
	}
	merchantID, reject = dep.ds.MerchantCreateMerchant(ctx, IMerchantDS.MerchantCreateMerchant{
		Name: data.Name,
	}) ; if reject != nil {
		return
	}
	return
}

func (dep Biz) VerifyMerchantID(ctx context.Context, merchantID pd.IDMerchant) (reject error) {
	has, reject := dep.ds.MerchantHasMerchantByID(ctx, merchantID) ; if reject != nil {
		return
	}
	if !has {
		return replyU.RejectMessage("merchantID(" + merchantID.String() + ") 不存在", true)
	}
	return
}