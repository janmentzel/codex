package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuoteColumnName(t *testing.T) {
	v := NewMySqlVisitor()
	err := v.QuoteColumnName("foo_bar1", v)
	assert.Nil(t, err)
	assert.Equal(t, "`foo_bar1`", v.String())
}

func TestQuoteTableName(t *testing.T) {
	v := NewMySqlVisitor()
	err := v.QuoteTableName("foo_bar2", v)
	assert.Nil(t, err)
	assert.Equal(t, "`foo_bar2`", v.String())
}
