package codex

import (
	"github.com/stretchr/testify/assert"
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

func TestDeleteManagerLimit(t *testing.T) {
	mgr := Deletion(Relation("users"))
	mgr.Limit(1)

	sql, args, err := mgr.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `DELETE FROM "users"  LIMIT ?`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestDeleteManagerScope(t *testing.T) {
	mgr := Deletion(Relation("users"))
	mgr.Scope("owner_id=?", 77)

	sql, args, err := mgr.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `DELETE FROM "users" WHERE (owner_id=?)`, sql)
	assert.Equal(t, []interface{}{77}, args)
}

func TestDeleteManagerScopeAndWhere(t *testing.T) {
	users := Relation("users")
	mgr := Deletion(users)
	mgr.Scope(users.Col("owner_id").Eq(77))
	mgr.Scope(users.Col("active"))

	sql, args, err := mgr.Where("id = ?", 1).ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `DELETE FROM "users" WHERE ("users"."owner_id"=?) AND ("users"."active") AND (id = ?)`, sql)
	assert.Equal(t, []interface{}{77, 1}, args)
}

func TestDeleteManagerScopeWithFunc(t *testing.T) {
	users := Relation("users")
	mgr := Deletion(users)

	scope1 := func(s Scoper) {
		s.Scope(users.Col("owner_id").Eq(77))
	}
	scope2 := func(s Scoper) {
		s.Scope(users.Col("active"))
	}
	mgr.Scopes(scope1, scope2)

	sql, args, err := mgr.Where("id = ?", 1).ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `DELETE FROM "users" WHERE ("users"."owner_id"=?) AND ("users"."active") AND (id = ?)`, sql)
	assert.Equal(t, []interface{}{77, 1}, args)
}

func TestDeleteManagerSelection(t *testing.T) {
	users := Relation("users")
	mgr := Deletion(users)
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

func TestDeleteManagerInsertion(t *testing.T) {
	users := Relation("users")
	mgr := Deletion(users)
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

func TestDeleteManagerModification(t *testing.T) {
	users := Relation("users")
	mgr := Deletion(users)
	scope1 := func(s Scoper) {
		s.Scope(users.Col("owner_id").Eq(77))
	}
	mgr.Scopes(scope1).Limit(1)
	mgr.Where("id > ?", 2).Limit(1)
	mod := mgr.Modification()
	mod.Set("name").To("new Name")
	sql, args, err := mod.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `UPDATE "users" SET "name"=? WHERE ("users"."owner_id"=?) AND (id > ?) LIMIT ?`, sql)
	assert.Equal(t, []interface{}{"new Name", 77, 2, 1}, args)
}
