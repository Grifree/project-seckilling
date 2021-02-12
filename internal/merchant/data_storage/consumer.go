package merchantDS

import (
	"context"
	IMerchantDS "github.com/goclub/project-seckilling/internal/merchant/data_storage/interface"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
	sq "github.com/goclub/sql"
)

func (dep DS) MerchantHasMerchantByName(ctx context.Context, name string) (has bool, reject error) {
	merchantCol := pd.Merchant{}.Column()
	has, reject = dep.rds.Main.Has(ctx, sq.QB{
		Table: pd.TableMerchant{},
		Where: sq.And(merchantCol.Name, sq.Equal(name)),
		CheckSQL: []string{"SELECT 1 FROM `merchant` WHERE `name` = ? AND `deleted_at` IS NULL LIMIT ?"},
	}) ; if reject != nil {
		return
	}
	return
}

func (dep DS) MerchantCreateMerchant(ctx context.Context, data IMerchantDS.MerchantCreateMerchant) (merchantID pd.IDMerchant, reject error) {
	merchant := pd.Merchant{
		Name: data.Name,
	}
	reject = dep.rds.Main.InsertModel(ctx, &merchant) ; if reject != nil {
		return
	}
	return merchant.ID, nil
}

func (dep DS) MerchantHasMerchantByID(ctx context.Context, id pd.IDMerchant) (has bool, reject error) {
	merchantCol := pd.Merchant{}.Column()
	has, reject = dep.rds.Main.Has(ctx, sq.QB{
		Table: pd.TableMerchant{},
		Where: sq.And(merchantCol.ID, sq.Equal(id)),
		CheckSQL: []string{"SELECT 1 FROM `merchant` WHERE `id` = ? AND `deleted_at` IS NULL LIMIT ?"},
	}) ; if reject != nil {
		return
	}
	return
}