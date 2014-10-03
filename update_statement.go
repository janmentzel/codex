package codex

// UpdateStatement is the base node for SQL Update Statements.
type UpdateStatementNode struct {
	Table  *TableNode    // Pointer to the Table the Delete Statement is acting on.
	Values []interface{} // Values is an array of expressions/nodes.
	Wheres []interface{} // Wheres is an array of expressions/nodes.
	Limit  *LimitNode    // Potential Limit node for limiting the number of rows effected.
}

// UpdateStatementNode factory method.
func UpdateStatement(relation *TableNode) (statement *UpdateStatementNode) {
	statement = new(UpdateStatementNode)
	statement.Table = relation
	statement.Values = make([]interface{}, 0)
	statement.Wheres = make([]interface{}, 0)
	return
}
