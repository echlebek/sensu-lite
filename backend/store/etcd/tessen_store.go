package etcd

import (
	"context"

	v2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/store"
)

const (
	tessenPathPrefix = "tessen"
)

var (
	tessenKeyBuilder = store.NewKeyBuilder(tessenPathPrefix)
)

// CreateOrUpdateTessenConfig creates or updates the tessen configuration
func (s *Store) CreateOrUpdateTessenConfig(ctx context.Context, config *v2.TessenConfig) error {
	return CreateOrUpdate(ctx, s.client, tessenKeyBuilder.Build(""), "", config)
}

// GetTessenConfig gets the tessen configuration
func (s *Store) GetTessenConfig(ctx context.Context) (*v2.TessenConfig, error) {
	config := &v2.TessenConfig{}
	err := Get(ctx, s.client, tessenKeyBuilder.Build(""), config)
	return config, err
}
