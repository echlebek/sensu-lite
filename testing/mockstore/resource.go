package mockstore

import (
	"context"

	corev2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/store"
)

// CreateResource ...
func (s *MockStore) CreateResource(ctx context.Context, resource corev2.Resource) error {
	args := s.Called(ctx, resource)
	return args.Error(0)
}

// CreateOrUpdateResource ...
func (s *MockStore) CreateOrUpdateResource(ctx context.Context, resource corev2.Resource) error {
	args := s.Called(ctx, resource)
	return args.Error(0)
}

// DeleteResource ...
func (s *MockStore) DeleteResource(ctx context.Context, kind, name string) error {
	args := s.Called(ctx, kind, name)
	return args.Error(0)
}

// GetResource ...
func (s *MockStore) GetResource(ctx context.Context, name string, resource corev2.Resource) error {
	args := s.Called(ctx, name, resource)
	return args.Error(0)
}

// ListResources ...
func (s *MockStore) ListResources(ctx context.Context, kind string, list interface{}, pred *store.SelectionPredicate) error {
	args := s.Called(ctx, kind, list, pred)
	return args.Error(0)
}
