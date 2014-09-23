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
	VisitUnqualifiedColumn(*UnqualifiedColumnNode, VisitorInterface) error
	VisitLimit(*LimitNode, VisitorInterface) error
	VisitOffset(*OffsetNode, VisitorInterface) error
	VisitHaving(*HavingNode, VisitorInterface) error
	VisitAscending(*AscendingNode, VisitorInterface) error
	VisitDescending(*DescendingNode, VisitorInterface) error
	VisitEngine(*EngineNode, VisitorInterface) error
	VisitIndexName(*IndexNameNode, VisitorInterface) error

	// Binary node visitors.
	VisitAssignment(*AssignmentNode, VisitorInterface) error
	VisitEqual(*EqualNode, VisitorInterface) error
	VisitNotEqual(*NotEqualNode, VisitorInterface) error
	VisitGreaterThan(*GreaterThanNode, VisitorInterface) error
	VisitGreaterThanOrEqual(*GreaterThanOrEqualNode, VisitorInterface) error
	VisitLessThan(*LessThanNode, VisitorInterface) error
	VisitLessThanOrEqual(*LessThanOrEqualNode, VisitorInterface) error
	VisitLike(*LikeNode, VisitorInterface) error
	VisitUnlike(*UnlikeNode, VisitorInterface) error
	VisitOr(*OrNode, VisitorInterface) error
	VisitAnd(*AndNode, VisitorInterface) error
	VisitRelation(*RelationNode, VisitorInterface) error
	VisitAttribute(*AttributeNode, VisitorInterface) error
	VisitInnerJoin(*InnerJoinNode, VisitorInterface) error
	VisitOuterJoin(*OuterJoinNode, VisitorInterface) error
	VisitJoinSource(*JoinSourceNode, VisitorInterface) error
	VisitValues(*ValuesNode, VisitorInterface) error
	VisitUnion(*UnionNode, VisitorInterface) error
	VisitExcept(*ExceptNode, VisitorInterface) error
	VisitIntersect(*IntersectNode, VisitorInterface) error
	VisitUnexistingColumn(*UnexistingColumnNode, VisitorInterface) error
	VisitExistingColumn(*ExistingColumnNode, VisitorInterface) error

	// Nary node visitors.
	//VisitConstraint(*ConstraintNode, VisitorInterface) error
	//VisitNotNull(*NotNullNode, VisitorInterface) error
	//VisitUnique(*UniqueNode, VisitorInterface) error
	//VisitPrimaryKey(*PrimaryKeyNode, VisitorInterface) error
	//VisitForeignKey(*ForeignKeyNode, VisitorInterface) error
	//VisitCheck(*CheckNode, VisitorInterface) error
	//VisitDefault(*DefaultNode, VisitorInterface) error
	VisitSelectCore(*SelectCoreNode, VisitorInterface) error
	VisitSelectStatement(*SelectStatementNode, VisitorInterface) error
	VisitInsertStatement(*InsertStatementNode, VisitorInterface) error
	VisitUpdateStatement(*UpdateStatementNode, VisitorInterface) error
	VisitDeleteStatement(*DeleteStatementNode, VisitorInterface) error
	//VisitAlterStatement(*AlterStatementNode, VisitorInterface) error
	//VisitCreateStatement(*CreateStatementNode, VisitorInterface) error

	// Function node visitors.
	VisitCount(*CountNode, VisitorInterface) error
	VisitAverage(*AverageNode, VisitorInterface) error
	VisitSum(*SumNode, VisitorInterface) error
	VisitMaximum(*MaximumNode, VisitorInterface) error
	VisitMinimum(*MinimumNode, VisitorInterface) error

	// SQL constant visitors.
	VisitSqlType(Type, VisitorInterface) error

	// Base visitors.
	VisitString(interface{}, VisitorInterface) error
	VisitInteger(interface{}, VisitorInterface) error
	VisitFloat(interface{}, VisitorInterface) error
	VisitBool(interface{}, VisitorInterface) error

	// Helpers.
	QuoteTableName(interface{}, VisitorInterface) error
	QuoteColumnName(interface{}, VisitorInterface) error
	FormatConstraintColumns([]interface{}, VisitorInterface) error
}
