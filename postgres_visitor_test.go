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
	assert.Equal(t, "? ILIKE ?", sql) // TODO $1 ILIKE $2
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestPostgresUnike(t *testing.T) {
	sql, args, err := NewPostgresVisitor().Accept(Unlike(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "? NOT ILIKE ?", sql) // TODO $1 NOT ILIKE $2
	assert.Equal(t, []interface{}{1, 2}, args)
}
