package codex

import (
	"strconv"
)

type PostgresVisitor struct {
	*ToSqlVisitor
	Bindings int // Number of bindings used in parameterization
}

var _ VisitorInterface = (*PostgresVisitor)(nil)

// creates PostgresVisitor with PostgresCollector
func NewPostgresVisitor() *PostgresVisitor {
	// can not use NewToSqlVisitor() because PostgresCollector needed instead of Collector.
	return &PostgresVisitor{&ToSqlVisitor{NewPostgresCollector()}, 0}
}

func (v *PostgresVisitor) Accept(o interface{}) (string, []interface{}, error) {
	err := v.Visit(o, v)

	return v.String(), v.Args(), err
}

// TODO obsolete
func (v *PostgresVisitor) VisitBinding(o *BindingNode, visitor VisitorInterface) (err error) {
	v.Bindings += 1
	v.AppendSqlByte('$')
	v.AppendSqlStr(strconv.Itoa(v.Bindings))
	return
}

func (v *PostgresVisitor) VisitLike(o *LikeNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlStr(" ILIKE ")
	err = visitor.Visit(o.Right, visitor)
	return
}

func (v *PostgresVisitor) VisitUnlike(o *UnlikeNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlStr(" NOT ILIKE ")
	err = visitor.Visit(o.Right, visitor)
	return
}
