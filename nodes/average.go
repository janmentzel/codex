// Package nodes provides nodes to use in codex AST's.
package nodes

// AverageNode is a FunctionNode struct.
type AverageNode FunctionNode

// Returns and Equal node containing a reference to the
// function and other
func (self *AverageNode) Eq(other interface{}) *EqualNode {
	return Equal(self, other)
}

// Returns and NotEqual node containing a reference to the
// function and other
func (self *AverageNode) Neq(other interface{}) *NotEqualNode {
	return NotEqual(self, other)
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (self *AverageNode) Gt(other interface{}) *GreaterThanNode {
	return GreaterThan(self, other)
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (self *AverageNode) Gte(other interface{}) *GreaterThanOrEqualNode {
	return GreaterThanOrEqual(self, other)
}

// Returns and LessThan node containing a reference to the
// function and other
func (self *AverageNode) Lt(other interface{}) *LessThanNode {
	return LessThan(self, other)
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (self *AverageNode) Lte(other interface{}) *LessThanOrEqualNode {
	return LessThanOrEqual(self, other)
}

// Returns and Like node containing a reference to the
// function and other
func (self *AverageNode) Like(other interface{}) *LikeNode {
	return Like(self, other)
}

// Returns and Unlike node containing a reference to the
// function and other
func (self *AverageNode) Unlike(other interface{}) *UnlikeNode {
	return Unlike(self, other)
}

// Returns and Or node containing a reference to the
// function and other
func (self *AverageNode) Or(other interface{}) *GroupingNode {
	return Grouping(Or(self, other))
}

// Returns and And node containing a reference to the
// function and other
func (self *AverageNode) And(other interface{}) *GroupingNode {
	return Grouping(And(self, other))
}

// AverageNode factory method.
func Average(expressions ...interface{}) (average *AverageNode) {
	average = new(AverageNode)
	average.Expressions = expressions
	return
}
