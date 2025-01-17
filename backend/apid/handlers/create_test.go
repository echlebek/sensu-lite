package handlers

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	corev2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/store"
	"github.com/echlebek/sensu-lite/testing/fixture"
	"github.com/echlebek/sensu-lite/testing/mockstore"
	"github.com/stretchr/testify/mock"
)

func TestHandlers_CreateResource(t *testing.T) {
	type storeFunc func(*mockstore.MockStore)
	tests := []struct {
		name      string
		body      []byte
		urlVars   map[string]string
		storeFunc storeFunc
		wantErr   bool
	}{
		{
			name:    "invalid request body",
			body:    []byte("foobar"),
			wantErr: true,
		},
		{
			name: "invalid resource meta",
			body: marshal(t, fixture.Resource{ObjectMeta: corev2.ObjectMeta{
				Name:      "foo",
				Namespace: "acme",
			}}),
			urlVars: map[string]string{"id": "bar", "namespace": "acme"},
			wantErr: true,
		},
		{
			name: "store err, already exists",
			body: marshal(t, fixture.Resource{ObjectMeta: corev2.ObjectMeta{}}),
			storeFunc: func(s *mockstore.MockStore) {
				s.On("CreateResource", mock.Anything, mock.AnythingOfType("*fixture.Resource")).
					Return(&store.ErrAlreadyExists{})
			},
			wantErr: true,
		},
		{
			name: "store err, not valid",
			body: marshal(t, fixture.Resource{ObjectMeta: corev2.ObjectMeta{}}),
			storeFunc: func(s *mockstore.MockStore) {
				s.On("CreateResource", mock.Anything, mock.AnythingOfType("*fixture.Resource")).
					Return(&store.ErrNotValid{})
			},
			wantErr: true,
		},
		{
			name: "store err, default",
			body: marshal(t, fixture.Resource{ObjectMeta: corev2.ObjectMeta{}}),
			storeFunc: func(s *mockstore.MockStore) {
				s.On("CreateResource", mock.Anything, mock.AnythingOfType("*fixture.Resource")).
					Return(&store.ErrInternal{})
			},
			wantErr: true,
		},
		{
			name: "successful create",
			body: marshal(t, fixture.Resource{ObjectMeta: corev2.ObjectMeta{}}),
			storeFunc: func(s *mockstore.MockStore) {
				s.On("CreateResource", mock.Anything, mock.AnythingOfType("*fixture.Resource")).
					Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &mockstore.MockStore{}
			if tt.storeFunc != nil {
				tt.storeFunc(store)
			}

			h := Handlers{
				Resource: &fixture.Resource{},
				Store:    store,
			}

			r, _ := http.NewRequest(http.MethodPost, "/", bytes.NewReader(tt.body))
			r = mux.SetURLVars(r, tt.urlVars)

			_, err := h.CreateResource(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handlers.CreateResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
