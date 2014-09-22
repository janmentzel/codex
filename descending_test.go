package codex

import (
	"testing"
)

func TestDescending(t *testing.T) {
	desc := Descending(1)

	// The following struct members should exist.
	_ = desc.Expr

	// The following receiver methods should exist.
	_ = desc.Or(1)
	_ = desc.And(1)
	_ = desc.Not()
}
