package api

import (
	"context"
	"errors"

	"github.com/echlebek/sensu-lite/backend/authentication/jwt"
	"github.com/echlebek/sensu-lite/backend/authorization"
	"github.com/echlebek/sensu-lite/types"
)

// addAuthUser finds the user in the provided context, and sets it on the
// attrs.
func addAuthUser(ctx context.Context, attrs *authorization.Attributes) error {
	// Get the claims from the request context
	claims := jwt.GetClaimsFromContext(ctx)
	if claims == nil {
		return errors.New("no claims found in the request context")
	}

	// Add the user to our request info
	attrs.User = types.User{
		Username: claims.Subject,
		Groups:   claims.Groups,
	}

	return nil
}

// authorize is a convenience function for clients that want to test if the
// operation specified by the attributes is authorized.
func authorize(ctx context.Context, auth authorization.Authorizer, attrs *authorization.Attributes) error {
	if err := addAuthUser(ctx, attrs); err != nil {
		return err
	}
	authorized, err := auth.Authorize(ctx, attrs)
	if err != nil {
		return err
	}
	if !authorized {
		return authorization.ErrUnauthorized
	}
	return nil
}
