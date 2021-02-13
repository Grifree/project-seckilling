package consumerBiz

import (
	"context"
	xerr "github.com/goclub/error"
	connectRDS "github.com/goclub/project-seckilling/internal/connect_rds"
	consumerDS "github.com/goclub/project-seckilling/internal/consumer/data_storage"
	IConsumerDS "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	IConsumerBiz "github.com/goclub/project-seckilling/internal/consumer/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	replyU "github.com/goclub/project-seckilling/internal/util_reply"
	sq "github.com/goclub/sql"
	xtest "github.com/goclub/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBiz_ConsumerSignIn(t *testing.T) {
	namespace := "TestBiz_ConsumerSignIn:" + xtest.String(10)
	ctx := context.TODO()
	rds := connectRDS.TestRDS(t)
	ds := consumerDS.TestDS(t)
	biz := TestConsumer(t)
	mock := struct {
		NewConsumerID pd.IDConsumer
	}{}
	// 清除数据
	_, err := rds.Main.ClearTestData(ctx,sq.QB{
		Table: pd.TableConsumer{},
		Where: sq.And(pd.Consumer{}.Column().Name, sq.LikeLeft(namespace)),
	}) ; if err != nil {panic(err)}
	// 查询 不存在
	has, reject := ds.HasConsumerByName(ctx, namespace)
	assert.Equal(t, has, false)
	assert.Equal(t, reject, nil)
	// 注册
	newID, err := biz.ConsumerSignIn(ctx, IConsumerBiz.ConsumerSignIn{
		Name: namespace,
	})
	mock.NewConsumerID = newID
	assert.NoError(t, err)
	assert.Equal(t, len(mock.NewConsumerID), 36)
	// 查询 存在
	has, reject = ds.HasConsumerByName(ctx, namespace)
	assert.Equal(t, has, true)
	assert.Equal(t, reject, nil)
	// 重复注册
	_, err = biz.ConsumerSignIn(ctx, IConsumerBiz.ConsumerSignIn{
		Name: namespace,
	})
	reject, asReject := xerr.AsReject(err)
	assert.Equal(t, asReject, true)
	assert.Equal(t, reject, replyU.RejectMessage("用户名已存在", false))
	// 查询数据只新增了1个
	count, err := rds.Main.Count(ctx, sq.QB{
		Table: pd.TableConsumer{},
		Where: sq.And(pd.TableConsumer{}.Column().Name, sq.Equal(namespace)),
	})
	assert.NoError(t, err)
	assert.Equal(t, count, uint64(1))
}

func TestBiz_VerifyConsumerID(t *testing.T) {
	namespace := "TestBiz_VerifyConsumerID:" + xtest.String(10)
	ctx := context.TODO()
	rds := connectRDS.TestRDS(t)
	ds := consumerDS.TestDS(t)
	_=ds
	biz := TestConsumer(t)
	// 清除数据
	_, err := rds.Main.ClearTestData(ctx,sq.QB{
		Table: pd.TableConsumer{},
		Where: sq.And(pd.Consumer{}.Column().Name, sq.LikeLeft(namespace)),
	}) ; if err != nil {panic(err)}
	// 查询 错误的id
	invalidID := pd.IDConsumer(sq.UUID())
	err = biz.VerifyConsumerID(ctx, invalidID)
	assert.Equal(t, err, replyU.RejectMessage("consumerID(" + invalidID.String() + ") 不存在", true))
	// 插入数据
	newConsumerID, err := ds.CreateConsumer(ctx, IConsumerDS.CreateConsumer{
		Name: namespace,
	})
	assert.NoError(t ,err)
	// 查询正确id
	err = biz.VerifyConsumerID(ctx, newConsumerID)
	assert.NoError(t, err)
}