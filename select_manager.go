package codex

// SelectManager manages a tree that compiles to a SQL select statement.
type SelectManager struct {
	Tree    *SelectStatementNode // The AST for the SQL SELECT statement.
	Adapter adapter              // The SQL adapter.
}

var _ Scoper = (*SelectManager)(nil)

func (self *SelectManager) Scopes(scopes ...ScopeFunc) *SelectManager {
	for _, scope := range scopes {
		scope(self)
	}
	return self
}

func (self *SelectManager) Scope(expr interface{}, args ...interface{}) {
	self.Where(expr, args...)
}

// // Clone returns
// func (m *SelectManager) Clone() *SelectManager{
// 	newMng := deepcopy.Copy(m)
// 	return  newMng.(*SelectManager)
// }

// Appends a projection to the current Context's Cols slice,
// typically an AttributeNode or string.  If a string is provided, it is
// inserted as a LiteralNode.
func (self *SelectManager) Select(projections ...interface{}) *SelectManager {
	for _, projection := range projections {
		if _, ok := projection.(string); ok {
			projection = UnqualifiedColumn(projection)
		}

		self.Tree.Cols = append(self.Tree.Cols, projection)
	}

	return self
}

// Where Appends an expression to the current Context's Wheres slice,
// typically a comparison, i.e. 1 = 1
//
//   Where("a")                             // no   args -> Group(Literal("a"))
//   Where("a = ?", 123)                    // with args -> Group(Literal("a = ?", 123))
//   Where("a = ? AND b = ?", 123, true)    // with args -> Group(Literal("a = ? AND b = ?", 123, true))
//   Where(Equal(Column("a"), Column("b"))) // no   args -> Group(Equal(Column("a"), Column("b")))
//
//
func (self *SelectManager) Where(expr interface{}, args ...interface{}) *SelectManager {

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

// Sets the Tree's Offset to the given integer.
func (self *SelectManager) Offset(skip int) *SelectManager {
	self.Tree.Offset = Offset(skip)
	return self
}

// Sets the Tree's Limit to the given integer.
func (self *SelectManager) Limit(take int) *SelectManager {
	self.Tree.Limit = Limit(take)
	return self
}

// Appends a new InnerJoin to the current Context's SourceNode.
func (self *SelectManager) InnerJoin(table interface{}) *SelectManager {
	switch table.(type) {
	case Accessor:
		self.Tree.Source.Right = append(self.Tree.Source.Right, InnerJoin(table.(Accessor).Relation(), nil))
	case *RelationNode:
		self.Tree.Source.Right = append(self.Tree.Source.Right, InnerJoin(table.(*RelationNode), nil))
	}

	return self
}

// Appends a new InnerJoin to the current Context's SourceNode.
func (self *SelectManager) OuterJoin(table interface{}) *SelectManager {
	switch table.(type) {
	case Accessor:
		self.Tree.Source.Right = append(self.Tree.Source.Right, OuterJoin(table.(Accessor).Relation(), nil))
	case *RelationNode:
		self.Tree.Source.Right = append(self.Tree.Source.Right, OuterJoin(table.(*RelationNode), nil))
	}

	return self
}

// Sets the last stored Join's Right leaf to a OnNode containing the
// given expression.
func (self *SelectManager) On(expr interface{}) *SelectManager {
	joins := self.Tree.Source.Right

	if 0 == len(joins) {
		return self
	}

	last := joins[len(joins)-1]

	switch last.(type) {
	case *InnerJoinNode:
		last.(*InnerJoinNode).Right = On(expr)
	case *OuterJoinNode:
		last.(*OuterJoinNode).Right = On(expr)
	}

	return self
}

// Appends an expression to the current Context's Orders slice,
// typically an attribute.
func (self *SelectManager) Order(expr interface{}) *SelectManager {
	if str, ok := expr.(string); ok {
		expr = Literal(str)
	}

	self.Tree.Orders = append(self.Tree.Orders, expr)
	return self
}

// Appends a node to the current Context's Groups slice,
// typically an attribute or column.
func (self *SelectManager) Group(groupings ...interface{}) *SelectManager {
	for _, group := range groupings {
		if str, ok := group.(string); ok {
			group = Literal(str)
		}

		self.Tree.Groups = append(self.Tree.Groups, group)
	}
	return self
}

// Sets the Tree's Having member to the given expression.
func (self *SelectManager) Having(expr interface{}) *SelectManager {
	if str, ok := expr.(string); ok {
		expr = Literal(str)
	}

	self.Tree.Having = Having(expr)
	return self
}

// Count returns a pointer to an new SelectManager, while keeping Wheres, Havings...
func (self *SelectManager) Count(expr interface{}) *SelectManager {
	if str, ok := expr.(string); ok {
		expr = UnqualifiedColumn(str)
	}

	cols := make([]interface{}, 1)
	cols[0] = Count(expr)

	tree := &SelectStatementNode{
		Relation:   self.Tree.Relation,
		Source:     self.Tree.Source,
		Cols:       cols,
		Wheres:     self.Tree.Wheres,
		Groups:     self.Tree.Groups,
		Having:     self.Tree.Having,
		Orders:     make([]interface{}, 0),
		Combinator: self.Tree.Combinator,
		Limit:      self.Tree.Limit,
		Offset:     self.Tree.Offset,
	}

	m := &SelectManager{
		Tree:    tree,
		Adapter: self.Adapter,
	}

	return m
}

// Union sets the SelectManager's Tree's Combination member to a
// UnionNode of itself and the parameter `manager`'s Tree.
func (self *SelectManager) Union(manager *SelectManager) *SelectManager {
	self.Tree.Combinator = Union(self.Tree, manager.Tree)
	return self
}

// Intersect sets the SelectManager's Tree's Combination member to a
// IntersectNode of itself and the parameter `manager`'s Tree.
func (self *SelectManager) Intersect(manager *SelectManager) *SelectManager {
	self.Tree.Combinator = Intersect(self.Tree, manager.Tree)
	return self
}

// Except sets the SelectManager's Tree's Combination member to a
// ExceptNode of itself and the parameter `manager`'s Tree.
func (self *SelectManager) Except(manager *SelectManager) *SelectManager {
	self.Tree.Combinator = Except(self.Tree, manager.Tree)
	return self
}

// ToSql calls a visitor's Accept method based on the manager's SQL adapter.
func (self *SelectManager) ToSql() (string, []interface{}, error) {
	// for _, core := range self.Tree.Cores {
	if 0 == len(self.Tree.Cols) {
		self.Tree.Cols = append(self.Tree.Cols, Attribute(Star(), self.Tree.Relation))
	}
	// }

	return VisitorFor(self.Adapter).Accept(self.Tree)
}

// SelectManager factory method.
func Selection(relation *RelationNode) (m *SelectManager) {
	m = new(SelectManager)
	m.Tree = SelectStatement(relation)
	m.Adapter = relation.Adapter
	return
}
