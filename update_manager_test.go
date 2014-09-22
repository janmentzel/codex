package codex

import (
	"testing"
)

func TestUpdateManager(t *testing.T) {
	relation := Relation("table")
	mgr := Modification(relation)

	// The following struct members should exist.
	_ = mgr.Tree

	// The following receiver methods should exist.
	_ = mgr.Set(1)
	_ = mgr.To(1)
	_ = mgr.Where(1)
	_ = mgr.Limit(1)
	_ = mgr.SetAdapter(1)
	_, _, _ = mgr.ToSql()
}
