package codex

// TableNode is a specific BinaryNode
type TableNode struct {
	Name    string  // Table's Name
	Alias   *string // Table's Alias
	Adapter adapter
	scopes  []ScopeFunc
}

func (self *TableNode) Scopes(scopes ...ScopeFunc) *TableNode {
	for _, scope := range scopes {
		self.scopes = append(self.scopes, scope)
	}
	return self
}

// TableNode factory method.
func Table(name string) (relation *TableNode) {
	relation = new(TableNode)
	relation.Name = name
	return
}

// Col returns a Column scoped to this table
func (t *TableNode) Col(name string) *AttributeNode {
	return Attribute(Column(name), t)
}

// Star returns a * scoped to this table e.g. renders to '"tablename".*' sql
func (t *TableNode) Star() *AttributeNode {
	return Attribute(Star(), t)
}

// Select returns a SelectManager
// appends the columns
func (t *TableNode) Select(cols ...interface{}) *SelectManager {
	// convert string to AttributeNode
	for i, col := range cols {
		if str, ok := col.(string); ok {
			cols[i] = t.Col(str)
		}
	}

	return Selection(t).Scopes(t.scopes...).Select(cols...)
}

// Where Returns a pointer to a SelectManager with the initial filter provided.
// see SelectManager.Where()
func (t *TableNode) Where(expr interface{}, args ...interface{}) *SelectManager {
	return Selection(t).Scopes(t.scopes...).Where(expr, args...)
}

// Returns a pointer to a SelectManager with an initial InnerJoinNode.
func (t *TableNode) InnerJoin(expr interface{}) *SelectManager {
	return Selection(t).Scopes(t.scopes...).InnerJoin(expr)
}

// Returns a pointer to a SelectManager with an initial OuterJoinNode.
func (t *TableNode) OuterJoin(expr interface{}) *SelectManager {
	return Selection(t).Scopes(t.scopes...).OuterJoin(expr)
}

// Returns a pointer to a SelectManager with an initial Ordering.
func (t *TableNode) Order(expr interface{}) *SelectManager {
	return Selection(t).Scopes(t.scopes...).Order(expr)
}

// Returns a pointer to a SelectManager with an initial Grouping.
func (t *TableNode) Group(groupings ...interface{}) *SelectManager {
	return Selection(t).Scopes(t.scopes...).Group(groupings...)
}

// Returns a pointer to a SelectManager with an initial Having.
func (t *TableNode) Having(expr interface{}) *SelectManager {
	return Selection(t).Scopes(t.scopes...).Having(expr)
}

// Count Returns a pointer to a new SelectManager
func (t *TableNode) Count(expr interface{}) *SelectManager {
	return Selection(t).Scopes(t.scopes...).Count(expr)
}

// // Returns a pointer to a SelectManager with for the given TableNode.
// func (t *TableNode) From(relation *TableNode) *SelectManager {
// 	return Selection(relation)
// }

// Returns a pointer to a InsertManager with initial the values provided.
func (t *TableNode) Insert(expr ...interface{}) *InsertManager {
	return Insertion(t).Insert(expr...)
}

// Returns a pointer to a UpdateManager with initial the columns provided.
func (t *TableNode) Set(expr ...interface{}) *UpdateManager {
	return Modification(t).Scopes(t.scopes...).Set(expr...)
}

// Returns a pointer to a DeleteManager with the initial filter provided.
func (t *TableNode) Delete(expr interface{}) *DeleteManager {
	return Deletion(t).Scopes(t.scopes...).Delete(expr)
}

// Returns a pointer to a InsertManager initialized with Table
func (t *TableNode) Insertion() *InsertManager {
	return Insertion(t)
}

// Returns a pointer to a UpdateManager initialized with Table
func (t *TableNode) Modification() *UpdateManager {
	return Modification(t)
}

// Returns a pointer to a SelectManager initialized with Table
func (t *TableNode) Selection() *SelectManager {
	return Selection(t)
}

// Returns a pointer to a DeleteManager initialized with Table
func (t *TableNode) Deletion() *DeleteManager {
	return Deletion(t)
}
