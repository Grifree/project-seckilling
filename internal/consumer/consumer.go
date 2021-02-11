package consumerBiz

import (
	"context"
	IConsumerDataStorage "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	IConsumerBiz "github.com/goclub/project-seckilling/internal/consumer/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	replyU "github.com/goclub/project-seckilling/internal/util_reply"
	reqU "github.com/goclub/project-seckilling/internal/util_request"
)

func (dep Biz) ConsumerSignIn(ctx context.Context, data IConsumerBiz.ConsumerSignIn) (consumerID pd.IDConsumer, reject error) {
	// 格式验证
	reject = reqU.Check(data) ; if reject != nil {
		return
	}
	has, reject := dep.ds.ConsumerHasConsumerByName(ctx, data.Name) ; if reject != nil {
		return
	}
	if has {
		return "", replyU.RejectMessage("用户名已存在", false)
	}
	consumerID, reject = dep.ds.ConsumerCreateConsumer(ctx, IConsumerDataStorage.ConsumerCreateConsumer{
		Name: data.Name,
	}) ; if reject != nil {
		return
	}
	return
}

func (dep Biz) VerifyConsumerID(ctx context.Context, consumerID pd.IDConsumer) (reject error) {
	has, reject := dep.ds.ConsumerHasConsumerByID(ctx, consumerID) ; if reject != nil {
		return
	}
	if !has {
		return replyU.RejectMessage("consumerID(" + consumerID.String() + ") 不存在", true)
	}
	return
}