package codex

// EqualNode is a BinaryNode struct
type InNode BinaryNode

// Returns a Grouping node with an expression containing a
// reference to an Or node of the Equal and other.
func (self *InNode) Or(other interface{}) *GroupingNode {
	return Grouping(Or(self, other))
}

// Returns a Grouping node with an expression containing a
// reference to an And node of the Equal and other.
func (self *InNode) And(other interface{}) *GroupingNode {
	return Grouping(And(self, other))
}

// Returns an Not node with and expression containing the
// Equal node.
func (self *InNode) Not() *NotNode {
	return Not(self)
}

// In factory method.
func In(left interface{}, args ...interface{}) *InNode {
	return &InNode{
		Left:  left,
		Right: args,
	}
}
