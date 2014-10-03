package codex

// DeleteStatement is the base node for SQL Delete Statements.
type DeleteStatementNode struct {
	Table  *TableNode    // Pointer to the Table the Delete Statement is acting on.
	Wheres []interface{} // Wheres is an array of expressions/nodes.
	Limit  *LimitNode    // Potential Limit node for limiting the number of rows effected.
}

// DeleteStatementNode factory method.
func DeleteStatement(relation *TableNode) (statement *DeleteStatementNode) {
	statement = new(DeleteStatementNode)
	statement.Table = relation
	return
}
