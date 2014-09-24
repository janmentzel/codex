package codex

import (
	"testing"
)

func TestInsertManager(t *testing.T) {
	relation := Relation("table")
	mgr := Insertion(relation)

	// The following struct members should exist.
	_ = mgr.Tree

	// The following receiver methods should exist.
	_ = mgr.Insert(1)
	_ = mgr.Into(1)
	_ = mgr.Returning(1)
	_, _, _ = mgr.ToSql()
}
