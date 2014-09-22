package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectorEmpty(t *testing.T) {
	c := &Collector{}
	assert.Equal(t, "", c.SqlBuf.String())
	assert.Equal(t, []interface{}(nil), c.Args)
}

func TestCollector(t *testing.T) {
	c := &Collector{}
	c.Sql("WHERE id")
	c.SqlB(EQUAL)
	c.SqlB(QUESTION)
	c.Arg(77)
	assert.Equal(t, "WHERE id=?", c.SqlBuf.String())
	assert.Equal(t, []interface{}{77}, c.Args)
}
