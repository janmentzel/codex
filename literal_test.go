package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLiteralEmpty(t *testing.T) {
	l := Literal("")
	assert.Equal(t, "", l.Sql)
	assert.Empty(t, l.Args)
}

func TestLiteralOneArg(t *testing.T) {
	l := Literal("id = ?", 1)
	assert.Equal(t, "id = ?", l.Sql)
	assert.Equal(t, []interface{}{1}, l.Args)
}

func TestLiteralTwoArgs(t *testing.T) {
	l := Literal("id = ? AND name = ?", 1234, "foobar")
	assert.Equal(t, "id = ? AND name = ?", l.Sql)
	assert.Equal(t, []interface{}{1234, "foobar"}, l.Args)
}

func TestLiteralExpand(t *testing.T) {
	l := Literal("?...")
	assert.Equal(t, "", l.Sql)
	assert.Empty(t, l.Args)

	l = Literal("?...", []interface{}{}...)
	assert.Equal(t, "", l.Sql)
	assert.Empty(t, l.Args)

	l = Literal("?...", 1)
	assert.Equal(t, "?", l.Sql)
	assert.Equal(t, []interface{}{1}, l.Args)

	l = Literal("?...", 1, 2)
	assert.Equal(t, "?,?", l.Sql)
	assert.Equal(t, []interface{}{1, 2}, l.Args)

	l = Literal("?...", 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assert.Equal(t, "?,?,?,?,?,?,?,?,?", l.Sql)
	assert.Equal(t, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}, l.Args)
}

func TestLiteralExpandTwiceRaisesError(t *testing.T) {
	l := Literal("bar IN(?...) OR foo IN(?...)", 1, 2, 3)
	assert.Equal(t, "bar IN(?...) OR foo IN(-- ERROR supports ?... only once per Literal --", l.Sql)
	assert.Equal(t, []interface{}{1, 2, 3}, l.Args)
}
