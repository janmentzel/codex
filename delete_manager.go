package codex

// DeleteManager manages a tree that compiles to a SQL delete statement.
type DeleteManager struct {
	Tree    *DeleteStatementNode // The AST for the SQL DELETE statement.
	adapter adapter              // The SQL adapter.
}

// Appends the expression to the Trees Wheres slice.
func (self *DeleteManager) Delete(expr interface{}) *DeleteManager {
	self.Tree.Wheres = append(self.Tree.Wheres, expr)
	return self
}

// Sets the SQL Adapter.
func (self *DeleteManager) SetAdapter(adapter adapter) *DeleteManager {
	self.adapter = adapter
	return self
}

// ToSql calls a visitor's Accept method based on the manager's SQL adapter.
func (self *DeleteManager) ToSql() (string, []interface{}, error) {
	return VisitorFor(self.adapter).Accept(self.Tree)
}

// DeleteManager factory methods.
func Deletion(relation *RelationNode) (m *DeleteManager) {
	m = new(DeleteManager)
	m.Tree = DeleteStatement(relation)
	m.adapter = relation.adapter
	return
}
