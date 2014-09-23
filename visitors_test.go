package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVisitorFor(t *testing.T) {
	var v VisitorInterface

	v = VisitorFor(adapter(0))
	assert.IsType(t, &ToSqlVisitor{}, v)

	v = VisitorFor(MYSQL)
	assert.IsType(t, &MySqlVisitor{}, v)

	v = VisitorFor(POSTGRES)
	assert.IsType(t, &PostgresVisitor{}, v)
}
