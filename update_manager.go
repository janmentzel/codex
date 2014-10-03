package codex

// UpdateManager manages a tree that compiles to a SQL update statement.
type UpdateManager struct {
	Tree    *UpdateStatementNode // The AST for the SQL UPDATE statement.
	Adapter adapter              // The SQL Engine.
}

var _ Scoper = (*UpdateManager)(nil)

func (self *UpdateManager) Scopes(scopes ...ScopeFunc) *UpdateManager {
	for _, scope := range scopes {
		scope(self)
	}
	return self
}

func (self *UpdateManager) Scope(expr interface{}, args ...interface{}) {
	self.Where(expr, args...)
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

// Where appends an sql WHERE condition to the current tree's Wheres slice,
//
//   Where("a")                             // no   args -> Group(Literal("a"))
//   Where("a = ?", 123)                    // with args -> Group(Literal("a = ?", 123))
//   Where("a = ? AND b = ?", 123, true)    // with args -> Group(Literal("a = ? AND b = ?", 123, true))
//   Where(Equal(Column("a"), Column("b"))) // no   args -> Group(Equal(Column("a"), Column("b")))
func (self *UpdateManager) Where(expr interface{}, args ...interface{}) *UpdateManager {

	if str, ok := expr.(string); ok {
		expr = Literal(str, args...)
	}
	// enclose expr in Grouping - except if expr is already a Grouping
	if _, ok := expr.(*GroupingNode); !ok {
		expr = Grouping(expr)
	}

	self.Tree.Wheres = append(self.Tree.Wheres, expr)
	return self
}

// Sets the Tree's Limit to the given integer.
func (self *UpdateManager) Limit(expr interface{}) *UpdateManager {
	self.Tree.Limit = Limit(expr)
	return self
}

// Selection returns a *SelectManager while keeping
// wheres, limit and adapter
func (self *UpdateManager) Selection() *SelectManager {
	m := Selection(self.Tree.Table)
	m.Tree.Wheres = self.Tree.Wheres
	m.Tree.Limit = self.Tree.Limit
	m.Adapter = self.Adapter
	return m
}

// Insertion returns a *InsertManager while keeping
// relation and adapter
func (self *UpdateManager) Insertion() *InsertManager {
	m := Insertion(self.Tree.Table)
	m.Adapter = self.Adapter
	return m
}

// Deletion returns a *DeleteManager while keeping
// wheres, limit and adapter
func (self *UpdateManager) Deletion() *DeleteManager {
	m := Deletion(self.Tree.Table)
	m.Tree.Wheres = self.Tree.Wheres
	m.Tree.Limit = self.Tree.Limit
	m.Adapter = self.Adapter
	return m
}

// ToSql calls a visitor's Accept method based on the manager's SQL adapter.
func (self *UpdateManager) ToSql() (string, []interface{}, error) {
	return VisitorFor(self.Adapter).Accept(self.Tree)
}

// UpdateManager factory method.
func Modification(relation *TableNode) (m *UpdateManager) {
	m = new(UpdateManager)
	m.Tree = UpdateStatement(relation)
	m.Adapter = relation.Adapter

	return
}
