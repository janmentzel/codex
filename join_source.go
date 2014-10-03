package codex

// JoinSourceNode is a specific BinaryNode.
type JoinSourceNode struct {
	Left  *TableNode    // Left child of the JoinSource node, a pointer to a Table.
	Right []interface{} // Right child of the JoinSource node contains joins and their instructions
}

// JoinSourceNode factory method.
func JoinSource(relation *TableNode) (source *JoinSourceNode) {
	source = new(JoinSourceNode)
	source.Left = relation
	source.Right = make([]interface{}, 0)
	return
}
