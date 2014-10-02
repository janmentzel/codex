package codex

import (
	"fmt"
)

const (
	MYSQL_QUOTE = '`'
)

type MySqlVisitor struct {
	*ToSqlVisitor
}

var _ VisitorInterface = (*MySqlVisitor)(nil)

func NewMySqlVisitor() *MySqlVisitor {
	return &MySqlVisitor{NewToSqlVisitor()}
}

func (v *MySqlVisitor) Accept(o interface{}) (string, []interface{}, error) {
	err := v.Visit(o, v)

	return v.String(), v.Args(), err
}

// Begin Helpers.

func (v *MySqlVisitor) QuoteTableName(o interface{}, visitor VisitorInterface) (err error) {
	s, ok := o.(string)
	if !ok {
		return fmt.Errorf("MySqlVisitor.QuoteTableName() expected string but got %#v", o)
	}
	visitor.AppendSqlByte(MYSQL_QUOTE)
	visitor.AppendSqlStr(s)
	visitor.AppendSqlByte(MYSQL_QUOTE)
	return
}

func (v *MySqlVisitor) QuoteColumnName(o interface{}, visitor VisitorInterface) (err error) {
	s, ok := o.(string)
	if !ok {
		return fmt.Errorf("MySqlVisitor.QuoteColumnName() expected string but got %#v", o)
	}
	visitor.AppendSqlByte(MYSQL_QUOTE)
	visitor.AppendSqlStr(s)
	visitor.AppendSqlByte(MYSQL_QUOTE)
	return
}

// End Helpers.
