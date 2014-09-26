package codex

import (
	"github.com/stretchr/testify/assert"
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

func TestInsertManagerInsert(t *testing.T) {
	tlb := Relation("users")
	mgr := Insertion(tlb)

	sql, args, err := mgr.Insert("john", "doe", "33").Into("first_name", "last_name", "age").ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `INSERT INTO "users" ("first_name","last_name","age") VALUES (?,?,?)`, sql)
	assert.Equal(t, []interface{}{"john", "doe", "33"}, args)
}
