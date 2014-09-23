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
