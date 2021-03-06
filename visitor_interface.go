package codex

type VisitorInterface interface {
	// Base methods.
	Accept(interface{}) (string, []interface{}, error)
	Visit(interface{}, VisitorInterface) error

	// Collector methods
	AppendSqlStr(string)
	AppendSqlByte(byte)
	AppendArg(interface{})

	// Unary node visitors.
	VisitGrouping(*GroupingNode, VisitorInterface) error
	VisitNot(*NotNode, VisitorInterface) error
	VisitLiteral(*LiteralNode, VisitorInterface) error
	VisitOn(*OnNode, VisitorInterface) error
	VisitColumn(*ColumnNode, VisitorInterface) error
	VisitStar(*StarNode, VisitorInterface) error
	VisitBinding(*BindingNode, VisitorInterface) error
	VisitLimit(*LimitNode, VisitorInterface) error
	VisitOffset(*OffsetNode, VisitorInterface) error
	VisitHaving(*HavingNode, VisitorInterface) error
	VisitAscending(*AscendingNode, VisitorInterface) error
	VisitDescending(*DescendingNode, VisitorInterface) error

	// Binary node visitors.
	VisitAs(*AsNode, VisitorInterface) error
	VisitAssignment(*AssignmentNode, VisitorInterface) error
	VisitEqual(*EqualNode, VisitorInterface) error
	VisitNotEqual(*NotEqualNode, VisitorInterface) error
	VisitGreaterThan(*GreaterThanNode, VisitorInterface) error
	VisitGreaterThanOrEqual(*GreaterThanOrEqualNode, VisitorInterface) error
	VisitLessThan(*LessThanNode, VisitorInterface) error
	VisitLessThanOrEqual(*LessThanOrEqualNode, VisitorInterface) error
	VisitIn(*InNode, VisitorInterface) error
	VisitLike(*LikeNode, VisitorInterface) error
	VisitUnlike(*UnlikeNode, VisitorInterface) error
	VisitOr(*OrNode, VisitorInterface) error
	VisitAnd(*AndNode, VisitorInterface) error
	VisitTable(*TableNode, VisitorInterface) error
	VisitAttribute(*AttributeNode, VisitorInterface) error
	VisitInnerJoin(*InnerJoinNode, VisitorInterface) error
	VisitOuterJoin(*OuterJoinNode, VisitorInterface) error
	VisitJoinSource(*JoinSourceNode, VisitorInterface) error
	VisitValues(*ValuesNode, VisitorInterface) error
	VisitUnion(*UnionNode, VisitorInterface) error
	VisitIntersect(*IntersectNode, VisitorInterface) error
	VisitExcept(*ExceptNode, VisitorInterface) error
	VisitBinaryLiteral(*BinaryLiteralNode, VisitorInterface) error

	// Nary node visitors.
	VisitSelectCore(*SelectStatementNode, VisitorInterface) error
	VisitSelectStatement(*SelectStatementNode, VisitorInterface) error
	VisitInsertStatement(*InsertStatementNode, VisitorInterface) error
	VisitUpdateStatement(*UpdateStatementNode, VisitorInterface) error
	VisitDeleteStatement(*DeleteStatementNode, VisitorInterface) error

	// Function node visitor.
	VisitFunction(*FunctionNode, VisitorInterface) error

	// Helpers.
	QuoteTableName(interface{}, VisitorInterface) error
	QuoteColumnName(interface{}, VisitorInterface) error
}
