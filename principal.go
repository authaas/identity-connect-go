//revive:disable:package-comments
package identity

import (
	"context"

	errors "github.com/pbrpc/connect-errors"
)

// InvalidPrincipal refuses a request whose principal id is not one the
// system can hold.
func InvalidPrincipal(ctx context.Context) error {
	return errors.InvalidArgument(ctx, "validation failed", errors.FieldViolation{
		Field:       "principal.id",
		Description: "is not a UUID",
	})
}
