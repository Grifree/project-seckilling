package consumerDataStorage

import (
	"context"
	IConsumerDataStorage "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	sq "github.com/goclub/sql"
)

func (dep DataStorage) ConsumerHasConsumerByName(ctx context.Context, name string) (has bool, reject error) {
	consumerCol := pd.Consumer{}.Column()
	has, reject = dep.rds.Main.Has(ctx, sq.QB{
		Table: pd.TableConsumer{},
		Where: sq.And(consumerCol.Name, sq.Equal(name)),
		CheckSQL: []string{"SELECT 1 FROM `consumer` WHERE `name` = ? AND `deleted_at` IS NULL LIMIT ?"},
	}) ; if reject != nil {
		return
	}
	return
}

func (dep DataStorage)  ConsumerCreateConsumer(ctx context.Context, data IConsumerDataStorage.ConsumerCreateConsumer) (consumerID pd.IDConsumer, reject error) {
	consumer := pd.Consumer{
		Name: data.Name,
	}
	reject = dep.rds.Main.InsertModel(ctx, &consumer) ; if reject != nil {
		return
	}
	return consumer.ID, nil
}