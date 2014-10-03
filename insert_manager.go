package codex

// InsertManager manages a tree that compiles to a SQL insert statement.
type InsertManager struct {
	Tree    *InsertStatementNode // The AST for the SQL INSERT statement.
	Adapter adapter              // The SQL adapter.
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

// Selection returns a *SelectManager while keeping
// Table and adapter
func (self *InsertManager) Selection() *SelectManager {
	m := Selection(self.Tree.Table)
	m.Adapter = self.Adapter
	return m
}

// Modification returns an *UpdateManager wwhile keeping
// Table and adapter
func (self *InsertManager) Modification() *UpdateManager {
	m := Modification(self.Tree.Table)
	m.Adapter = self.Adapter
	return m
}

// Deletion returns a *DeleteManager while keeping
// Table and adapter
func (self *InsertManager) Deletion() *DeleteManager {
	m := Deletion(self.Tree.Table)
	m.Adapter = self.Adapter
	return m
}

// ToSql calls a visitor's Accept method based on the manager's SQL adapter.
func (self *InsertManager) ToSql() (string, []interface{}, error) {
	return VisitorFor(self.Adapter).Accept(self.Tree)
}

func Insertion(relation *TableNode) (m *InsertManager) {
	m = new(InsertManager)
	m.Tree = InsertStatement(relation)
	m.Adapter = relation.Adapter
	return
}
