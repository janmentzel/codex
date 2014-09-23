package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodexDialectPostgres(t *testing.T) {
	psql := Dialect(POSTGRES)
	users := psql.Table("users")
	q := users.Where("id = ?", 2)

	sql, args, err := q.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "users".* FROM "users" WHERE (id = $1)`, sql)
	assert.Equal(t, []interface{}{2}, args)
}

func TestCodexDialectMysql(t *testing.T) {
	mysql := Dialect(MYSQL)
	users := mysql.Table("users")
	q := users.Where("id = ?", 2)

	sql, args, err := q.ToSql()
	assert.Nil(t, err)
	assert.Equal(t, "SELECT `users`.* FROM `users` WHERE (id = ?)", sql)
	assert.Equal(t, []interface{}{2}, args)
}
