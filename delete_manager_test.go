package codex

import (
	"testing"
)

func TestDeleteManager(t *testing.T) {
	relation := Relation("table")
	mgr := Deletion(relation)

	// The following struct members should exist.
	_ = mgr.Tree

	// The following receiver methods should exist.
	_ = mgr.Delete(1)
	_, _, _ = mgr.ToSql()
}
