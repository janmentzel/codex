package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgresBinding(t *testing.T) {
	sql, args, err := NewPostgresVisitor().Accept(Binding())
	assert.Nil(t, err)
	assert.Equal(t, "$1", sql)
	assert.Equal(t, []interface{}(nil), args)
}

func TestPostgresLike(t *testing.T) {
	sql, args, err := NewPostgresVisitor().Accept(Like(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "$1 ILIKE $2", sql)
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestPostgresUnike(t *testing.T) {
	sql, args, err := NewPostgresVisitor().Accept(Unlike(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "$1 NOT ILIKE $2", sql)
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestPostgresSubSelect(t *testing.T) {
	subTab := Table("sub_table")
	sub := subTab.Select(Sum(Column("s"))).Where(Literal("group_id = ?", 77)).Group(Column("id"))
	one := Table("").Select(Coalesce(sub, 0).As("total"))

	sql, args, err := NewPostgresVisitor().Accept(one)
	assert.Nil(t, err)
	assert.Equal(t, `(SELECT COALESCE((SELECT SUM("s") FROM "sub_table" WHERE (group_id = $1) GROUP BY "id"),$2) AS "total")`, sql)
	assert.Equal(t, []interface{}{77, 0}, args)
}
