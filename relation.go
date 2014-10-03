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
func (r *TableNode) Col(name string) *AttributeNode {
	return Attribute(Column(name), r)
}

// Select returns a SelectManager
// appends the columns
func (r *TableNode) Select(cols ...interface{}) *SelectManager {
	// convert string to AttributeNode
	for i, col := range cols {
		if str, ok := col.(string); ok {
			cols[i] = r.Col(str)
		}
	}

	return Selection(r).Scopes(r.scopes...).Select(cols...)
}

// Where Returns a pointer to a SelectManager with the initial filter provided.
// see SelectManager.Where()
func (r *TableNode) Where(expr interface{}, args ...interface{}) *SelectManager {
	return Selection(r).Scopes(r.scopes...).Where(expr, args...)
}

// Returns a pointer to a SelectManager with an initial InnerJoinNode.
func (r *TableNode) InnerJoin(expr interface{}) *SelectManager {
	return Selection(r).Scopes(r.scopes...).InnerJoin(expr)
}

// Returns a pointer to a SelectManager with an initial OuterJoinNode.
func (r *TableNode) OuterJoin(expr interface{}) *SelectManager {
	return Selection(r).Scopes(r.scopes...).OuterJoin(expr)
}

// Returns a pointer to a SelectManager with an initial Ordering.
func (r *TableNode) Order(expr interface{}) *SelectManager {
	return Selection(r).Scopes(r.scopes...).Order(expr)
}

// Returns a pointer to a SelectManager with an initial Grouping.
func (r *TableNode) Group(groupings ...interface{}) *SelectManager {
	return Selection(r).Scopes(r.scopes...).Group(groupings...)
}

// Returns a pointer to a SelectManager with an initial Having.
func (r *TableNode) Having(expr interface{}) *SelectManager {
	return Selection(r).Scopes(r.scopes...).Having(expr)
}

// Count Returns a pointer to a new SelectManager
func (r *TableNode) Count(expr interface{}) *SelectManager {
	return Selection(r).Scopes(r.scopes...).Count(expr)
}

// // Returns a pointer to a SelectManager with for the given TableNode.
// func (r *TableNode) From(relation *TableNode) *SelectManager {
// 	return Selection(relation)
// }

// Returns a pointer to a InsertManager with initial the values provided.
func (r *TableNode) Insert(expr ...interface{}) *InsertManager {
	return Insertion(r).Insert(expr...)
}

// Returns a pointer to a UpdateManager with initial the columns provided.
func (r *TableNode) Set(expr ...interface{}) *UpdateManager {
	return Modification(r).Scopes(r.scopes...).Set(expr...)
}

// Returns a pointer to a DeleteManager with the initial filter provided.
func (r *TableNode) Delete(expr interface{}) *DeleteManager {
	return Deletion(r).Scopes(r.scopes...).Delete(expr)
}
