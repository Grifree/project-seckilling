package goodsBiz

import (
	"context"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	replyU "github.com/goclub/project-seckilling/internal/util_reply"
)

func (dep Biz) OwnershipGoodsByMerchantID(ctx context.Context, goodsID pd.IDGoods, merchantID pd.IDMerchant) (reject error) {
	targetID, reject := dep.ds.MerchantIDByGoodsID(ctx, goodsID) ; if reject != nil {
		return
	}
	if targetID != merchantID {
		return replyU.RejectMessage("商品（" + goodsID.String() + ") 不属于你", true)
	}
	return
}
