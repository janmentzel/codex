package codex

// SelectStatement is the base node for SQL Select Statements.
type SelectStatementNode struct {
	Relation   *RelationNode   // Pointer to the relation the SelectCore is acting on.
	Source     *JoinSourceNode // JoinSouce for joining other SQL tables.
	Cols       []interface{}   // Cols is an array, normally columns found on the SQL table.
	Wheres     []interface{}   // Wheres is an array of filters for the acting on the SelectCore.
	Groups     []interface{}   // GROUP BY nodes.
	Having     interface{}     // HAVING expression.
	Orders     []interface{}   // An array of nodes for ordering results.
	Combinator interface{}     // Potential Union/Intersect/Except node.
	Limit      *LimitNode      // Potential Limit node for limiting the number of results returned.
	Offset     *OffsetNode     // Potential Offset node for skipping records.
}

// SelectStatementNode factory method.
func SelectStatement(relation *RelationNode) (stm *SelectStatementNode) {
	stm = new(SelectStatementNode)
	stm.Relation = relation
	stm.Source = JoinSource(relation)
	stm.Wheres = make([]interface{}, 0)
	stm.Cols = make([]interface{}, 0)
	stm.Groups = make([]interface{}, 0)
	stm.Orders = make([]interface{}, 0)
	return
}
