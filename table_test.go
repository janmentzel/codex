package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTableFactory(t *testing.T) {
	rel := Table("foo")
	assert.Equal(t, "foo", rel.Name)
	assert.Nil(t, rel.Alias)
	assert.Empty(t, rel.scopes)
}

func TestTableScopes(t *testing.T) {
	rel := Table("foo")

	scope1 := func(s Scoper) {
		s.Scope(rel.Col("owner_id").Eq(77))
	}
	scope2 := func(s Scoper) {
		s.Scope(rel.Col("active"))
	}
	rel.Scopes(scope1, scope2)
	assert.Equal(t, []ScopeFunc{scope1, scope2}, rel.scopes)
}

func TestTableCol(t *testing.T) {
	rel := Table("foo")
	col := rel.Col("id")
	assert.IsType(t, &AttributeNode{}, col)
	assert.IsType(t, Column("id"), col.Name)
	assert.IsType(t, rel, col.Table)
}

func TestTableSelect(t *testing.T) {
	rel := Table("foo")
	m := rel.Select("id", Column("name"), Literal("age"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo"."id","name",age FROM "foo"`, sql)
	assert.Empty(t, args)
}

func TestTableSelectScoped(t *testing.T) {
	rel := Table("foo")
	scope := func(s Scoper) {
		s.Scope(rel.Col("owner_id").Eq(77))
	}
	m := rel.Scopes(scope).Select("id", Column("name"), Literal("age"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo"."id","name",age FROM "foo" WHERE ("foo"."owner_id"=?)`, sql)
	assert.Equal(t, []interface{}{77}, args)
}

func TestTableSelectStar(t *testing.T) {
	rel := Table("foo")

	m := rel.Select(Star())

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT * FROM "foo"`, sql)
	assert.Empty(t, args)
}

func TestTableSelectTableStar(t *testing.T) {
	rel := Table("foo")

	m := rel.Select(rel.Star())

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo"`, sql)
	assert.Empty(t, args)
}

func TestTableWhereWithNode(t *testing.T) {
	rel := Table("foo")
	m := rel.Where(rel.Col("id").Eq(1))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" WHERE ("foo"."id"=?)`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestTableWhereWithNodeAndScope(t *testing.T) {
	rel := Table("foo")
	scope := func(s Scoper) {
		s.Scope(rel.Col("owner_id").Eq(77))
	}
	m := rel.Scopes(scope).Where(rel.Col("id").Eq(1))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" WHERE ("foo"."owner_id"=?) AND ("foo"."id"=?)`, sql)
	assert.Equal(t, []interface{}{77, 1}, args)
}

func TestTableWhereWithSqlAndArgs(t *testing.T) {
	rel := Table("foo")
	m := rel.Where("a = ? AND b = ?", 22, "bar")

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" WHERE (a = ? AND b = ?)`, sql)
	assert.Equal(t, []interface{}{22, "bar"}, args)
}

func TestTableInnerJoin(t *testing.T) {
	m := Table("foo").InnerJoin(Table("bar"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" INNER JOIN "bar"`, sql)
	assert.Empty(t, args)
}

func TestTableInnerJoinScoped(t *testing.T) {
	rel := Table("foo")
	scope := func(s Scoper) {
		s.Scope(rel.Col("owner_id").Eq(77))
	}

	m := rel.Scopes(scope).InnerJoin(Table("bar"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" INNER JOIN "bar" WHERE ("foo"."owner_id"=?)`, sql)
	assert.Equal(t, []interface{}{77}, args)
}

func TestTableOuterJoin(t *testing.T) {
	m := Table("foo").OuterJoin(Table("bar"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" LEFT OUTER JOIN "bar"`, sql)
	assert.Empty(t, args)
}

func TestTableOuterJoinScoped(t *testing.T) {
	rel := Table("foo")
	scope := func(s Scoper) {
		s.Scope(rel.Col("owner_id").Eq(77))
	}

	m := rel.Scopes(scope).OuterJoin(Table("bar"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" LEFT OUTER JOIN "bar" WHERE ("foo"."owner_id"=?)`, sql)
	assert.Equal(t, []interface{}{77}, args)
}

func TestTableOrder(t *testing.T) {
	rel := Table("foo")
	m := rel.Order(rel.Col("id").Asc())

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" ORDER BY "foo"."id" ASC`, sql)
	assert.Empty(t, args)
}

func TestTableOrderScoped(t *testing.T) {
	rel := Table("foo")
	scope := func(s Scoper) {
		s.Scope(rel.Col("owner_id").Eq(77))
	}

	m := rel.Scopes(scope).Order(rel.Col("id").Asc())

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" WHERE ("foo"."owner_id"=?) ORDER BY "foo"."id" ASC`, sql)
	assert.Equal(t, []interface{}{77}, args)
}

func TestTableOrderTwo(t *testing.T) {
	rel := Table("foo")
	m := rel.Order(rel.Col("group_id").Desc()).Order(rel.Col("name").Asc())

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" ORDER BY "foo"."group_id" DESC,"foo"."name" ASC`, sql)
	assert.Empty(t, args)
}

func TestTableGroup(t *testing.T) {
	rel := Table("foo")
	m := rel.Group(rel.Col("id"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" GROUP BY "foo"."id"`, sql)
	assert.Empty(t, args)
}

func TestTableGroupScoped(t *testing.T) {
	rel := Table("foo")
	scope := func(s Scoper) {
		s.Scope(rel.Col("owner_id").Eq(77))
	}

	m := rel.Scopes(scope).Group(rel.Col("id"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" WHERE ("foo"."owner_id"=?) GROUP BY "foo"."id"`, sql)
	assert.Equal(t, []interface{}{77}, args)
}

func TestTableGroupTwoCols(t *testing.T) {
	rel := Table("foo")
	m := rel.Group(rel.Col("id"), rel.Col("bar_id"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" GROUP BY "foo"."id","foo"."bar_id"`, sql)
	assert.Empty(t, args)
}

func TestTableHaving(t *testing.T) {
	rel := Table("foo")
	m := rel.Having(rel.Col("id"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" HAVING "foo"."id"`, sql)
	assert.Empty(t, args)
}

func TestTableHavingScoped(t *testing.T) {
	rel := Table("foo")
	scope := func(s Scoper) {
		s.Scope(rel.Col("owner_id").Eq(77))
	}

	m := rel.Scopes(scope).Having(rel.Col("id"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" WHERE ("foo"."owner_id"=?) HAVING "foo"."id"`, sql)
	assert.Equal(t, []interface{}{77}, args)
}

func TestTableCount(t *testing.T) {
	rel := Table("foo")
	m := rel.Select("id").Where(rel.Col("id").Eq(1))
	m1 := m.Count("id")

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo"."id" FROM "foo" WHERE ("foo"."id"=?)`, sql)
	assert.Equal(t, []interface{}{1}, args)

	sql, args, err = m1.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT COUNT("id") FROM "foo" WHERE ("foo"."id"=?)`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestTableCountScoped(t *testing.T) {
	rel := Table("foo")
	scope := func(s Scoper) {
		s.Scope(rel.Col("owner_id").Eq(77))
	}
	rel.Scopes(scope)

	m := rel.Select("id").Where(rel.Col("id").Eq(1))
	m1 := m.Count("id")

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo"."id" FROM "foo" WHERE ("foo"."owner_id"=?) AND ("foo"."id"=?)`, sql)
	assert.Equal(t, []interface{}{77, 1}, args)

	sql, args, err = m1.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT COUNT("id") FROM "foo" WHERE ("foo"."owner_id"=?) AND ("foo"."id"=?)`, sql)
	assert.Equal(t, []interface{}{77, 1}, args)
}

func TestTableInsert(t *testing.T) {
	rel := Table("foo")
	m := rel.Insert(1)

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `INSERT INTO "foo" VALUES (?)`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestTableSet(t *testing.T) {
	rel := Table("foo")
	m := rel.Set("id").To(1)

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `UPDATE "foo" SET "id"=? `, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestTableSetScoped(t *testing.T) {
	rel := Table("foo")
	scope := func(s Scoper) {
		s.Scope(rel.Col("owner_id").Eq(77))
	}
	m := rel.Scopes(scope).Set("id").To(1)

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `UPDATE "foo" SET "id"=? WHERE ("foo"."owner_id"=?)`, sql)
	assert.Equal(t, []interface{}{1, 77}, args)
}

func TestTableDelete(t *testing.T) {
	rel := Table("foo")
	m := rel.Delete(1)

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `DELETE FROM "foo" WHERE (?)`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestTableDeleteScoped(t *testing.T) {
	rel := Table("foo")
	scope := func(s Scoper) {
		s.Scope(rel.Col("owner_id").Eq(77))
	}
	m := rel.Scopes(scope).Delete(1)

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `DELETE FROM "foo" WHERE ("foo"."owner_id"=?) AND (?)`, sql)
	assert.Equal(t, []interface{}{77, 1}, args)
}

func TestTableInsertion(t *testing.T) {
	tab := Table("foo")
	assert.IsType(t, &InsertManager{}, tab.Insertion())
	assert.Equal(t, tab, tab.Insertion().Table())
}

func TestTableModification(t *testing.T) {
	tab := Table("foo")
	assert.IsType(t, &UpdateManager{}, tab.Modification())
	assert.Equal(t, tab, tab.Modification().Table())
}

func TestTableSelection(t *testing.T) {
	tab := Table("foo")
	assert.IsType(t, &SelectManager{}, tab.Selection())
	assert.Equal(t, tab, tab.Selection().Table())
}

func TestTableDeletion(t *testing.T) {
	tab := Table("foo")
	assert.IsType(t, &DeleteManager{}, tab.Deletion())
	assert.Equal(t, tab, tab.Deletion().Table())
}

func TestTableColLiteral(t *testing.T) {

	psql := Dialect(POSTGRES)
	products := psql.Table("products")
	sql, args, err := products.Where(products.Col("tags").Literal("@> ARRAY[?...]", "fancy", "cheap", "retro")).ToSql()

	assert.Nil(t, err)
	assert.Equal(t, `SELECT "products".* FROM "products" WHERE ("products"."tags" @> ARRAY[$1,$2,$3])`, sql)
	assert.Equal(t, []interface{}{"fancy", "cheap", "retro"}, args)
}
