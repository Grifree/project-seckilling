package IMerchantDS

import (
	"context"
	pd "github.com/goclub/project-seckilling/internal/persistence_data"
)

type Interface interface {
	MerchantHasMerchantByName(ctx context.Context, name string) (has bool, reject error)
	MerchantCreateMerchant(ctx context.Context, data MerchantCreateMerchant) (merchantID pd.IDMerchant, reject error)
	MerchantHasMerchantByID(ctx context.Context, merchantID pd.IDMerchant) (has bool, reject error)
}
type MerchantCreateMerchant struct {
	Name string
}