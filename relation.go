package codex

// RelationNode is a specific BinaryNode
type RelationNode struct {
	Name    string  // Relation's Name
	Alias   *string // Relation's Alias
	Adapter adapter
}

// RelationNode factory method.
func Relation(name string) (relation *RelationNode) {
	relation = new(RelationNode)
	relation.Name = name
	return
}

// Col returns a Column scoped to this table
func (r *RelationNode) Col(name string) *AttributeNode {
	return Attribute(Column(name), r)
}

// Select returns a SelectManager
// appends the columns
func (r *RelationNode) Select(cols ...interface{}) *SelectManager {
	// convert string to AttributeNode
	for i, col := range cols {
		if str, ok := col.(string); ok {
			cols[i] = r.Col(str)
		}
	}

	return Selection(r).Select(cols...)
}

// Where Returns a pointer to a SelectManager with the initial filter provided.
// see SelectManager.Where()
func (r *RelationNode) Where(expr interface{}, args ...interface{}) *SelectManager {
	return Selection(r).Where(expr, args...)
}

// Returns a pointer to a SelectManager with an initial InnerJoinNode.
func (r *RelationNode) InnerJoin(expr interface{}) *SelectManager {
	return Selection(r).InnerJoin(expr)
}

// Returns a pointer to a SelectManager with an initial OuterJoinNode.
func (r *RelationNode) OuterJoin(expr interface{}) *SelectManager {
	return Selection(r).OuterJoin(expr)
}

// Returns a pointer to a SelectManager with an initial Ordering.
func (r *RelationNode) Order(expr interface{}) *SelectManager {
	return Selection(r).Order(expr)
}

// Returns a pointer to a SelectManager with an initial Grouping.
func (r *RelationNode) Group(groupings ...interface{}) *SelectManager {
	return Selection(r).Group(groupings...)
}

// Returns a pointer to a SelectManager with an initial Having.
func (r *RelationNode) Having(expr interface{}) *SelectManager {
	return Selection(r).Having(expr)
}

// Count Returns a pointer to a new SelectManager
func (r *RelationNode) Count(expr interface{}) *SelectManager {
	return Selection(r).Count(expr)
}

// // Returns a pointer to a SelectManager with for the given RelationNode.
// func (r *RelationNode) From(relation *RelationNode) *SelectManager {
// 	return Selection(relation)
// }

// Returns a pointer to a InsertManager with initial the values provided.
func (r *RelationNode) Insert(expr ...interface{}) *InsertManager {
	return Insertion(r).Insert(expr...)
}

// Returns a pointer to a UpdateManager with initial the columns provided.
func (r *RelationNode) Set(expr ...interface{}) *UpdateManager {
	return Modification(r).Set(expr...)
}

// Returns a pointer to a DeleteManager with the initial filter provided.
func (r *RelationNode) Delete(expr interface{}) *DeleteManager {
	return Deletion(r).Delete(expr)
}
