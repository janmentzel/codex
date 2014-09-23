package codex

// UpdateManager manages a tree that compiles to a SQL update statement.
type UpdateManager struct {
	Tree    *UpdateStatementNode // The AST for the SQL UPDATE statement.
	adapter adapter              // The SQL Engine.
}

// Set appends to the trees Values slice a list of UnqualifiedColumnNodes
// which are to be modified in the query.
func (self *UpdateManager) Set(columns ...interface{}) *UpdateManager {
	for _, column := range columns {
		self.Tree.Values = append(self.Tree.Values, UnqualifiedColumn(column))
	}

	return self
}

// To alters the trees Values slice to be an AssignmentNode, containing the
// column from Set at the same index of the value.
func (self *UpdateManager) To(values ...interface{}) *UpdateManager {
	for index, value := range values {
		if index < len(self.Tree.Values) {
			column := self.Tree.Values[index]
			self.Tree.Values[index] = Assignment(column, value)
		}
	}

	return self
}

// Appends an expression to the current tree's Wheres slice,
// typically a comparison, i.e. 1 = 1
func (self *UpdateManager) Where(expr interface{}) *UpdateManager {
	self.Tree.Wheres = append(self.Tree.Wheres, expr)
	return self
}

// Sets the Tree's Limit to the given integer.
func (self *UpdateManager) Limit(expr interface{}) *UpdateManager {
	self.Tree.Limit = Limit(expr)
	return self
}

// Sets the SQL Adapter.
func (self *UpdateManager) SetAdapter(adapter adapter) *UpdateManager {
	self.adapter = adapter
	return self
}

// ToSql calls a visitor's Accept method based on the manager's SQL adapter.
func (self *UpdateManager) ToSql() (string, []interface{}, error) {
	return VisitorFor(self.adapter).Accept(self.Tree)
}

// UpdateManager factory method.
func Modification(relation *RelationNode) (modification *UpdateManager) {
	modification = new(UpdateManager)
	modification.Tree = UpdateStatement(relation)
	return
}
