package codex

// SelectManager manages a tree that compiles to a SQL select statement.
type SelectManager struct {
	Tree    *SelectStatementNode // The AST for the SQL SELECT statement.
	Context *SelectCoreNode      // Reference to the Core the manager is curretly operating on.
	adapter interface{}          // The SQL adapter.
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

		self.Context.Cols = append(self.Context.Cols, projection)
	}

	return self
}

// Where Appends an expression to the current Context's Wheres slice,
// typically a comparison, i.e. 1 = 1
//
//   Where("a")                             // no   args -> Grouping(Literal("a"))
//   Where("a = ?", 123)                    // with args -> Grouping(Literal("a = ?", 123))
//   Where("a = ? AND b = ?", 123, true)    // with args -> Grouping(Literal("a = ? AND b = ?", 123, true))
//   Where(Equal(Column("a"), Column("b"))) // no   args -> Grouping(Equal(Column("a"), Column("b")))
//
//
func (self *SelectManager) Where(expr interface{}, args ...interface{}) *SelectManager {

	if str, ok := expr.(string); ok {
		if len(args) == 0 {
			expr = Literal(str)
		} else {
			expr = Literal(str, args...)
		}
	}
	// enclose expr in Grouping - except if expr is already a Grouping
	if _, ok := expr.(*GroupingNode); !ok {
		expr = Grouping(expr)
	}

	self.Context.Wheres = append(self.Context.Wheres, expr)
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
		self.Context.Source.Right = append(self.Context.Source.Right, InnerJoin(table.(Accessor).Relation(), nil))
	case *RelationNode:
		self.Context.Source.Right = append(self.Context.Source.Right, InnerJoin(table.(*RelationNode), nil))
	}

	return self
}

// Appends a new InnerJoin to the current Context's SourceNode.
func (self *SelectManager) OuterJoin(table interface{}) *SelectManager {
	switch table.(type) {
	case Accessor:
		self.Context.Source.Right = append(self.Context.Source.Right, OuterJoin(table.(Accessor).Relation(), nil))
	case *RelationNode:
		self.Context.Source.Right = append(self.Context.Source.Right, OuterJoin(table.(*RelationNode), nil))
	}

	return self
}

// Sets the last stored Join's Right leaf to a OnNode containing the
// given expression.
func (self *SelectManager) On(expr interface{}) *SelectManager {
	joins := self.Context.Source.Right

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

		self.Context.Groups = append(self.Context.Groups, group)
	}
	return self
}

// Sets the Context's Having member to the given expression.
func (self *SelectManager) Having(expr interface{}) *SelectManager {
	if str, ok := expr.(string); ok {
		expr = Literal(str)
	}

	self.Context.Having = Having(expr)
	return self
}

// Count returns a pointer to an new SelectManager, while keeping Wheres, Havings...
func (self *SelectManager) Count(expr interface{}) *SelectManager {
	if str, ok := expr.(string); ok {
		expr = UnqualifiedColumn(str)
	}

	cols := make([]interface{}, 1)
	cols[0] = Count(expr)

	ctx := &SelectCoreNode{
		Relation: self.Context.Relation,
		Source:   self.Context.Source,
		Cols:     cols,
		Wheres:   self.Context.Wheres,
		Groups:   self.Context.Groups,
		Having:   self.Context.Having,
	}

	tree := &SelectStatementNode{
		Cores:      make([]*SelectCoreNode, len(self.Tree.Cores)),
		Orders:     make([]interface{}, 0),
		Combinator: self.Tree.Combinator,
		Limit:      self.Tree.Limit,
		Offset:     self.Tree.Offset,
	}

	for i, c := range self.Tree.Cores {
		core := &SelectCoreNode{

			Relation: c.Relation,
			Source:   c.Source,
			Cols:     cols,
			Wheres:   c.Wheres,
			Groups:   c.Groups,
			Having:   c.Having,
		}
		tree.Cores[i] = core
	}

	m := &SelectManager{
		Tree:    tree,
		Context: ctx,
		adapter: self.adapter,
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

// Sets the SQL Adapter.
func (self *SelectManager) SetAdapter(adapter interface{}) *SelectManager {
	self.adapter = adapter
	return self
}

// ToSql calls a visitor's Accept method based on the manager's SQL adapter.
func (self *SelectManager) ToSql() (string, []interface{}, error) {
	for _, core := range self.Tree.Cores {
		if 0 == len(core.Cols) {
			core.Cols = append(core.Cols, Attribute(Star(), core.Relation))
		}
	}

	if nil == self.adapter {
		self.adapter = "to_sql"
	}

	return VisitorFor(self.adapter).Accept(self.Tree)
}

// SelectManager factory method.
func Selection(relation *RelationNode) (selection *SelectManager) {
	selection = new(SelectManager)
	selection.Tree = SelectStatement(relation)
	selection.Context = selection.Tree.Cores[0]
	return
}
