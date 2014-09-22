package codex

import (
	"testing"
)

func TestNot(t *testing.T) {
	not := Not(1)

	// The following struct members should exist.
	_ = not.Expr

	// The following receiver methods should exist.
	_ = not.Or(1)
	_ = not.And(1)
	_ = not.Not()
}
