package asset

import (
	"context"
	"os"
	"path/filepath"
	"sync"

	"github.com/echlebek/sensu-lite/types"
	bolt "go.etcd.io/bbolt"
)

const (
	dbName = "assets.db"
)

// Manager ...
type Manager struct {
	cacheDir string
	entity   *types.Entity
	stopping chan struct{}
	wg       *sync.WaitGroup
}

// NewManager ...
func NewManager(cacheDir string, entity *types.Entity, wg *sync.WaitGroup) *Manager {
	return &Manager{
		cacheDir: cacheDir,
		entity:   entity,
		wg:       wg,
	}
}

// StartAssetManager starts the asset manager for a backend or agent.
func (m *Manager) StartAssetManager(ctx context.Context) (Getter, error) {
	// create agent cache directory if it doesn't already exist
	if err := os.MkdirAll(m.cacheDir, 0755); err != nil {
		return nil, err
	}

	logger.WithField("cache", m.cacheDir).Debug("initializing cache directory")
	db, err := bolt.Open(filepath.Join(m.cacheDir, dbName), 0600, &bolt.Options{})
	if err != nil {
		return nil, err
	}
	logger.WithField("cache", m.cacheDir).Debug("done initializing cache directory")

	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		<-ctx.Done()
		if err := db.Close(); err != nil {
			logger.Debug(err)
		}
	}()
	boltDBGetter := NewBoltDBGetter(
		db, m.cacheDir, nil, nil, nil)

	return NewFilteredManager(boltDBGetter, m.entity), nil
}
