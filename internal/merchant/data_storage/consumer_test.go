package merchantDS

import (
	"context"
	connectRDS "github.com/goclub/project-seckilling/internal/connect_rds"
	IMerchantDS "github.com/goclub/project-seckilling/internal/merchant/data_storage/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	sq "github.com/goclub/sql"
	xtest "github.com/goclub/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDS_MerchantHasMerchantByName(t *testing.T) {
	namespace := "TestDS_MerchantHasMerchantByName:" + xtest.String(10)
	ctx := context.TODO()
	rds := connectRDS.TestRDS(t)
	ds := TestDS(t)
	// 清除数据
	_, err := rds.Main.ClearTestData(ctx,sq.QB{
		Table: pd.TableMerchant{},
		Where: sq.And(pd.Merchant{}.Column().Name, sq.LikeLeft(namespace)),
	}) ; if err != nil {panic(err)}
	// 查询 nimoc
	has, reject := ds.MerchantHasMerchantByName(ctx, namespace)
	assert.Equal(t, has, false)
	assert.Equal(t, reject, nil)
	// 插入数据
	err = rds.Main.InsertModel(ctx, &pd.Merchant{
		Name: namespace,
	})
	assert.NoError(t, err)
	// 查询 nimoc
	has, reject = ds.MerchantHasMerchantByName(ctx, namespace)
	assert.Equal(t, has, true)
	assert.Equal(t, reject, nil)
}

func TestDS_MerchantCreateMerchant(t *testing.T) {
	namespace := "TestDS_MerchantCreateMerchant" + xtest.String(10)
	ctx := context.TODO()
	rds := connectRDS.TestRDS(t)
	ds := TestDS(t)
	mock := struct {
		NewMerchantID pd.IDMerchant
	}{}
	// 清除数据
	{
		_, err := rds.Main.ClearTestData(ctx,sq.QB{
			Table: pd.TableMerchant{},
			Where: sq.And(pd.Merchant{}.Column().Name, sq.LikeLeft(namespace)),
		}) ; if err != nil {panic(err)}
	}
	// 插入数据
	{
		newID, err := ds.MerchantCreateMerchant(ctx, IMerchantDS.MerchantCreateMerchant{
			Name: namespace,
		})
		mock.NewMerchantID = newID
		assert.NoError(t, err)
		assert.Equal(t, 36, len(mock.NewMerchantID))
		newMerchant := pd.Merchant{}
		has, err := rds.Main.QueryStruct(ctx, &newMerchant, sq.QB{
			Where: sq.And(pd.Merchant{}.Column().ID, sq.Equal(mock.NewMerchantID)),
		})
		assert.Equal(t, has, true)
		assert.Equal(t, newMerchant.Name, namespace)
		assert.Equal(t, newMerchant.CreatedAt.Sub(time.Now()) < time.Second * 2, true)
		assert.Equal(t, newMerchant.UpdatedAt.Sub(time.Now()) < time.Second * 2, true)
	}
	// 检查数据
	{
		merchant := pd.Merchant{}
		has, err := rds.Main.QueryStruct(ctx, &merchant, sq.QB{
			Where: sq.And(merchant.Column().ID, sq.Equal(mock.NewMerchantID)),
		})
		assert.NoError(t, err)
		assert.Equal(t, has, true)
		assert.Equal(t, merchant.CreatedAt.Sub(time.Now()).Seconds() < 2, true)
	}
}

func TestDS_MerchantHasMerchantByID(t *testing.T) {
	namespace := "TestDS_MerchantHasMerchantByID:" + xtest.String(10)
	ctx := context.TODO()
	rds := connectRDS.TestRDS(t)
	ds := TestDS(t)
	// 清除数据
	_, err := rds.Main.ClearTestData(ctx,sq.QB{
		Table: pd.TableMerchant{},
		Where: sq.And(pd.Merchant{}.Column().Name, sq.LikeLeft(namespace)),
	}) ; if err != nil {panic(err)}
	// 查询 nimoc
	has, reject := ds.MerchantHasMerchantByName(ctx, namespace)
	assert.Equal(t, has, false)
	assert.Equal(t, reject, nil)
	// 插入数据
	merchant := pd.Merchant{
		Name: namespace,
	}
	err = rds.Main.InsertModel(ctx, &merchant)
	assert.NoError(t, err)
	// 查询 nimoc
	has, reject = ds.MerchantHasMerchantByID(ctx, merchant.ID)
	assert.Equal(t, has, true)
	assert.Equal(t, reject, nil)
}