package consumerDataStorage

import (
	"context"
	connectRDS "github.com/goclub/project-seckilling/internal/connect_rds"
	IConsumerDataStorage "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	sq "github.com/goclub/sql"
	xtest "github.com/goclub/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDataStorage_ConsumerHasConsumerByName(t *testing.T) {
	namespace := "TestDataStorage_ConsumerHasConsumerByName:" + xtest.String(10)
	ctx := context.TODO()
	rds := connectRDS.TestRDS(t)
	ds := TestDataStorage(t)
	// 清除数据
	_, err := rds.Main.ClearTestData(ctx,sq.QB{
		Table: pd.TableConsumer{},
		Where: sq.And(pd.Consumer{}.Column().Name, sq.LikeLeft(namespace)),
	}) ; if err != nil {panic(err)}
	// 查询 nimoc
	has, reject := ds.ConsumerHasConsumerByName(ctx, namespace)
	assert.Equal(t, has, false)
	assert.Equal(t, reject, nil)
	// 插入数据
	err = rds.Main.InsertModel(ctx, &pd.Consumer{
		Name: namespace,
	})
	assert.NoError(t, err)
	// 查询 nimoc
	has, reject = ds.ConsumerHasConsumerByName(ctx, namespace)
	assert.Equal(t, has, true)
	assert.Equal(t, reject, nil)
}

func TestDataStorage_ConsumerCreateConsumer(t *testing.T) {
	namespace := "TestDataStorage_ConsumerCreateConsumer" + xtest.String(10)
	ctx := context.TODO()
	rds := connectRDS.TestRDS(t)
	ds := TestDataStorage(t)
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
	{
		newID, err := ds.ConsumerCreateConsumer(ctx, IConsumerDataStorage.ConsumerCreateConsumer{
			Name: namespace,
		})
		mock.NewConsumerID = newID
		assert.NoError(t, err)
		assert.Equal(t, 36, len(mock.NewConsumerID))
		newConsumer := pd.Consumer{}
		has, err := rds.Main.QueryStruct(ctx, &newConsumer, sq.QB{
			Where: sq.And(pd.Consumer{}.Column().ID, sq.Equal(mock.NewConsumerID)),
		})
		assert.Equal(t, has, true)
		assert.Equal(t, newConsumer.Name, namespace)
		assert.Equal(t, newConsumer.CreatedAt.Sub(time.Now()) < time.Second * 2, true)
		assert.Equal(t, newConsumer.UpdatedAt.Sub(time.Now()) < time.Second * 2, true)
	}
	// 检查数据
	{
		consumer := pd.Consumer{}
		has, err := rds.Main.QueryStruct(ctx, &consumer, sq.QB{
			Where: sq.And(consumer.Column().ID, sq.Equal(mock.NewConsumerID)),
		})
		assert.NoError(t, err)
		assert.Equal(t, has, true)
		assert.Equal(t, consumer.CreatedAt.Sub(time.Now()).Seconds() < 2, true)
	}
}

func TestDataStorage_ConsumerHasConsumerByID(t *testing.T) {
	namespace := "TestDataStorage_ConsumerHasConsumerByID:" + xtest.String(10)
	ctx := context.TODO()
	rds := connectRDS.TestRDS(t)
	ds := TestDataStorage(t)
	// 清除数据
	_, err := rds.Main.ClearTestData(ctx,sq.QB{
		Table: pd.TableConsumer{},
		Where: sq.And(pd.Consumer{}.Column().Name, sq.LikeLeft(namespace)),
	}) ; if err != nil {panic(err)}
	// 查询 nimoc
	has, reject := ds.ConsumerHasConsumerByName(ctx, namespace)
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