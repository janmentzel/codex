package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateManager(t *testing.T) {
	relation := Relation("table")
	mgr := Modification(relation)

	// The following struct members should exist.
	_ = mgr.Tree

	// The following receiver methods should exist.
	_ = mgr.Set(1)
	_ = mgr.To(2)
	_ = mgr.Where(1)
	_ = mgr.Limit(1)
	_, _, _ = mgr.ToSql()
}

func TestUpdateManagerScope(t *testing.T) {
	mgr := Modification(Relation("users"))
	mgr.Scope("owner_id=?", 77)

	sql, args, err := mgr.Set("id").To(2).ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `UPDATE "users" SET "id"=? WHERE (owner_id=?)`, sql)
	assert.Equal(t, []interface{}{2, 77}, args)
}

func TestUpdateManagerScopeAndWhere(t *testing.T) {
	users := Relation("users")
	mgr := Modification(users)
	mgr.Scope(users.Col("owner_id").Eq(77))
	mgr.Scope(users.Col("active"))
	mgr.Limit(123)

	sql, args, err := mgr.Set("id").To(2).Where("id = ?", 1).ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `UPDATE "users" SET "id"=? WHERE ("users"."owner_id"=?) AND ("users"."active") AND (id = ?) LIMIT ?`, sql)
	assert.Equal(t, []interface{}{2, 77, 1, 123}, args)
}

func TestUpdateManagerScopeWithFunc(t *testing.T) {
	users := Relation("users")
	mgr := Modification(users)

	scope1 := func(s Scoper) {
		s.Scope(users.Col("owner_id").Eq(77))
	}
	scope2 := func(s Scoper) {
		s.Scope(users.Col("active"))
	}
	mgr.Scopes(scope1, scope2)

	sql, args, err := mgr.Set("id").To(2).Where("id = ?", 1).ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `UPDATE "users" SET "id"=? WHERE ("users"."owner_id"=?) AND ("users"."active") AND (id = ?)`, sql)
	assert.Equal(t, []interface{}{2, 77, 1}, args)
}

func TestUpdateManagerSelection(t *testing.T) {
	users := Relation("users")
	mgr := Modification(users)
	scope1 := func(s Scoper) {
		s.Scope(users.Col("owner_id").Eq(77))
	}
	mgr.Scopes(scope1).Limit(1)
	mgr.Where("id > ?", 2).Limit(1)
	sel := mgr.Selection()

	sql, args, err := sel.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "users".* FROM "users" WHERE ("users"."owner_id"=?) AND (id > ?) LIMIT ?`, sql)
	assert.Equal(t, []interface{}{77, 2, 1}, args)
}

func TestUpdateManagerInsertion(t *testing.T) {
	users := Relation("users")
	mgr := Modification(users)
	scope1 := func(s Scoper) {
		s.Scope(users.Col("owner_id").Eq(77))
	}

	mgr.Scopes(scope1).Limit(1)
	mgr.Where("id > ?", 2).Limit(1)

	mod := mgr.Insertion()
	mod.Insert("Undo").Into("name")

	sql, args, err := mod.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `INSERT INTO "users" ("name") VALUES (?)`, sql)
	assert.Equal(t, []interface{}{"Undo"}, args)
}

func TestUpdateManagerDeletion(t *testing.T) {
	users := Relation("users")
	mgr := Modification(users)
	scope1 := func(s Scoper) {
		s.Scope(users.Col("owner_id").Eq(77))
	}
	mgr.Scopes(scope1).Limit(1)
	mgr.Where("id > ?", 2).Limit(1)
	mod := mgr.Deletion()
	sql, args, err := mod.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `DELETE FROM "users" WHERE ("users"."owner_id"=?) AND (id > ?) LIMIT ?`, sql)
	assert.Equal(t, []interface{}{77, 2, 1}, args)
}
