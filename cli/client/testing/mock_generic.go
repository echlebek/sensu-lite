package testing

import (
	"net/http"

	"github.com/echlebek/sensu-lite/cli/client"
	"github.com/echlebek/sensu-lite/types"
)

// Delete ...
func (c *MockClient) Delete(path string) error {
	args := c.Called(path)
	return args.Error(0)
}

// Get ...
func (c *MockClient) Get(path string, obj interface{}) error {
	args := c.Called(path, obj)
	return args.Error(0)
}

// List ...
func (c *MockClient) List(path string, objs interface{}, options *client.ListOptions, header *http.Header) error {
	args := c.Called(path, objs, options, header)
	return args.Error(0)
}

// Post ...
func (c *MockClient) Post(path string, obj interface{}) error {
	args := c.Called(path, obj)
	return args.Error(0)
}

// Put ...
func (c *MockClient) Put(path string, obj interface{}) error {
	args := c.Called(path, obj)
	return args.Error(0)
}

// PutResource ...
func (c *MockClient) PutResource(r types.Wrapper) error {
	args := c.Called(r)
	return args.Error(0)
}
