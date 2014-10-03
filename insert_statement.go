package codex

// InsertStatement is the base node for SQL Insert Statements.
type InsertStatementNode struct {
	Table     *TableNode    // Pointer to the Table the Insert Statement is acting on.
	Columns   []interface{} // Columns the Insert Statement is effecting.
	Returning interface{}   // Columns to return after the Insert Statement is executed.
	Values    *ValuesNode   // Pointer to the Values for insertion.
}

// InsertStatementNode factory method.
func InsertStatement(relation *TableNode) (statement *InsertStatementNode) {
	statement = new(InsertStatementNode)
	statement.Table = relation
	statement.Columns = make([]interface{}, 0)
	statement.Values = Values()
	return
}
