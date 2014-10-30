package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunction(t *testing.T) {
	f := Function("TEST_FUNC", 1, 2, 3, 4)

	// The following struct members should exist.
	assert.Equal(t, "TEST_FUNC", f.Name)
	assert.Equal(t, []interface{}{1, 2, 3, 4}, f.Args)
	assert.Nil(t, f.Alias)
	assert.False(t, f.Distinct)

	// The following receiver methods should exist.
	_ = f.And(1)
	_ = f.Or(1)
	_ = f.Eq(1)
	_ = f.Neq(1)
	_ = f.Gt(1)
	_ = f.Gte(1)
	_ = f.Lt(1)
	_ = f.Lte(1)
	_ = f.Like(1)
	_ = f.Unlike(1)
}

func TestFunctions(t *testing.T) {

	f := Avg(1, 2, 3)
	assert.Equal(t, "AVG", f.Name)
	assert.Equal(t, []interface{}{1, 2, 3}, f.Args)

	f = Coalesce(1, 2, 3)
	assert.Equal(t, "COALESCE", f.Name)
	assert.Equal(t, []interface{}{1, 2, 3}, f.Args)

	f = Count(1, 2, 3)
	assert.Equal(t, "COUNT", f.Name)
	assert.Equal(t, []interface{}{1, 2, 3}, f.Args)

	f = Max(1, 2, 3)
	assert.Equal(t, "MAX", f.Name)
	assert.Equal(t, []interface{}{1, 2, 3}, f.Args)

	f = Min(1, 2, 3)
	assert.Equal(t, "MIN", f.Name)
	assert.Equal(t, []interface{}{1, 2, 3}, f.Args)

	f = Sum(1, 2, 3)
	assert.Equal(t, "SUM", f.Name)
	assert.Equal(t, []interface{}{1, 2, 3}, f.Args)

	f = Lower(1, 2, 3)
	assert.Equal(t, "LOWER", f.Name)
	assert.Equal(t, []interface{}{1, 2, 3}, f.Args)

	f = Upper(1, 2, 3)
	assert.Equal(t, "UPPER", f.Name)
	assert.Equal(t, []interface{}{1, 2, 3}, f.Args)

	f = Substring(1, 2, 3)
	assert.Equal(t, "SUBSTRING", f.Name)
	assert.Equal(t, []interface{}{1, 2, 3}, f.Args)
}
