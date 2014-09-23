package codex

// InsertManager manages a tree that compiles to a SQL insert statement.
type InsertManager struct {
	Tree    *InsertStatementNode // The AST for the SQL INSERT statement.
	adapter adapter              // The SQL adapter.
}

// Appends the values to the trees Values node
func (self *InsertManager) Insert(values ...interface{}) *InsertManager {
	self.Tree.Values.Expressions = append(self.Tree.Values.Expressions, values...)
	return self
}

// Appends the columns to the trees Columns slice and Values node.
func (self *InsertManager) Into(columns ...interface{}) *InsertManager {
	self.Tree.Values.Columns = append(self.Tree.Values.Columns, columns...)
	self.Tree.Columns = append(self.Tree.Columns, columns...)
	return self
}

// Return sets the InsertStatementNodes Return to the `column` paramenter
// after ensureing it is a ColumnNode.
func (self *InsertManager) Returning(column interface{}) *InsertManager {
	if _, ok := column.(string); ok {
		column = Column(column)
	}

	self.Tree.Returning = column
	return self
}

// Sets the SQL Adapter.
func (self *InsertManager) SetAdapter(adapter adapter) *InsertManager {
	self.adapter = adapter
	return self
}

// ToSql calls a visitor's Accept method based on the manager's SQL adapter.
func (self *InsertManager) ToSql() (string, []interface{}, error) {
	return VisitorFor(self.adapter).Accept(self.Tree)
}

func Insertion(relation *RelationNode) (m *InsertManager) {
	m = new(InsertManager)
	m.Tree = InsertStatement(relation)
	m.adapter = relation.adapter
	return
}
