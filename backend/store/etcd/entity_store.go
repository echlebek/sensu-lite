package etcd

import (
	"context"
	"errors"
	"fmt"

	"github.com/coreos/etcd/clientv3"
	"github.com/gogo/protobuf/proto"
	corev2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/store"
)

const (
	entityPathPrefix = "entities"
)

var (
	entityKeyBuilder = store.NewKeyBuilder(entityPathPrefix)
)

func getEntityPath(entity *corev2.Entity) string {
	return entityKeyBuilder.WithResource(entity).Build(entity.Name)
}

// GetEntitiesPath gets the path of the entity store
func GetEntitiesPath(ctx context.Context, name string) string {
	return entityKeyBuilder.WithContext(ctx).Build(name)
}

// DeleteEntity deletes an Entity.
func (s *Store) DeleteEntity(ctx context.Context, e *corev2.Entity) error {
	if err := e.Validate(); err != nil {
		return err
	}
	_, err := s.client.Delete(ctx, getEntityPath(e))
	return err
}

// DeleteEntityByName deletes an Entity by its name.
func (s *Store) DeleteEntityByName(ctx context.Context, name string) error {
	if name == "" {
		return errors.New("must specify name")
	}

	key := GetEntitiesPath(ctx, name)
	resp, err := s.client.Delete(ctx, key)
	if err != nil {
		return err
	}
	if resp.Deleted == 0 {
		return &store.ErrNotFound{Key: key}
	} else if resp.Deleted > 1 {
		return &store.ErrInternal{
			Message: fmt.Sprintf("expected to delete exactly 1 key, deleted %d", resp.Deleted),
		}
	}

	return nil
}

// GetEntityByName gets an Entity by its name.
func (s *Store) GetEntityByName(ctx context.Context, name string) (*corev2.Entity, error) {
	if name == "" {
		return nil, errors.New("must specify name")
	}

	resp, err := s.client.Get(ctx, GetEntitiesPath(ctx, name), clientv3.WithLimit(1))
	if err != nil {
		return nil, err
	}
	if len(resp.Kvs) != 1 {
		return nil, nil
	}
	entity := &corev2.Entity{}
	if err := unmarshal(resp.Kvs[0].Value, entity); err != nil {
		return nil, err
	}

	if entity.Labels == nil {
		entity.Labels = make(map[string]string)
	}
	if entity.Annotations == nil {
		entity.Annotations = make(map[string]string)
	}
	return entity, nil
}

// GetEntities returns the entities for the namespace in the supplied context.
func (s *Store) GetEntities(ctx context.Context, pred *store.SelectionPredicate) ([]*corev2.Entity, error) {
	entities := []*corev2.Entity{}
	err := List(ctx, s.client, GetEntitiesPath, &entities, pred)
	return entities, err
}

// UpdateEntity updates an Entity.
func (s *Store) UpdateEntity(ctx context.Context, e *corev2.Entity) error {
	if err := e.Validate(); err != nil {
		return err
	}

	eStr, err := proto.Marshal(e)
	if err != nil {
		return err
	}

	cmp := clientv3.Compare(clientv3.Version(getNamespacePath(e.Namespace)), ">", 0)
	req := clientv3.OpPut(getEntityPath(e), string(eStr))
	res, err := s.client.Txn(ctx).If(cmp).Then(req).Commit()
	if err != nil {
		return err
	}
	if !res.Succeeded {
		return fmt.Errorf(
			"could not create the entity %s in namespace %s",
			e.Name,
			e.Namespace,
		)
	}

	return nil
}
