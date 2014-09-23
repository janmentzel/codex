package codex

import (
	"fmt"
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
	v.AppendSqlStr(fmt.Sprintf("`%v`", o))
	return
}

func (v *MySqlVisitor) QuoteColumnName(o interface{}, visitor VisitorInterface) (err error) {
	v.AppendSqlStr(fmt.Sprintf("`%v`", o))
	return
}

// End Helpers.
