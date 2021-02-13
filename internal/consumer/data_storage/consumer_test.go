package consumerDS

import (
	"context"
	"errors"
	connectRDS "github.com/goclub/project-seckilling/internal/connect_rds"
	IConsumerDS "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	sq "github.com/goclub/sql"
	xtest "github.com/goclub/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDS_HasConsumerByName(t *testing.T) {
	namespace := "TestDS_HasConsumerByName:" + xtest.String(10)
	ctx := context.TODO()
	rds := connectRDS.TestRDS(t)
	ds := TestDS(t)
	// 清除数据
	_, err := rds.Main.ClearTestData(ctx,sq.QB{
		Table: pd.TableConsumer{},
		Where: sq.And(pd.Consumer{}.Column().Name, sq.LikeLeft(namespace)),
	}) ; if err != nil {panic(err)}
	// 查询 nimoc
	has, reject := ds.HasConsumerByName(ctx, namespace)
	assert.Equal(t, has, false)
	assert.Equal(t, reject, nil)
	// 插入数据
	err = rds.Main.InsertModel(ctx, &pd.Consumer{
		Name: namespace,
	})
	assert.NoError(t, err)
	// 查询 nimoc
	has, reject = ds.HasConsumerByName(ctx, namespace)
	assert.Equal(t, has, true)
	assert.Equal(t, reject, nil)
}

func TestDS_CreateConsumer(t *testing.T) {
	namespace := "TestDS_CreateConsumer" + xtest.String(10)
	ctx := context.TODO()
	rds := connectRDS.TestRDS(t)
	ds := TestDS(t)
	mock := struct {
		NewConsumerID pd.IDConsumer
	}{}
	// 清除数据
	{
		_, err := rds.Main.ClearTestData(ctx,sq.QB{
			Table: pd.TableConsumer{},
			Where: sq.And(pd.Consumer{}.Column().Name, sq.LikeLeft(namespace)),
		}) ; if err != nil {panic(err)}
	}
	// 插入数据
	// {
	// 	newID,isRollback, err := ds.CreateConsumer(ctx, IConsumerDS.CreateConsumer{
	// 		Name: namespace,
	// 	}, func() error {
	// 		return nil
	// 	})
	// 	mock.NewConsumerID = newID
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, isRollback, false)
	// 	assert.Equal(t, 36, len(mock.NewConsumerID))
	// 	// 查询验证
	// 	newConsumer := pd.Consumer{}
	// 	has, err := rds.Main.QueryStruct(ctx, &newConsumer, sq.QB{
	// 		Where: sq.And(pd.Consumer{}.Column().ID, sq.Equal(mock.NewConsumerID)),
	// 	})
	// 	assert.Equal(t, has, true)
	// 	assert.Equal(t, newConsumer.Name, namespace)
	// 	assert.Equal(t, newConsumer.CreatedAt.Sub(time.Now()) < time.Second * 2, true)
	// 	assert.Equal(t, newConsumer.UpdatedAt.Sub(time.Now()) < time.Second * 2, true)
	// }
	// 插入数据
	{
		newName := namespace + ":rollback"
		_,isRollback, err := ds.CreateConsumer(ctx, IConsumerDS.CreateConsumer{
			Name: newName,
		}, func() error {
			return errors.New("mock error")
		})

		assert.EqualError(t, err, "mock error")
		assert.Equal(t, isRollback, true)
		assert.Equal(t, 0, len(mock.NewConsumerID))
		// 查询验证
		has, err := rds.Main.Has(ctx, sq.QB{
			Table: pd.TableConsumer{},
			Where: sq.And(pd.TableConsumer{}.Column().Name, sq.Equal(newName)),
		})
		assert.NoError(t ,err)
		assert.Equal(t, has, false)
	}

}

func TestDS_ConsumerHasConsumerByID(t *testing.T) {
	namespace := "TestDS_ConsumerHasConsumerByID:" + xtest.String(10)
	ctx := context.TODO()
	rds := connectRDS.TestRDS(t)
	ds := TestDS(t)
	// 清除数据
	_, err := rds.Main.ClearTestData(ctx,sq.QB{
		Table: pd.TableConsumer{},
		Where: sq.And(pd.Consumer{}.Column().Name, sq.LikeLeft(namespace)),
	}) ; if err != nil {panic(err)}
	// 查询 nimoc
	has, reject := ds.HasConsumerByName(ctx, namespace)
	assert.Equal(t, has, false)
	assert.Equal(t, reject, nil)
	// 插入数据
	consumer := pd.Consumer{
		Name: namespace,
	}
	err = rds.Main.InsertModel(ctx, &consumer)
	assert.NoError(t, err)
	// 查询 nimoc
	has, reject = ds.ConsumerHasConsumerByID(ctx, consumer.ID)
	assert.Equal(t, has, true)
	assert.Equal(t, reject, nil)
}