package handlers

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	corev2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/apid/actions"
	"github.com/echlebek/sensu-lite/backend/store"
)

// CreateOrUpdateResource creates or updates the resource given in the request
// body, regardless of whether it already exists or not
func (h Handlers) CreateOrUpdateResource(r *http.Request) (interface{}, error) {
	payload := reflect.New(reflect.TypeOf(h.Resource).Elem())
	if err := json.NewDecoder(r.Body).Decode(payload.Interface()); err != nil {
		return nil, actions.NewError(actions.InvalidArgument, err)
	}

	if err := CheckMeta(payload.Interface(), mux.Vars(r)); err != nil {
		return nil, actions.NewError(actions.InvalidArgument, err)
	}

	resource, ok := payload.Interface().(corev2.Resource)
	if !ok {
		return nil, actions.NewErrorf(actions.InvalidArgument)
	}

	if err := h.Store.CreateOrUpdateResource(r.Context(), resource); err != nil {
		switch err := err.(type) {
		case *store.ErrNotValid:
			return nil, actions.NewErrorf(actions.InvalidArgument)
		default:
			return nil, actions.NewError(actions.InternalErr, err)
		}
	}

	return nil, nil
}
