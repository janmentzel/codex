package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertManager(t *testing.T) {
	relation := Table("table")
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
	tlb := Table("users")
	mgr := Insertion(tlb)

	sql, args, err := mgr.Insert("john", "doe", "33").Into("first_name", "last_name", "age").ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `INSERT INTO "users" ("first_name","last_name","age") VALUES (?,?,?)`, sql)
	assert.Equal(t, []interface{}{"john", "doe", "33"}, args)
}

func TestInsertManagerSelection(t *testing.T) {
	users := Table("users")
	mgr := Insertion(users)

	sel := mgr.Selection()

	sql, args, err := sel.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "users".* FROM "users"`, sql)
	assert.Empty(t, args)
}

func TestInsertManagerModification(t *testing.T) {
	users := Table("users")
	mgr := Insertion(users)

	mod := mgr.Modification()
	sql, args, err := mod.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `UPDATE "users" `, sql)
	assert.Empty(t, args)
}

func TestInsertManagerDeletion(t *testing.T) {
	users := Table("users")
	mgr := Insertion(users)

	mod := mgr.Deletion()
	sql, args, err := mod.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `DELETE FROM "users" `, sql)
	assert.Empty(t, args)
}
