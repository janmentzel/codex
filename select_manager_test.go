package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectManager(t *testing.T) {
	relation := Relation("table")
	mgr := Selection(relation)

	// The following struct members should exist.
	_ = mgr.Tree

	// The following receiver methods should exist.
	_ = mgr.Select(1)
	_ = mgr.Where(1)
	_ = mgr.Offset(1)
	_ = mgr.Limit(1)
	_ = mgr.InnerJoin(1)
	_ = mgr.OuterJoin(1)
	_ = mgr.On(1)
	_ = mgr.Order(1)
	_ = mgr.Group(1)
	_ = mgr.Having(1)
	_ = mgr.Union(Selection(relation))
	_ = mgr.Intersect(Selection(relation))
	_ = mgr.Except(Selection(relation))
	_, _, _ = mgr.ToSql()
}

func TestSelectManagerWhereWithString(t *testing.T) {
	mgr := Selection(Relation("users"))
	sql, args, err := mgr.Where("a").ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "users".* FROM "users" WHERE (a)`, sql)
	assert.Empty(t, args)
}

func TestSelectManagerWhereWithGrouping(t *testing.T) {
	mgr := Selection(Relation("users"))
	sql, args, err := mgr.Where(Grouping(Literal("a"))).ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "users".* FROM "users" WHERE (a)`, sql)
	assert.Empty(t, args)
}

func TestSelectManagerWhereWithEqual(t *testing.T) {
	mgr := Selection(Relation("users"))
	sql, args, err := mgr.Where(Equal(Column("a"), Column("b"))).ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "users".* FROM "users" WHERE ("a"="b")`, sql)
	assert.Empty(t, args)
}

func TestSelectManagerWhereWithSqlAndArg(t *testing.T) {
	mgr := Selection(Relation("users"))
	sql, args, err := mgr.Where("a = ?", 1).ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "users".* FROM "users" WHERE (a = ?)`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestSelectManagerWhereWithSqlAndArgs(t *testing.T) {
	mgr := Selection(Relation("users"))
	sql, args, err := mgr.Where("a = ? AND b = ?", 1, true).ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "users".* FROM "users" WHERE (a = ? AND b = ?)`, sql)
	assert.Equal(t, []interface{}{1, true}, args)
}

func TestSelectManagerScope(t *testing.T) {
	mgr := Selection(Relation("users"))
	mgr.Scope("owner_id=?", 77)

	sql, args, err := mgr.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "users".* FROM "users" WHERE (owner_id=?)`, sql)
	assert.Equal(t, []interface{}{77}, args)
}

func TestSelectManagerScopeAndWhere(t *testing.T) {
	users := Relation("users")
	mgr := Selection(users)
	mgr.Scope(users.Col("owner_id").Eq(77))
	mgr.Scope(users.Col("active"))

	sql, args, err := mgr.Where("id = ?", 1).ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "users".* FROM "users" WHERE ("users"."owner_id"=?) AND ("users"."active") AND (id = ?)`, sql)
	assert.Equal(t, []interface{}{77, 1}, args)
}

func TestSelectManagerScopeWithFunc(t *testing.T) {
	users := Relation("users")
	mgr := Selection(users)

	scope1 := func(s Scoper) {
		s.Scope(users.Col("owner_id").Eq(77))
	}
	scope2 := func(s Scoper) {
		s.Scope(users.Col("active"))
	}
	mgr.Scopes(scope1, scope2)

	sql, args, err := mgr.Where("id = ?", 1).ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "users".* FROM "users" WHERE ("users"."owner_id"=?) AND ("users"."active") AND (id = ?)`, sql)
	assert.Equal(t, []interface{}{77, 1}, args)
}

func TestSelectManagerModification(t *testing.T) {
	users := Relation("users")
	mgr := Selection(users)
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

func TestSelectManagerInsertion(t *testing.T) {
	users := Relation("users")
	mgr := Selection(users)
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

func TestSelectManagerDeletion(t *testing.T) {
	users := Relation("users")
	mgr := Selection(users)
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
