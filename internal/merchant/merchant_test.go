package merchantBiz

import (
	"context"
	xerr "github.com/goclub/error"
	connectRDS "github.com/goclub/project-seckilling/internal/connect_rds"
	merchantDS "github.com/goclub/project-seckilling/internal/merchant/data_storage"
	IMerchantDS "github.com/goclub/project-seckilling/internal/merchant/data_storage/interface"
	IMerchantBiz "github.com/goclub/project-seckilling/internal/merchant/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	replyU "github.com/goclub/project-seckilling/internal/util_reply"
	sq "github.com/goclub/sql"
	xtest "github.com/goclub/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBiz_MerchantSignIn(t *testing.T) {
	namespace := "TestBiz_MerchantSignIn:" + xtest.String(10)
	ctx := context.TODO()
	rds := connectRDS.TestRDS(t)
	ds := merchantDS.TestDS(t)
	biz := TestMerchant(t)
	mock := struct {
		NewMerchantID pd.IDMerchant
	}{}
	// 清除数据
	_, err := rds.Main.ClearTestData(ctx,sq.QB{
		Table: pd.TableMerchant{},
		Where: sq.And(pd.Merchant{}.Column().Name, sq.LikeLeft(namespace)),
	}) ; if err != nil {panic(err)}
	// 查询 不存在
	has, reject := ds.MerchantHasMerchantByName(ctx, namespace)
	assert.Equal(t, has, false)
	assert.Equal(t, reject, nil)
	// 注册
	newID, err := biz.MerchantSignIn(ctx, IMerchantBiz.MerchantSignIn{
		Name: namespace,
	})
	mock.NewMerchantID = newID
	assert.NoError(t, err)
	assert.Equal(t, len(mock.NewMerchantID), 36)
	// 查询 存在
	has, reject = ds.MerchantHasMerchantByName(ctx, namespace)
	assert.Equal(t, has, true)
	assert.Equal(t, reject, nil)
	// 重复注册
	_, err = biz.MerchantSignIn(ctx, IMerchantBiz.MerchantSignIn{
		Name: namespace,
	})
	reject, asReject := xerr.AsReject(err)
	assert.Equal(t, asReject, true)
	assert.Equal(t, reject, replyU.RejectMessage("用户名已存在", false))
	// 查询数据只新增了1个
	count, err := rds.Main.Count(ctx, sq.QB{
		Table: pd.TableMerchant{},
		Where: sq.And(pd.TableMerchant{}.Column().Name, sq.Equal(namespace)),
	})
	assert.NoError(t, err)
	assert.Equal(t, count, uint64(1))
}

func TestBiz_VerifyMerchantID(t *testing.T) {
	namespace := "TestBiz_VerifyMerchantID:" + xtest.String(10)
	ctx := context.TODO()
	rds := connectRDS.TestRDS(t)
	ds := merchantDS.TestDS(t)
	_=ds
	biz := TestMerchant(t)
	// 清除数据
	_, err := rds.Main.ClearTestData(ctx,sq.QB{
		Table: pd.TableMerchant{},
		Where: sq.And(pd.Merchant{}.Column().Name, sq.LikeLeft(namespace)),
	}) ; if err != nil {panic(err)}
	// 查询 错误的id
	invalidID := pd.IDMerchant(sq.UUID())
	err = biz.VerifyMerchantID(ctx, invalidID)
	assert.Equal(t, err, replyU.RejectMessage("merchantID(" + invalidID.String() + ") 不存在", true))
	// 插入数据
	newMerchantID, err := ds.MerchantCreateMerchant(ctx, IMerchantDS.MerchantCreateMerchant{
		Name: namespace,
	})
	assert.NoError(t ,err)
	// 查询正确id
	err = biz.VerifyMerchantID(ctx, newMerchantID)
	assert.NoError(t, err)
}
