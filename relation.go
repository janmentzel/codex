// Package nodes provides nodes to use in codex AST's.
package codex

// RelationNode is a specific BinaryNode
type RelationNode struct {
	Name  string  // Relation's Name
	Alias *string // Relation's Alias
}

func (r *RelationNode) Col(name string) *AttributeNode {
	return Attribute(Column(name), r)
}

func (t *RelationNode) Select(cols ...interface{}) *SelectManager {
	// convert string to column object
	for i, col := range cols {
		if str, ok := col.(string); ok {
			cols[i] = t.Col(str)
		}
	}

	return Selection(t).Project(cols...)
}

func (t *RelationNode) Where(expr interface{}) *SelectManager {
	return Selection(t).Where(expr)
}

// TODO more convenience like select, where, group, join, offset, having, ...

// RelationNode factory method.
func Relation(name string) (relation *RelationNode) {
	relation = new(RelationNode)
	relation.Name = name
	return
}
