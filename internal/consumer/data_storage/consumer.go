package consumerDS

import (
	"context"
	IConsumerDS "github.com/goclub/project-seckilling/internal/consumer/data_storage/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	sq "github.com/goclub/sql"
	"log"
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

func (dep DS) CreateConsumer(ctx context.Context, data IConsumerDS.CreateConsumer, execUnlock func() error) (consumerID pd.IDConsumer, isRollback bool, reject error) {
	isRollback, reject = dep.rds.Main.BeginTransaction(ctx, func(tx *sq.Transaction) sq.TxResult {
		// 确保后续代码不会使用 dep.rds.Main
		dep := ""
		_=dep
		consumer := pd.Consumer{
			Name: data.Name,
		}
		err := tx.InsertModel(ctx, &consumer) ; if err != nil {
			// 确保解锁函数一定会运行，否则会导致锁必须等到过期时间才释放
			unlockErr := execUnlock()
			// 当事务插入数据失败时，解锁失败只记录，优先返回事务的失败信息
			log.Print(unlockErr)
			return tx.RollbackWithError(err)
		}
		unlockErr := execUnlock() ; if unlockErr != nil {
			return tx.RollbackWithError(unlockErr)
		}
		consumerID = consumer.ID
		return tx.Commit()
	},) ; if reject != nil {
		return
	}
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