package codex

import (
	"testing"
)

func TestOr(t *testing.T) {
	or := Or(1, 2)

	// The following struct members should exist.
	_ = or.Left
	_ = or.Right

	// The following receiver methods should exist.
	_ = or.Or(1)
	_ = or.And(1)
	_ = or.Not()
}
