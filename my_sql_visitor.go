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

// Begin Unary node visitors.

// func (v *MySqlVisitor) VisitIndexName(o *IndexNameNode, visitor VisitorInterface) string {
// 	return fmt.Sprintf("`%v`", o.Expr)
// }

// // End Unary node visitors.

// // Being Binary node visitors.

// func (v *MySqlVisitor) VisitExistingColumn(o *ExistingColumnNode, visitor VisitorInterface) string {
// 	return fmt.Sprintf("MODIFY COLUMN %v %v", visitor.Visit(o.Left, visitor), visitor.Visit(o.Right, visitor))
// }

// // End Binary node visitors.

// // Begin Nary node visitors.

// func (v *MySqlVisitor) VisitNotNull(o *NotNullNode, visitor VisitorInterface) string {
// 	var typ interface{}

// 	if 0 >= len(o.Options) {
// 		panic("Missing column type definition for MySql NOT NULL constraint.")
// 	} else {
// 		typ = o.Options[0]
// 	}

// 	return fmt.Sprintf("MODIFY %v %v NOT NULL", visitor.Visit(o.Columns, visitor), visitor.Visit(typ, visitor))
// }

// End Nary node visitors.

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
