package codex

// Accessor is a type def of a function that returns an AttributeNode
type Accessor func(interface{}) *AttributeNode

// Returns the RelationNode scoped to the Accessor
func (a Accessor) Relation() *RelationNode {
	return a("").Relation
}

// Returns the name of the RelationNode scoped to the Accessor
func (a Accessor) Name() string {
	return a("").Relation.Name
}

// Returns a pointer to a SelectManager with the initial projections provided.
func (a Accessor) Select(projections ...interface{}) *SelectManager {
	return a.From(a.Relation()).Select(projections...)
}

// Returns a pointer to a SelectManager with the initial filter provided.
func (a Accessor) Where(expr interface{}) *SelectManager {
	return a.From(a.Relation()).Where(expr)
}

// Returns a pointer to a SelectManager with an initial InnerJoinNode.
func (a Accessor) InnerJoin(expr interface{}) *SelectManager {
	return a.From(a.Relation()).InnerJoin(expr)
}

// Returns a pointer to a SelectManager with an initial OuterJoinNode.
func (a Accessor) OuterJoin(expr interface{}) *SelectManager {
	return a.From(a.Relation()).OuterJoin(expr)
}

// Returns a pointer to a SelectManager with an initial Ordering.
func (a Accessor) Order(expr interface{}) *SelectManager {
	return a.From(a.Relation()).Order(expr)
}

// Returns a pointer to a SelectManager with an initial Grouping.
func (a Accessor) Group(groupings ...interface{}) *SelectManager {
	return a.From(a.Relation()).Group(groupings...)
}

// Returns a pointer to a SelectManager with an initial Having.
func (a Accessor) Having(expr interface{}) *SelectManager {
	return a.From(a.Relation()).Having(expr)
}

// Returns a pointer to a SelectManager with for the given RelationNode.
func (a Accessor) From(relation *RelationNode) *SelectManager {
	return Selection(relation)
}

// Returns a pointer to a InsertManager with initial the values provided.
func (a Accessor) Insert(expr ...interface{}) *InsertManager {
	return Insertion(a.Relation()).Insert(expr...)
}

// Returns a pointer to a UpdateManager with initial the columns provided.
func (a Accessor) Set(expr ...interface{}) *UpdateManager {
	return Modification(a.Relation()).Set(expr...)
}

// Returns a pointer to a DeleteManager with the initial filter provided.
func (a Accessor) Delete(expr interface{}) *DeleteManager {
	return Deletion(a.Relation()).Delete(expr)
}

// Returns string, error generated by calling ToSql on an SelectManager.
func (a Accessor) ToSql() (string, []interface{}, error) {
	return a.From(a.Relation()).ToSql()
}
