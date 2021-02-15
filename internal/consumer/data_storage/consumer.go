package consumerDS

import (
	"context"
	IConsumerDS "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	sq "github.com/goclub/sql"
)

func (dep DS) HasConsumerByName(ctx context.Context, name string) (has bool, reject error) {
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

func (dep DS) CreateConsumer(ctx context.Context, data IConsumerDS.CreateConsumer) (consumerID pd.IDConsumer, reject error) {
	consumer := pd.Consumer{
		Name: data.Name,
	}
	reject = dep.rds.Main.InsertModel(ctx, &consumer);if reject != nil {
		return
	}
	consumerID = consumer.ID
	return
}

func (dep DS) ConsumerHasConsumerByID(ctx context.Context, id pd.IDConsumer) (has bool, reject error) {
	consumerCol := pd.Consumer{}.Column()
	has, reject = dep.rds.Main.Has(ctx, sq.QB{
		Table: pd.TableConsumer{},
		Where: sq.And(consumerCol.ID, sq.Equal(id)),
		CheckSQL: []string{"SELECT 1 FROM `consumer` WHERE `id` = ? AND `deleted_at` IS NULL LIMIT ?"},
	}) ; if reject != nil {
		return
	}
	return
}