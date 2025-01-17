package graphql

import (
	"github.com/echlebek/sensu-lite/backend/apid/graphql/schema"
	"github.com/echlebek/sensu-lite/backend/authentication/jwt"
	"github.com/echlebek/sensu-lite/graphql"
)

var _ schema.ViewerFieldResolvers = (*viewerImpl)(nil)

//
// Implement QueryFieldResolvers
//

type viewerImpl struct {
	userClient UserClient
}

// Namespaces implements response to request for 'namespaces' field.
func (r *viewerImpl) Namespaces(p graphql.ResolveParams) (interface{}, error) {
	return loadNamespaces(p.Context)
}

// User implements response to request for 'user' field.
func (r *viewerImpl) User(p graphql.ResolveParams) (interface{}, error) {
	claims := jwt.GetClaimsFromContext(p.Context)
	if claims == nil {
		return nil, nil
	}

	res, err := r.userClient.FetchUser(p.Context, claims.Subject)
	return handleFetchResult(res, err)
}
