package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectorEmpty(t *testing.T) {
	c := &Collector{}
	assert.Equal(t, "", c.String())
	assert.Equal(t, []interface{}(nil), c.Args())
}

func TestCollector(t *testing.T) {
	c := &Collector{}
	c.AppendSqlStr("WHERE id")
	c.AppendSqlByte(EQUAL)
	c.AppendSqlByte(QUESTION)
	c.AppendArg(77)
	assert.Equal(t, "WHERE id=?", c.String())
	assert.Equal(t, []interface{}{77}, c.Args())
}
