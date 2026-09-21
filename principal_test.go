//revive:disable:package-comments
package identity

import (
	"testing"

	"connectrpc.com/connect/v2"
)

func TestInvalidPrincipal(t *testing.T) {
	err := InvalidPrincipal(t.Context())

	if got := connect.CodeOf(err); got != connect.CodeInvalidArgument {
		t.Errorf("code = %v, want %v", got, connect.CodeInvalidArgument)
	}
}
