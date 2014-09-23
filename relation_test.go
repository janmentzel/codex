package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelationFactory(t *testing.T) {
	rel := Relation("foo")
	assert.Equal(t, "foo", rel.Name)
	assert.Nil(t, rel.Alias)
}

func TestRelationCol(t *testing.T) {
	rel := Relation("foo")
	col := rel.Col("id")
	assert.IsType(t, &AttributeNode{}, col)
	assert.IsType(t, Column("id"), col.Name)
	assert.IsType(t, rel, col.Relation)
}

func TestRelationSelect(t *testing.T) {
	rel := Relation("foo")
	m := rel.Select("id", Column("name"), Literal("age"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo"."id","name",age FROM "foo"`, sql)
	assert.Empty(t, args)
}

func TestRelationWhereWithNode(t *testing.T) {
	rel := Relation("foo")
	m := rel.Where(rel.Col("id").Eq(1))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" WHERE ("foo"."id"=?)`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestRelationWhereWithSqlAndArgs(t *testing.T) {
	rel := Relation("foo")
	m := rel.Where("a = ? AND b = ?", 22, "bar")

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" WHERE (a = ? AND b = ?)`, sql)
	assert.Equal(t, []interface{}{22, "bar"}, args)
}

func TestRelationInnerJoin(t *testing.T) {
	m := Relation("foo").InnerJoin(Relation("bar"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" INNER JOIN "bar"`, sql)
	assert.Empty(t, args)
}

func TestRelationOuterJoin(t *testing.T) {
	m := Relation("foo").OuterJoin(Relation("bar"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" LEFT OUTER JOIN "bar"`, sql)
	assert.Empty(t, args)
}

func TestRelationOrder(t *testing.T) {
	rel := Relation("foo")
	m := rel.Order(rel.Col("id").Asc())

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" ORDER BY "foo"."id" ASC`, sql)
	assert.Empty(t, args)
}

func TestRelationGroup(t *testing.T) {
	rel := Relation("foo")
	m := rel.Group(rel.Col("id"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" GROUP BY "foo"."id"`, sql)
	assert.Empty(t, args)
}

func TestRelationHaving(t *testing.T) {
	rel := Relation("foo")
	m := rel.Having(rel.Col("id"))

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo".* FROM "foo" HAVING "foo"."id"`, sql)
	assert.Empty(t, args)
}

func TestRelationCount(t *testing.T) {
	rel := Relation("foo")
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

func TestRelationInsert(t *testing.T) {
	rel := Relation("foo")
	m := rel.Insert(1)

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `INSERT INTO "foo" VALUES (?)`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestRelationSet(t *testing.T) {
	rel := Relation("foo")
	m := rel.Set("id").To(1)

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `UPDATE "foo" SET "id"=? `, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestRelationDelete(t *testing.T) {
	rel := Relation("foo")
	m := rel.Delete(1)

	sql, args, err := m.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `DELETE FROM "foo" WHERE ?`, sql)
	assert.Equal(t, []interface{}{1}, args)
}
