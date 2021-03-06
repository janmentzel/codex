package codex

// DeleteManager manages a tree that compiles to a SQL delete statement.
type DeleteManager struct {
	Tree    *DeleteStatementNode // The AST for the SQL DELETE statement.
	Adapter adapter              // The SQL adapter.
}

var _ Scoper = (*DeleteManager)(nil)

func (self *DeleteManager) Scopes(scopes ...ScopeFunc) *DeleteManager {
	for _, scope := range scopes {
		scope(self)
	}
	return self
}

func (self *DeleteManager) Scope(expr interface{}, args ...interface{}) {
	self.Where(expr, args...)
}

// Delete appends the expression to the Trees Wheres slice.
// alias Where
func (self *DeleteManager) Delete(expr interface{}, args ...interface{}) *DeleteManager {
	return self.Where(expr, args...)
}

// Where appends an sql WHERE condition to the current tree's Wheres slice,
//
//   Where("a")                             // no   args -> Group(Literal("a"))
//   Where("a = ?", 123)                    // with args -> Group(Literal("a = ?", 123))
//   Where("a = ? AND b = ?", 123, true)    // with args -> Group(Literal("a = ? AND b = ?", 123, true))
//   Where(Equal(Column("a"), Column("b"))) // no   args -> Group(Equal(Column("a"), Column("b")))
func (self *DeleteManager) Where(expr interface{}, args ...interface{}) *DeleteManager {

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

// Limit Sets the Tree's Limit to the given integer.
func (self *DeleteManager) Limit(expr interface{}) *DeleteManager {
	self.Tree.Limit = Limit(expr)
	return self
}

// Selection returns a *SelectManager while keeping
// wheres, limit and adapter
func (self *DeleteManager) Selection() *SelectManager {
	m := Selection(self.Tree.Table)
	m.Tree.Wheres = self.Tree.Wheres
	m.Tree.Limit = self.Tree.Limit
	m.Adapter = self.Adapter
	return m
}

// Modification returns an *UpdateManager while keeping
// wheres, limit and adapter
func (self *DeleteManager) Modification() *UpdateManager {
	m := Modification(self.Tree.Table)
	m.Tree.Wheres = self.Tree.Wheres
	m.Tree.Limit = self.Tree.Limit
	m.Adapter = self.Adapter
	return m
}

// Insertion returns a *InsertManager while keeping
// relation and adapter
func (self *DeleteManager) Insertion() *InsertManager {
	m := Insertion(self.Tree.Table)
	m.Adapter = self.Adapter
	return m
}

// ToSql calls a visitor's Accept method based on the manager's SQL adapter.
func (self *DeleteManager) ToSql() (string, []interface{}, error) {
	return VisitorFor(self.Adapter).Accept(self.Tree)
}

func (self *DeleteManager) Table() *TableNode {
	return self.Tree.Table
}

// DeleteManager factory methods.
func Deletion(relation *TableNode) (m *DeleteManager) {
	m = new(DeleteManager)
	m.Tree = DeleteStatement(relation)
	m.Adapter = relation.Adapter
	return
}
