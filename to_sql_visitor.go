package codex

import (
	"errors"
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
func NewToSqlVisitor() *ToSqlVisitor {
	return &ToSqlVisitor{NewCollector()}
}

func (v *ToSqlVisitor) Accept(o interface{}) (string, []interface{}, error) {
	err := v.Visit(o, v)

	return v.String(), v.Args(), err
}

func (_ *ToSqlVisitor) Visit(o interface{}, visitor VisitorInterface) error {

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
	case *UnqualifiedColumnNode:
		return visitor.VisitUnqualifiedColumn(o.(*UnqualifiedColumnNode), visitor)
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
	case *EngineNode:
		return visitor.VisitEngine(o.(*EngineNode), visitor)
	case *IndexNameNode:
		return visitor.VisitIndexName(o.(*IndexNameNode), visitor)

	// Binary node visitors.
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
	case *LikeNode:
		return visitor.VisitLike(o.(*LikeNode), visitor)
	case *UnlikeNode:
		return visitor.VisitUnlike(o.(*UnlikeNode), visitor)
	case *OrNode:
		return visitor.VisitOr(o.(*OrNode), visitor)
	case *AndNode:
		return visitor.VisitAnd(o.(*AndNode), visitor)
	case *RelationNode:
		return visitor.VisitRelation(o.(*RelationNode), visitor)
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
	case *UnexistingColumnNode:
		return visitor.VisitUnexistingColumn(o.(*UnexistingColumnNode), visitor)
	case *ExistingColumnNode:
		return visitor.VisitExistingColumn(o.(*ExistingColumnNode), visitor)

	// Nary node visitors.
	// case *ConstraintNode:
	// 	return visitor.VisitConstraint(o.(*ConstraintNode), visitor)
	// case *NotNullNode:
	// 	return visitor.VisitNotNull(o.(*NotNullNode), visitor)
	// case *UniqueNode:
	// 	return visitor.VisitUnique(o.(*UniqueNode), visitor)
	// case *PrimaryKeyNode:
	// 	return visitor.VisitPrimaryKey(o.(*PrimaryKeyNode), visitor)
	// case *ForeignKeyNode:
	// 	return visitor.VisitForeignKey(o.(*ForeignKeyNode), visitor)
	// case *CheckNode:
	// 	return visitor.VisitCheck(o.(*CheckNode), visitor)
	// case *DefaultNode:
	// 	return visitor.VisitDefault(o.(*DefaultNode), visitor)
	case *SelectCoreNode:
		return visitor.VisitSelectCore(o.(*SelectCoreNode), visitor)
	case *SelectStatementNode:
		return visitor.VisitSelectStatement(o.(*SelectStatementNode), visitor)
	case *InsertStatementNode:
		return visitor.VisitInsertStatement(o.(*InsertStatementNode), visitor)
	case *UpdateStatementNode:
		return visitor.VisitUpdateStatement(o.(*UpdateStatementNode), visitor)
	case *DeleteStatementNode:
		return visitor.VisitDeleteStatement(o.(*DeleteStatementNode), visitor)
	// case *AlterStatementNode:
	// 	return visitor.VisitAlterStatement(o.(*AlterStatementNode), visitor)
	// case *CreateStatementNode:
	// 	return visitor.VisitCreateStatement(o.(*CreateStatementNode), visitor)

	// Function node visitors.
	case *CountNode:
		return visitor.VisitCount(o.(*CountNode), visitor)
	case *AverageNode:
		return visitor.VisitAverage(o.(*AverageNode), visitor)
	case *SumNode:
		return visitor.VisitSum(o.(*SumNode), visitor)
	case *MaximumNode:
		return visitor.VisitMaximum(o.(*MaximumNode), visitor)
	case *MinimumNode:
		return visitor.VisitMinimum(o.(*MinimumNode), visitor)

	// SQL constant visitors.
	case Type:
		return visitor.VisitSqlType(o.(Type), visitor)

	// Base visitors.
	case string:
		return visitor.VisitString(o, visitor)
	case int, int16, int32, int64:
		return visitor.VisitInteger(o, visitor)
	case float32, float64:
		return visitor.VisitFloat(o, visitor)
	case bool:
		return visitor.VisitBool(o, visitor)
	default:
		return errors.New(fmt.Sprintf("No visitor method for <%T>.", o))
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
	visitor.AppendSqlStr("NOT (")
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

func (_ *ToSqlVisitor) VisitUnqualifiedColumn(o *UnqualifiedColumnNode, visitor VisitorInterface) (err error) {
	return visitor.QuoteColumnName(o.Expr, visitor)
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

func (_ *ToSqlVisitor) VisitEngine(o *EngineNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Expr, visitor)
	return
}

func (_ *ToSqlVisitor) VisitIndexName(o *IndexNameNode, visitor VisitorInterface) (err error) {
	err = visitor.QuoteColumnName(o.Expr, visitor)
	return
}

// End Unary node visitors.

// Begin Binary node visitors.

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

func (_ *ToSqlVisitor) VisitRelation(o *RelationNode, visitor VisitorInterface) (err error) {
	if o.Alias != nil {
		return visitor.QuoteTableName(o.Alias, visitor)
	}

	return visitor.QuoteTableName(o.Name, visitor)
}

func (_ *ToSqlVisitor) VisitAttribute(o *AttributeNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Relation, visitor)
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

func (_ *ToSqlVisitor) VisitUnexistingColumn(o *UnexistingColumnNode, visitor VisitorInterface) (err error) {
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlByte(SPACE)

	err = visitor.Visit(o.Right, visitor)

	return
}

func (_ *ToSqlVisitor) VisitExistingColumn(o *ExistingColumnNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlStr("ALTER COLUMN ")
	err = visitor.Visit(o.Left, visitor)
	if err != nil {
		return
	}
	visitor.AppendSqlStr(" TYPE ")

	err = visitor.Visit(o.Right, visitor)

	return
}

// End Binary node visitors.

// Begin Nary node visitors.

// func (_ *ToSqlVisitor) VisitConstraint(o *ConstraintNode, visitor VisitorInterface) (err error) {
// 	panic("VisitConstraint is unimplemented.")
// }

// func (_ *ToSqlVisitor) VisitNotNull(o *NotNullNode, visitor VisitorInterface) (err error) {
// 	visitor.AppendSqlStr("ALTER ")
// 	err = visitor.FormatConstraintColumns(o.Columns, visitor)
// 	visitor.AppendSqlStr(" SET NOT NULL")
// 	return
// }

// func (_ *ToSqlVisitor) VisitUnique(o *UniqueNode, visitor VisitorInterface) (err error) {
// 	visitor.AppendSqlStr("ADD ")
// 	// Optional index name.
// 	if 0 < len(o.Options) {
// 		expr := o.Options[0]
// 		if _, ok := expr.(string); ok {
// 			expr = IndexName(expr)
// 		}

// 		visitor.AppendSqlStr("CONSTRAINT ")
// 		err = visitor.Visit(expr, visitor)
// 		if err != nil {
// 			return
// 		}
// 		visitor.AppendSqlByte(SPACE)
// 	}

// 	visitor.AppendSqlStr("UNIQUE(")
// 	err = visitor.FormatConstraintColumns(o.Columns, visitor)
// 	visitor.AppendSqlByte(')')

// 	return
// }

// func (_ *ToSqlVisitor) VisitPrimaryKey(o *PrimaryKeyNode, visitor VisitorInterface) (err error) {
// 	visitor.AppendSqlStr("ADD ")
// 	// Optional index name.
// 	if 0 < len(o.Options) {
// 		expr := o.Options[0]
// 		if _, ok := expr.(string); ok {
// 			expr = IndexName(expr)
// 		}

// 		visitor.AppendSqlStr("CONSTRAINT ")
// 		err = visitor.Visit(expr, visitor)
// 		if err != nil {
// 			return
// 		}
// 		visitor.AppendSqlByte(SPACE)
// 	}

// 	visitor.AppendSqlStr("PRIMARY KEY(")
// 	err = visitor.FormatConstraintColumns(o.Columns, visitor)
// 	visitor.AppendSqlByte(')')

// 	return
// }

// func (_ *ToSqlVisitor) VisitForeignKey(o *ForeignKeyNode, visitor VisitorInterface) (err error) {
// 	if 0 >= len(o.Options) {
// 		return errors.New("Missing column REFERENCE name for FOREIGN KEY constraint.")
// 	}

// 	visitor.AppendSqlStr("ADD ")

// 	// For FOREIGN KEY, index name is optional, REFERENCES is not.
// 	//
// 	// No index name ex.
// 	//
// 	//  CreateTable("orders").
// 	//    AddColumn("user_id").
// 	//    AddConstraint("user_id", FOREIGN_KEY, "users")
// 	//
// 	// With option index name ex.
// 	//
// 	//  CreateTable("orders").
// 	//    AddColumn("user_id").
// 	//    AddConstraint("user_id", FOREIGN_KEY, "users_fkey", "users")
// 	if 1 < len(o.Options) {
// 		expr := o.Options[0]
// 		if _, ok := expr.(string); ok {
// 			expr = IndexName(expr)
// 		}

// 		// Remove this item from the array, avoiding any potential memory leak.
// 		// https://code.google.com/p/go-wiki/wiki/SliceTricks
// 		length := len(o.Options) - 1
// 		copy(o.Options[0:], o.Options[1:])
// 		o.Options[length] = nil
// 		o.Options = o.Options[:length]

// 		visitor.AppendSqlStr("CONSTRAINT ")
// 		err = visitor.Visit(expr, visitor)
// 		if err != nil {
// 			return
// 		}
// 		visitor.AppendSqlByte(SPACE)
// 	}

// 	visitor.AppendSqlStr("FOREIGN KEY(")
// 	err = visitor.FormatConstraintColumns(o.Columns, visitor)
// 	if err != nil {
// 		return
// 	}
// 	visitor.AppendSqlByte(')')

// 	// If option is not here, user didn't do it right, but don't dereference and panic.
// 	if 0 < len(o.Options) {
// 		relation := o.Options[0]
// 		if _, ok := relation.(string); ok {
// 			relation = Relation(relation.(string))
// 		}

// 		visitor.AppendSqlStr(" REFERENCES ")
// 		err = visitor.Visit(relation, visitor)
// 		if err != nil {
// 			return
// 		}
// 	}

// 	return
// }

// func (_ *ToSqlVisitor) VisitCheck(o *CheckNode, visitor VisitorInterface) (err error) {
// 	panic("VisitCheck is unimplemented.")
// }

// func (_ *ToSqlVisitor) VisitDefault(o *DefaultNode, visitor VisitorInterface) (err error) {
// 	str := fmt.Sprintf("ALTER %v SET DEFAULT", visitor.FormatConstraintColumns(o.Columns, visitor))

// 	if 0 < len(o.Options) {
// 		str = fmt.Sprintf("%v%v%v", str, SPACE, visitor.Visit(o.Options[0], visitor))
// 	}

// 	return str
// }

func (_ *ToSqlVisitor) VisitSelectCore(o *SelectCoreNode, visitor VisitorInterface) (err error) {

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

	if nil != o.Combinator {
		combinator := o.Combinator
		o.Combinator = nil
		return visitor.Visit(combinator, visitor)
	}

	for _, core := range o.Cores {
		err = visitor.Visit(core, visitor)
		if err != nil {
			return
		}
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
	err = visitor.Visit(o.Relation, visitor)
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
				visitor.AppendSqlByte(SPACE)
			}
		}
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
	err = visitor.Visit(o.Relation, visitor)
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

	return
}

func (_ *ToSqlVisitor) VisitDeleteStatement(o *DeleteStatementNode, visitor VisitorInterface) (err error) {

	visitor.AppendSqlStr("DELETE FROM ")
	err = visitor.Visit(o.Relation, visitor)
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

	return
}

// func (_ *ToSqlVisitor) VisitAlterStatement(o *AlterStatementNode, visitor VisitorInterface) (err error) {
// 	str := ""

// 	columns := len(o.RemovedColumns) - 1

// 	for i, column := range o.RemovedColumns {
// 		str = fmt.Sprintf("%vALTER TABLE %v DROP COLUMN %v;", visitor.Visit(o.Relation, visitor), visitor.Visit(column, visitor))

// 		if i != columns {
// 			str = fmt.Sprintf("%v\n", str)
// 		}
// 	}

// 	indicies := len(o.RemovedIndicies) - 1

// 	for i, index := range o.RemovedIndicies {
// 		str = fmt.Sprintf("%vALTER TABLE %v DROP INDEX %v;", visitor.Visit(o.Relation, visitor), visitor.Visit(index, visitor))

// 		if i != indicies {
// 			str = fmt.Sprintf("%v\n", str)
// 		}
// 	}

// 	columns = len(o.UnexistingColumns) - 1

// 	for i, column := range o.UnexistingColumns {
// 		str = fmt.Sprintf("%vALTER TABLE %v ADD %v", str, visitor.Visit(o.Relation, visitor), visitor.Visit(column, visitor))

// 		if i != columns {
// 			str = fmt.Sprintf("%v\n", str)
// 		}
// 	}

// 	if 0 <= columns {
// 		str = fmt.Sprintf("%v\n", str)
// 	}

// 	columns = len(o.ModifiedColumns) - 1

// 	for i, column := range o.ModifiedColumns {
// 		str = fmt.Sprintf("%vALTER TABLE %v %v;", str, visitor.Visit(o.Relation, visitor), visitor.Visit(column, visitor))

// 		if i != columns {
// 			str = fmt.Sprintf("%v\n", str)
// 		}
// 	}

// 	if 0 <= columns {
// 		str = fmt.Sprintf("%v\n", str)
// 	}

// 	constraints := len(o.Constraints) - 1

// 	for i, constraint := range o.Constraints {
// 		str = fmt.Sprintf("%vALTER TABLE %v %v;", str, visitor.Visit(o.Relation, visitor), visitor.Visit(constraint, visitor))

// 		if i != constraints {
// 			str = fmt.Sprintf("%v\n", str)
// 		}
// 	}

// 	return strings.Trim(str, "\n")
// }

// func (_ *ToSqlVisitor) VisitCreateStatement(o *CreateStatementNode, visitor VisitorInterface) (err error) {
// 	str := fmt.Sprintf("CREATE TABLE %v (", visitor.Visit(o.Relation, visitor))

// 	if columns := len(o.UnexistingColumns) - 1; 0 <= columns {
// 		str = fmt.Sprintf("%v\n", str)

// 		for i, column := range o.UnexistingColumns {
// 			str = fmt.Sprintf("%v\t%v", str, visitor.Visit(column, visitor))

// 			if i != columns {
// 				str = fmt.Sprintf("%v,", str)
// 			}

// 			str = fmt.Sprintf("%v\n", str)
// 		}
// 	} else {
// 		// FIXME: Include default column if none provided.
// 	}

// 	str = fmt.Sprintf("%v);\n", str)

// 	constraints := len(o.Constraints) - 1

// 	for i, constraint := range o.Constraints {
// 		str = fmt.Sprintf("%vALTER TABLE %v %v;", str, visitor.Visit(o.Relation, visitor), visitor.Visit(constraint, visitor))

// 		if i != constraints {
// 			str = fmt.Sprintf("%v\n", str)
// 		}
// 	}

// 	return strings.Trim(str, "\n")
// }

// End Nary node visitors.

// Begin Function node visitors.

// End Function node visitors.

func (_ *ToSqlVisitor) VisitCount(o *CountNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlStr("COUNT(")

	if o.Distinct {
		visitor.AppendSqlStr("DISTINCT ")
	}

	if o.Expressions == nil {
		visitor.AppendSqlByte(STAR)
	} else {
		n := len(o.Expressions) - 1
		for i, expression := range o.Expressions {
			err = visitor.Visit(expression, visitor)
			if err != nil {
				return
			}
			if i != n {
				visitor.AppendSqlByte(COMMA)
			}
		}
	}

	visitor.AppendSqlByte(')')

	if nil != o.Alias {
		visitor.AppendSqlStr(AS)
		err = visitor.Visit(o.Alias, visitor)
	}

	return
}

func (_ *ToSqlVisitor) VisitAverage(o *AverageNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlStr("AVG(")

	if o.Distinct {
		visitor.AppendSqlStr("DISTINCT ")
	}

	if length := len(o.Expressions) - 1; 0 <= length {
		for index, expression := range o.Expressions {
			err = visitor.Visit(expression, visitor)
			if err != nil {
				return
			}
			if index != length {
				visitor.AppendSqlByte(COMMA)
			}
		}
		visitor.AppendSqlByte(')')
	}

	if nil != o.Alias {
		visitor.AppendSqlStr(AS)
		err = visitor.Visit(o.Alias, visitor)
	}

	return
}

func (_ *ToSqlVisitor) VisitSum(o *SumNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlStr("SUM(")

	if o.Distinct {
		visitor.AppendSqlStr("DISTINCT ")
	}

	if length := len(o.Expressions) - 1; 0 <= length {
		for index, expression := range o.Expressions {
			err = visitor.Visit(expression, visitor)
			if err != nil {
				return
			}
			if index != length {
				visitor.AppendSqlByte(COMMA)
			}
		}
		visitor.AppendSqlByte(')')
	}

	if nil != o.Alias {
		visitor.AppendSqlStr(AS)
		err = visitor.Visit(o.Alias, visitor)
	}

	return
}

func (_ *ToSqlVisitor) VisitMaximum(o *MaximumNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlStr("MAX(")

	if o.Distinct {
		visitor.AppendSqlStr("DISTINCT ")
	}

	if length := len(o.Expressions) - 1; 0 <= length {
		for index, expression := range o.Expressions {
			err = visitor.Visit(expression, visitor)
			if err != nil {
				return
			}
			if index != length {
				visitor.AppendSqlByte(COMMA)
			}
		}
		visitor.AppendSqlByte(')')
	}

	if nil != o.Alias {
		visitor.AppendSqlStr(AS)
		err = visitor.Visit(o.Alias, visitor)
	}

	return
}

func (_ *ToSqlVisitor) VisitMinimum(o *MinimumNode, visitor VisitorInterface) (err error) {
	visitor.AppendSqlStr("MIN(")

	if o.Distinct {
		visitor.AppendSqlStr("DISTINCT ")
	}

	if length := len(o.Expressions) - 1; 0 <= length {
		for index, expression := range o.Expressions {
			err = visitor.Visit(expression, visitor)
			if err != nil {
				return
			}
			if index != length {
				visitor.AppendSqlByte(COMMA)
			}
		}
		visitor.AppendSqlByte(')')
	}

	if nil != o.Alias {
		visitor.AppendSqlStr(AS)
		err = visitor.Visit(o.Alias, visitor)
	}

	return
}

// Begin SQL constant visitors.

func (_ *ToSqlVisitor) VisitSqlType(o Type, visitor VisitorInterface) (err error) {
	switch o {
	case String:
		visitor.AppendSqlStr("varchar(255)")
	case Text:
		visitor.AppendSqlStr("text")
	case Boolean:
		visitor.AppendSqlStr("boolean")
	case Integer:
		visitor.AppendSqlStr("integer")
	case Float:
		visitor.AppendSqlStr("float")
	case Decimal:
		visitor.AppendSqlStr("decimal")
	case Date:
		visitor.AppendSqlStr("date")
	case Time:
		visitor.AppendSqlStr("time")
	case Datetime:
		visitor.AppendSqlStr("datetime")
	case Timestamp:
		visitor.AppendSqlStr("timestamp")
	default:
		err = errors.New(fmt.Sprintf("Unkown SQL Type constant: %T", o))
	}
	return
}

// End SQL constant visitors.

// Begin Base visitors.

func (_ *ToSqlVisitor) VisitString(o interface{}, visitor VisitorInterface) (err error) {
	visitor.AppendSqlByte(QUESTION)
	visitor.AppendArg(o)
	return
}

func (_ *ToSqlVisitor) VisitInteger(o interface{}, visitor VisitorInterface) (err error) {
	visitor.AppendSqlByte(QUESTION)
	visitor.AppendArg(o)
	return
}

func (_ *ToSqlVisitor) VisitFloat(o interface{}, visitor VisitorInterface) (err error) {
	visitor.AppendSqlByte(QUESTION)
	visitor.AppendArg(o)
	return
}

func (_ *ToSqlVisitor) VisitBool(o interface{}, visitor VisitorInterface) (err error) {
	visitor.AppendSqlByte(QUESTION)
	visitor.AppendArg(o)
	return
}

// End Base visitors.

// Begin Helpers.

func (_ *ToSqlVisitor) QuoteTableName(o interface{}, visitor VisitorInterface) (err error) {
	// TODO remove Sprintf  with (o).(Relation).Name
	visitor.AppendSqlStr(fmt.Sprintf(`"%v"`, o))
	return
}

func (_ *ToSqlVisitor) QuoteColumnName(o interface{}, visitor VisitorInterface) (err error) {
	// TODO remove Sprintf  with (o).(Table).Name
	visitor.AppendSqlStr(fmt.Sprintf(`"%v"`, o))
	return
}

// FIXME: Not sure if I like this as a solution to indexing
// on multiple columns.
func (_ *ToSqlVisitor) FormatConstraintColumns(cols []interface{}, visitor VisitorInterface) (err error) {

	n := len(cols)
	for i, col := range cols {
		err = visitor.Visit(col, visitor)
		if err != nil {
			return
		}

		if i != n {
			visitor.AppendSqlByte(COMMA)
		}
	}

	return
}

// End Helpers.
