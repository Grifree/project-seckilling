package consumerBiz

import (
	"context"
	"fmt"
	IConsumerDS "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	IConsumerBiz "github.com/goclub/project-seckilling/internal/consumer/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	replyU "github.com/goclub/project-seckilling/internal/util_reply"
	reqU "github.com/goclub/project-seckilling/internal/util_request"
	"log"
	"time"
)

func (dep Biz) ConsumerSignIn(ctx context.Context, data IConsumerBiz.ConsumerSignIn) (consumerID pd.IDConsumer, reject error) {
	// 格式验证
	reject = reqU.Check(data) ; if reject != nil {
		return
	}
	// 查询重名
	{
		has, reject := dep.ds.HasConsumerByName(ctx, data.Name) ; if reject != nil {
			return
		};if has {
			return "", replyU.RejectMessage("用户名已存在", false)
		}
	}
	// 创建占锁
	ok, unlock, reject := dep.ms.LockConsumerCreateName(ctx, data.Name, time.Minute * 10) ; if reject != nil {
		return
	};if ok == false {
		return "", replyU.RejectMessage("用户名已存在", false)
	}
	defer func() {
		err := unlock(ctx) ; if err != nil {
			// todo: sentry log
			log.Print( fmt.Errorf("create conumser name unlock fail: name: %s consumerID %s : %w", data.Name, consumerID, err) )
		}
	}()
	consumerID, reject = dep.ds.CreateConsumer(ctx, IConsumerDS.CreateConsumer{
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