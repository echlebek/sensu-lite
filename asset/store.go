package asset

import (
	"context"

	"github.com/echlebek/sensu-lite/backend/store"
	"github.com/echlebek/sensu-lite/types"
)

// GetAssets retrieves all Assets from the store if contained in the list of asset names
func GetAssets(ctx context.Context, store store.Store, assetList []string) []types.Asset {
	assets := []types.Asset{}

	for _, assetName := range assetList {
		asset, err := store.GetAssetByName(ctx, assetName)
		if err != nil {
			logger.WithField("asset", assetName).WithError(err).Error("error fetching asset from store")
		} else if asset == nil {
			logger.WithField("asset", assetName).Info("asset does not exist")
		} else {
			assets = append(assets, *asset)
		}
	}

	return assets
}
