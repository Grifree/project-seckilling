package goodsBiz

import (
	"context"
	IGoodsBiz "github.com/goclub/project-seckilling/internal/goods/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
)

func (dep Biz) ConsumerGoods(ctx context.Context, consumerID pd.IDConsumer) (goods IGoodsBiz.ConsumerGoodsReply, reject error) {
	return
}
