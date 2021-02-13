package consumerBiz

import (
	"context"
	IConsumerDS "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	IConsumerBiz "github.com/goclub/project-seckilling/internal/consumer/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	replyU "github.com/goclub/project-seckilling/internal/util_reply"
	reqU "github.com/goclub/project-seckilling/internal/util_request"
	"time"
)

func (dep Biz) ConsumerSignIn(ctx context.Context, data IConsumerBiz.ConsumerSignIn) (consumerID pd.IDConsumer, reject error) {
	// 格式验证
	reject = reqU.Check(data) ; if reject != nil {
		return
	}
	// 查询重名
	has, reject := dep.ds.HasConsumerByName(ctx, data.Name) ; if reject != nil {
		return
	}
	if has {
		return "", replyU.RejectMessage("用户名已存在", false)
	}
	ok, unlock, reject := dep.ms.LockConsumerCreateName(data.Name, time.Second * 10) ; if reject != nil {
		return
	}
	unlockCtx, cancelCtx := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelCtx()
	if ok == false {
		return "", replyU.RejectMessage("用户名已存在", false)
	}
	// 有可能以后允许出现重名用户，所以不在 sql 做 unique 约束
	var execUnlock = func() error {
		return unlock(unlockCtx)
	}
	var isRollback bool
	consumerID,isRollback, reject = dep.ds.CreateConsumer(ctx, IConsumerDS.CreateConsumer{
		Name: data.Name,
	}, execUnlock) ; if reject != nil {
		return
	}
	if isRollback == false {
		return "", replyU.RejectMessage("用户名已存在或其他错误", false)
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