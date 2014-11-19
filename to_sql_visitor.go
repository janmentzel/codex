package codex

import (
	"fmt"
)

var DEBUG = false

const (
	SPACE    = ' '
	COMMA    = ','
	STAR     = '*'
	QUESTION = '?'
	DOT      = '.'
	EQUAL    = '='
	GT       = '>'
	LT       = '<'
	GTEQUAL  = ">="
	LTEQUAL  = "<="
	QUOTE    = '"'

	// Keywords
	SELECT   = `SELECT `
	FROM     = ` FROM `
	WHERE    = ` WHERE `
	ORDER_BY = ` ORDER BY `
	GROUP_BY = ` GROUP BY `
	AND      = ` AND `
	OR       = ` OR `
	AS       = ` AS `
)

type ToSqlVisitor struct {
	CollectorInterface
}

var _ VisitorInterface = (*ToSqlVisitor)(nil)

// creates ToSqlVisitor with standard Collector
func NewToSqlVisitor(collectors ...CollectorInterface) *ToSqlVisitor {
	var collector CollectorInterface
	if len(collectors) == 0 {
		collector = NewCollector()
	} else {
		collector = collectors[0]
	}
	return &ToSqlVisitor{collector}
}

func (v *ToSqlVisitor) Accept(o interface{}) (string, []interface{}, error) {
	err := v.Visit(o, v)

	return v.String(), v.Args(), err
}

func (v *ToSqlVisitor) Visit(o interface{}, visitor VisitorInterface) error {

	if DEBUG {
		fmt.Printf("DEBUG: Visiting %T\n", o)
	}

	switch o.(type) {
	// Unary node visitors.
	case *GroupingNode:
		return visitor.VisitGrouping(o.(*GroupingNode), visitor)
	case *NotNode:
		return visitor.VisitNot(o.(*NotNode), visitor)
	case *LiteralNode:
		return visitor.VisitLiteral(o.(*LiteralNode), visitor)
	case *OnNode:
		return visitor.VisitOn(o.(*OnNode), visitor)
	case *ColumnNode:
		return visitor.VisitColumn(o.(*ColumnNode), visitor)
	case *StarNode:
		return visitor.VisitStar(o.(*StarNode), visitor)
	case *BindingNode:
		return visitor.VisitBinding(o.(*BindingNode), visitor)
	case *LimitNode:
		return visitor.VisitLimit(o.(*LimitNode), visitor)
	case *OffsetNode:
		return visitor.VisitOffset(o.(*OffsetNode), visitor)
	case *HavingNode:
		return visitor.VisitHaving(o.(*HavingNode), visitor)
	case *AscendingNode:
		return visitor.VisitAscending(o.(*AscendingNode), visitor)
	case *DescendingNode:
		return visitor.VisitDescending(o.(*DescendingNode), visitor)

	// Binary node visitors.
	case *AsNode:
		return visitor.VisitAs(o.(*AsNode), visitor)
	case *AssignmentNode:
		return visitor.VisitAssignment(o.(*AssignmentNode), visitor)
	case *EqualNode:
		return visitor.VisitEqual(o.(*EqualNode), visitor)
	case *NotEqualNode:
		return visitor.VisitNotEqual(o.(*NotEqualNode), visitor)
	case *GreaterThanNode:
		return visitor.VisitGreaterThan(o.(*GreaterThanNode), visitor)
	case *GreaterThanOrEqualNode:
		return visitor.VisitGreaterThanOrEqual(o.(*GreaterThanOrEqualNode), visitor)
	case *LessThanNode:
		return visitor.VisitLessThan(o.(*LessThanNode), visitor)
	case *LessThanOrEqualNode:
		return visitor.VisitLessThanOrEqual(o.(*LessThanOrEqualNode), visitor)
	case *InNode:
		return visitor.VisitIn(o.(*InNode), visitor)
	case *LikeNode:
		return visitor.VisitLike(o.(*LikeNode), visitor)
	case *UnlikeNode:
		return visitor.VisitUnlike(o.(*UnlikeNode), visitor)
	case *OrNode:
		return visitor.VisitOr(o.(*OrNode), visitor)
	case *AndNode:
		return visitor.VisitAnd(o.(*AndNode), visitor)
	case *TableNode:
		return visitor.VisitTable(o.(*TableNode), visitor)
	case *AttributeNode:
		return visitor.VisitAttribute(o.(*AttributeNode), visitor)
	case *InnerJoinNode:
		return visitor.VisitInnerJoin(o.(*InnerJoinNode), visitor)
	case *OuterJoinNode:
		return visitor.VisitOuterJoin(o.(*OuterJoinNode), visitor)
	case *JoinSourceNode:
		return visitor.VisitJoinSource(o.(*JoinSourceNode), visitor)
	case *ValuesNode:
		return visitor.VisitValues(o.(*ValuesNode), visitor)
	case *UnionNode:
		return visitor.VisitUnion(o.(*UnionNode), visitor)
	case *IntersectNode:
		return visitor.VisitIntersect(o.(*IntersectNode), visitor)
	case *ExceptNode:
		return visitor.VisitExcept(o.(*ExceptNode), visitor)
	case *BinaryLiteralNode:
		return visitor.VisitBinaryLiteral(o.(*BinaryLiteralNode), visitor)

	// Nary node visitors.
	case *SelectStatementNode:
		return visitor.VisitSelectStatement(o.(*SelectStatementNode), visitor)
	case *InsertStatementNode:
		return visitor.VisitInsertStatement(o.(*InsertStatementNode), visitor)
	case *UpdateStatementNode:
		return visitor.VisitUpdateStatement(o.(*UpdateStatementNode), visitor)
	case *DeleteStatementNode:
		return visitor.VisitDeleteStatement(o.(*DeleteStatementNode), visitor)

	// Function node visitors.
	case *FunctionNode:
		return visitor.VisitFunction(o.(*FunctionNode), visitor)

	// Base visitor.
	default:
		visitor.AppendSqlByte(QUESTION)
		visitor.AppendArg(o)
		return nil
	}
}

// Begin Unary node visitors.

func (_ *ToSqlVisitor) VisitGrouping(o *GroupingNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlByte('(')
	err = visitor.Visit(o.Expr, visitor)
	visitor.AppendSqlByte(')')
	return
}

func (_ *ToSqlVisitor) VisitNot(o *NotNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlStr("NOT(")
	err = visitor.Visit(o.Expr, visitor)
	visitor.AppendSqlByte(')')
	return
}

func (_ *ToSqlVisitor) VisitLiteral(o *LiteralNode, visitor VisitorInterface) (err error) {

	visitor.AppendSqlStr(o.Sql)
	for _, arg := range o.Args {
		visitor.AppendArg(arg)
	}

	return
}

func (_ *ToSqlVisitor) VisitOn(o *OnNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlStr("ON ")
	err = visitor.Visit(o.Expr, visitor)
	return
}

func (_ *ToSqlVisitor) VisitColumn(o *ColumnNode, visitor VisitorInterface) (err error) {
	err = visitor.QuoteColumnName(o.Expr, visitor)
	return
}

func (_ *ToSqlVisitor) VisitStar(o *StarNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlByte(STAR)
	return
}

// TODO everything will be bound - no more in string replacement sql injection danger shit, so VisitBinding may be obsolete
func (_ *ToSqlVisitor) VisitBinding(o *BindingNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlByte(QUESTION)
	return
}

func (_ *ToSqlVisitor) VisitLimit(o *LimitNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlStr("LIMIT ")
	err = visitor.Visit(o.Expr, visitor)
	return
}

func (_ *ToSqlVisitor) VisitOffset(o *OffsetNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlStr("OFFSET ")
	err = visitor.Visit(o.Expr, visitor)
	return
}

func (_ *ToSqlVisitor) VisitHaving(o *HavingNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlStr(" HAVING ")
	err = visitor.Visit(o.Expr, visitor)
	return
}

func (_ *ToSqlVisitor) VisitAscending(o *AscendingNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Expr, visitor)
	visitor.AppendSqlStr(" ASC")
	return
}

func (_ *ToSqlVisitor) VisitDescending(o *DescendingNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Expr, visitor)
	visitor.AppendSqlStr(" DESC")
	return
}

// End Unary node visitors.

// Begin Binary node visitors.

func (_ *ToSqlVisitor) VisitAs(o *AsNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlStr(AS)
	err = visitor.Visit(o.Right, visitor)
	return
}

func (_ *ToSqlVisitor) VisitAssignment(o *AssignmentNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlByte(EQUAL)
	err = visitor.Visit(o.Right, visitor)
	return
}

func (_ *ToSqlVisitor) VisitEqual(o *EqualNode, visitor VisitorInterface) (err error) {
	if nil == o.Right {
		err = visitor.Visit(o.Left, visitor)
		visitor.AppendSqlStr(" IS NULL")
		return
	}
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlByte(EQUAL)
	err = visitor.Visit(o.Right, visitor)
	return
}

func (_ *ToSqlVisitor) VisitNotEqual(o *NotEqualNode, visitor VisitorInterface) (err error) {
	if nil == o.Right {
		err = visitor.Visit(o.Left, visitor)
		visitor.AppendSqlStr(" IS NOT NULL")
		return
	}
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlStr("!=")
	err = visitor.Visit(o.Right, visitor)
	return
}

func (_ *ToSqlVisitor) VisitGreaterThan(o *GreaterThanNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlByte(GT)
	err = visitor.Visit(o.Right, visitor)
	return
}

func (_ *ToSqlVisitor) VisitGreaterThanOrEqual(o *GreaterThanOrEqualNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlStr(GTEQUAL)
	err = visitor.Visit(o.Right, visitor)
	return
}

func (_ *ToSqlVisitor) VisitLessThan(o *LessThanNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlByte(LT)
	err = visitor.Visit(o.Right, visitor)
	return
}

func (_ *ToSqlVisitor) VisitLessThanOrEqual(o *LessThanOrEqualNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlStr(LTEQUAL)
	err = visitor.Visit(o.Right, visitor)
	return
}

func (_ *ToSqlVisitor) VisitIn(o *InNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	vals, ok := o.Right.([]interface{})
	if !ok {
		visitor.AppendSqlStr("-- ERROR --")
		return fmt.Errorf("IN() requires parameters to be []interface{} but is: %#v", o.Right)
	}
	visitor.AppendSqlStr(" IN(")
	for i, val := range vals {
		if i > 0 {
			visitor.AppendSqlByte(COMMA)
		}
		err = visitor.Visit(val, visitor)
		if err != nil {
			return
		}
	}
	visitor.AppendSqlByte(')')
	return
}

func (_ *ToSqlVisitor) VisitLike(o *LikeNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlStr(" LIKE ")
	err = visitor.Visit(o.Right, visitor)
	return
}

func (_ *ToSqlVisitor) VisitUnlike(o *UnlikeNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlStr(" NOT LIKE ")
	err = visitor.Visit(o.Right, visitor)
	return
}

func (_ *ToSqlVisitor) VisitOr(o *OrNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlStr(OR)
	err = visitor.Visit(o.Right, visitor)
	return
}

func (_ *ToSqlVisitor) VisitAnd(o *AndNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlStr(AND)
	err = visitor.Visit(o.Right, visitor)
	return
}

func (_ *ToSqlVisitor) VisitTable(o *TableNode, visitor VisitorInterface) (err error) {
	if o.Alias != nil {
		return visitor.QuoteTableName(o.Alias, visitor)
	}

	return visitor.QuoteTableName(o.Name, visitor)
}

func (_ *ToSqlVisitor) VisitAttribute(o *AttributeNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Table, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlByte(DOT)
	err = visitor.Visit(o.Name, visitor)
	return
}

func (_ *ToSqlVisitor) VisitInnerJoin(o *InnerJoinNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlStr("INNER JOIN ")
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	if nil != o.Right {
		visitor.AppendSqlByte(SPACE)
		err = visitor.Visit(o.Right, visitor)
	}
	return
}

func (_ *ToSqlVisitor) VisitOuterJoin(o *OuterJoinNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlStr("LEFT OUTER JOIN ")
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	if nil != o.Right {
		visitor.AppendSqlByte(SPACE)
		err = visitor.Visit(o.Right, visitor)
	}
	return
}

func (_ *ToSqlVisitor) VisitJoinSource(o *JoinSourceNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	if length := len(o.Right) - 1; 0 <= length {
		visitor.AppendSqlByte(SPACE)
		for index, join := range o.Right {
			err = visitor.Visit(join, visitor)
			if err != nil {
				return
			}
			if index != length {
				visitor.AppendSqlByte(SPACE)
			}
		}
	}
	return
}

func (_ *ToSqlVisitor) VisitValues(o *ValuesNode, visitor VisitorInterface) (err error) {

	if length := len(o.Expressions) - 1; 0 <= length {
		visitor.AppendSqlStr("VALUES (")
		for index, value := range o.Expressions {
			err = visitor.Visit(value, visitor)
			if err != nil {
				return
			}
			if index != length {
				visitor.AppendSqlByte(COMMA)
			}
		}
		visitor.AppendSqlByte(')')
	}

	return
}

func (_ *ToSqlVisitor) VisitUnion(o *UnionNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlByte('(')
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlStr(" UNION ")

	err = visitor.Visit(o.Right, visitor)

	visitor.AppendSqlByte(')')
	return
}

func (_ *ToSqlVisitor) VisitIntersect(o *IntersectNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlByte('(')
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlStr(" INTERSECT ")

	err = visitor.Visit(o.Right, visitor)

	visitor.AppendSqlByte(')')
	return
}

func (_ *ToSqlVisitor) VisitExcept(o *ExceptNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlByte('(')
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlStr(" EXCEPT ")

	err = visitor.Visit(o.Right, visitor)

	visitor.AppendSqlByte(')')
	return
}

func (_ *ToSqlVisitor) VisitBinaryLiteral(o *BinaryLiteralNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlByte(SPACE)

	err = visitor.Visit(o.Right, visitor)

	return
}

// End Binary node visitors.

// Begin Nary node visitors.

func (_ *ToSqlVisitor) VisitSelectCore(o *SelectStatementNode, visitor VisitorInterface) (err error) {

	visitor.AppendSqlStr(SELECT)
	n := len(o.Cols)
	if n == 0 {
		visitor.AppendSqlByte(STAR)
	} else {
		n--
		for i, col := range o.Cols {
			err = visitor.Visit(col, visitor)
			if err != nil {
				return
			}

			if i < n {
				visitor.AppendSqlByte(COMMA)
			}
		}
	}

	visitor.AppendSqlStr(FROM)
	err = visitor.Visit(o.Source, visitor)
	if err != nil {
		return
	}

	if length := len(o.Wheres) - 1; 0 <= length {
		visitor.AppendSqlStr(WHERE)
		for index, where := range o.Wheres {
			err = visitor.Visit(where, visitor)
			if err != nil {
				return
			}
			if index != length {
				visitor.AppendSqlStr(AND)
			}
		}
	}

	if length := len(o.Groups) - 1; 0 <= length {
		visitor.AppendSqlStr(GROUP_BY)
		for index, group := range o.Groups {
			err = visitor.Visit(group, visitor)
			if err != nil {
				return
			}

			if index != length {
				visitor.AppendSqlByte(COMMA)
			}
		}
	}

	if nil != o.Having {
		err = visitor.Visit(o.Having, visitor)
	}
	return
}

func (_ *ToSqlVisitor) VisitSelectStatement(o *SelectStatementNode, visitor VisitorInterface) (err error) {

	// Union, Intersect, Except
	if nil != o.Combinator {
		// TODO: do not change the tree
		combinator := o.Combinator
		o.Combinator = nil
		return visitor.Visit(combinator, visitor)
	}

	err = visitor.VisitSelectCore(o, visitor)
	if err != nil {
		return err
	}

	if length := len(o.Orders) - 1; 0 <= length {
		visitor.AppendSqlStr(ORDER_BY)
		for index, order := range o.Orders {
			err = visitor.Visit(order, visitor)
			if err != nil {
				return
			}
			if index != length {
				visitor.AppendSqlByte(COMMA)
			}
		}
	}

	if nil != o.Limit {
		visitor.AppendSqlByte(SPACE)
		err = visitor.Visit(o.Limit, visitor)
		if err != nil {
			return
		}
	}

	if nil != o.Offset {
		visitor.AppendSqlByte(SPACE)
		err = visitor.Visit(o.Offset, visitor)
	}

	return
}

func (_ *ToSqlVisitor) VisitInsertStatement(o *InsertStatementNode, visitor VisitorInterface) (err error) {

	visitor.AppendSqlStr("INSERT INTO ")
	err = visitor.Visit(o.Table, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlByte(SPACE)

	if length := len(o.Columns) - 1; 0 <= length {
		visitor.AppendSqlByte('(')
		for index, column := range o.Columns {
			switch column.(type) {
			case *LiteralNode, *BindingNode:
				err = visitor.Visit(column, visitor)

			default:
				err = visitor.QuoteColumnName(column, visitor)
			}
			if err != nil {
				return
			}
			if index != length {
				visitor.AppendSqlByte(COMMA)
			} else {
				visitor.AppendSqlByte(')')
			}
		}
		visitor.AppendSqlByte(SPACE)
	}

	err = visitor.Visit(o.Values, visitor)
	if err != nil {
		return
	}

	if nil != o.Returning {
		visitor.AppendSqlStr(" RETURNING ")
		err = visitor.Visit(o.Returning, visitor)
	}

	return
}

func (_ *ToSqlVisitor) VisitUpdateStatement(o *UpdateStatementNode, visitor VisitorInterface) (err error) {

	visitor.AppendSqlStr("UPDATE ")
	err = visitor.Visit(o.Table, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlByte(SPACE)

	if length := len(o.Values) - 1; 0 <= length {
		visitor.AppendSqlStr("SET ")
		for index, assignment := range o.Values {
			err = visitor.Visit(assignment, visitor)
			if err != nil {
				return
			}
			if index != length {
				visitor.AppendSqlByte(COMMA)
			} else {
				visitor.AppendSqlByte(SPACE)
			}
		}
	}

	if length := len(o.Wheres) - 1; 0 <= length {
		visitor.AppendSqlStr("WHERE ")
		for index, filter := range o.Wheres {
			err = visitor.Visit(filter, visitor)
			if err != nil {
				return
			}
			if index != length {
				visitor.AppendSqlStr(AND)
			}
		}
	}

	if nil != o.Limit {
		visitor.AppendSqlByte(SPACE)
		err = visitor.Visit(o.Limit, visitor)
		if err != nil {
			return
		}
	}

	return
}

func (_ *ToSqlVisitor) VisitDeleteStatement(o *DeleteStatementNode, visitor VisitorInterface) (err error) {

	visitor.AppendSqlStr("DELETE FROM ")
	err = visitor.Visit(o.Table, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlByte(SPACE)

	if length := len(o.Wheres) - 1; 0 <= length {
		visitor.AppendSqlStr("WHERE ")
		for index, filter := range o.Wheres {
			err = visitor.Visit(filter, visitor)
			if err != nil {
				return
			}
			if index != length {
				visitor.AppendSqlStr(AND)
			}
		}
	}

	if nil != o.Limit {
		visitor.AppendSqlByte(SPACE)
		err = visitor.Visit(o.Limit, visitor)
		if err != nil {
			return
		}
	}

	return
}

// End Nary node visitors.

func (v *ToSqlVisitor) VisitFunction(o *FunctionNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlStr(o.Name)

	visitor.AppendSqlByte('(')

	if o.Distinct {
		visitor.AppendSqlStr("DISTINCT ")
	}

	if o.Args == nil {
		visitor.AppendSqlByte(STAR)
	} else {
		if length := len(o.Args) - 1; 0 <= length {
			for index, arg := range o.Args {
				err = visitor.Visit(arg, visitor)
				if err != nil {
					return
				}
				if index != length {
					visitor.AppendSqlByte(COMMA)
				}
			}
		}
	}
	visitor.AppendSqlByte(')')

	if nil != o.Alias {
		visitor.AppendSqlStr(AS)
		visitor.QuoteColumnName(o.Alias, visitor)
	}
	return
}

// Begin Helpers.

func (_ *ToSqlVisitor) QuoteTableName(o interface{}, visitor VisitorInterface) (err error) {
	s, ok := o.(string)
	if !ok {
		return fmt.Errorf("ToSqlVisitor.QuoteTableName() expected string but got %#v", o)
	}

	if !VALID_TABLE_NAME_PATTERN.MatchString(s) {
		visitor.AppendSqlStr("-- ERROR --")
		return fmt.Errorf("invalid table name: '%s'", s)
	}

	visitor.AppendSqlByte(QUOTE)
	visitor.AppendSqlStr(s)
	visitor.AppendSqlByte(QUOTE)
	return
}

func (_ *ToSqlVisitor) QuoteColumnName(o interface{}, visitor VisitorInterface) (err error) {
	s, ok := o.(string)
	if !ok {
		return fmt.Errorf("ToSqlVisitor.QuoteColumnName() expected string but got %#v", o)
	}

	if !VALID_COL_NAME_PATTERN.MatchString(s) {
		visitor.AppendSqlStr("-- ERROR --")
		return fmt.Errorf("invalid column name: '%s'", s)
	}

	visitor.AppendSqlByte(QUOTE)
	visitor.AppendSqlStr(s)
	visitor.AppendSqlByte(QUOTE)
	return
}

// End Helpers.
