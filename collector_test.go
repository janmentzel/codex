package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectorEmpty(t *testing.T) {
	c := &Collector{}
	assert.Equal(t, 0, cap(c.sqlBuf.Bytes()))
	assert.Equal(t, "", c.String())
	assert.Empty(t, c.Args())
}

func TestCollector(t *testing.T) {
	c := NewCollector()
	assert.Equal(t, EXPECTED_SQL_QUERY_LEN, cap(c.sqlBuf.Bytes()))

	c.AppendSqlStr("WHERE id")
	c.AppendSqlByte(EQUAL)
	c.AppendSqlByte(QUESTION)
	c.AppendArg(77)
	assert.Equal(t, "WHERE id=?", c.String())
	assert.Equal(t, []interface{}{77}, c.Args())
}

func TestPostgresCollectorEmpty(t *testing.T) {
	c := &PostgresCollector{}
	assert.Equal(t, 0, cap(c.sqlBuf.Bytes()))
	assert.Equal(t, "", c.String())
	assert.Empty(t, c.Args())
}

func TestPostgresCollectorAppendByte(t *testing.T) {
	c := NewPostgresCollector()
	assert.Equal(t, EXPECTED_SQL_QUERY_LEN, cap(c.sqlBuf.Bytes()))

	c.AppendSqlByte(LT)
	assert.Equal(t, "<", c.String())
	c.AppendSqlByte(EQUAL)
	assert.Equal(t, "<=", c.String())
	c.AppendSqlByte(SPACE)
	assert.Equal(t, "<= ", c.String())
	c.AppendSqlByte(QUESTION)
	assert.Equal(t, "<= $1", c.String())
	c.AppendSqlByte(COMMA)
	c.AppendSqlByte(QUESTION)
	assert.Equal(t, "<= $1,$2", c.String())
	c.AppendSqlByte(COMMA)
	c.AppendSqlByte(QUESTION)
	assert.Equal(t, "<= $1,$2,$3", c.String())
	c.AppendSqlByte(COMMA)
	c.AppendSqlByte(QUESTION)
	assert.Equal(t, "<= $1,$2,$3,$4", c.String())
	c.AppendSqlByte(COMMA)
	c.AppendSqlByte(QUESTION)
	assert.Equal(t, "<= $1,$2,$3,$4,$5", c.String())
	c.AppendSqlByte(COMMA)
	c.AppendSqlByte(QUESTION)
	assert.Equal(t, "<= $1,$2,$3,$4,$5,$6", c.String())
	c.AppendSqlByte(COMMA)
	c.AppendSqlByte(QUESTION)
	assert.Equal(t, "<= $1,$2,$3,$4,$5,$6,$7", c.String())
	c.AppendSqlByte(COMMA)
	c.AppendSqlByte(QUESTION)
	assert.Equal(t, "<= $1,$2,$3,$4,$5,$6,$7,$8", c.String())
	c.AppendSqlByte(COMMA)
	c.AppendSqlByte(QUESTION)
	assert.Equal(t, "<= $1,$2,$3,$4,$5,$6,$7,$8,$9", c.String())
	c.AppendSqlByte(COMMA)
	c.AppendSqlByte(QUESTION)
	assert.Equal(t, "<= $1,$2,$3,$4,$5,$6,$7,$8,$9,$10", c.String())
	c.AppendSqlByte(COMMA)
	c.AppendSqlByte(QUESTION)
	assert.Equal(t, "<= $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11", c.String())
}

func TestPostgresCollector(t *testing.T) {
	c := NewPostgresCollector()
	c.AppendSqlStr("WHERE id")
	c.AppendSqlByte(EQUAL)
	c.AppendSqlByte(QUESTION)
	c.AppendArg(77)
	assert.Equal(t, "WHERE id=$1", c.String())
	assert.Equal(t, []interface{}{77}, c.Args())
}

func TestPostgresCollectorAppendSqlStr(t *testing.T) {
	c := NewPostgresCollector()
	c.AppendSqlStr("WHERE id=? AND name=?")
	c.AppendArg(77)
	c.AppendArg("Hans")
	assert.Equal(t, "WHERE id=$1 AND name=$2", c.String())
	assert.Equal(t, []interface{}{77, "Hans"}, c.Args())
}

func TestPostgresCollectorAppendSqlStrWithoutPlaceholder(t *testing.T) {
	c := NewPostgresCollector()
	c.AppendSqlStr("WHERE id=foo_id AND bar=bla")
	assert.Equal(t, "WHERE id=foo_id AND bar=bla", c.String())
	assert.Empty(t, c.Args())
}
