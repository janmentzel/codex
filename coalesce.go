package codex

// CoalesceNode is a FunctionNode struct.
type CoalesceNode FunctionNode

// Returns and Equal node containing a reference to the
// function and other
func (self *CoalesceNode) Eq(other interface{}) *EqualNode {
	return Equal(self, other)
}

// Returns and NotEqual node containing a reference to the
// function and other
func (self *CoalesceNode) Neq(other interface{}) *NotEqualNode {
	return NotEqual(self, other)
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (self *CoalesceNode) Gt(other interface{}) *GreaterThanNode {
	return GreaterThan(self, other)
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (self *CoalesceNode) Gte(other interface{}) *GreaterThanOrEqualNode {
	return GreaterThanOrEqual(self, other)
}

// Returns and LessThan node containing a reference to the
// function and other
func (self *CoalesceNode) Lt(other interface{}) *LessThanNode {
	return LessThan(self, other)
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (self *CoalesceNode) Lte(other interface{}) *LessThanOrEqualNode {
	return LessThanOrEqual(self, other)
}

// Returns and Like node containing a reference to the
// function and other
func (self *CoalesceNode) Like(other interface{}) *LikeNode {
	return Like(self, other)
}

// Returns and Unlike node containing a reference to the
// function and other
func (self *CoalesceNode) Unlike(other interface{}) *UnlikeNode {
	return Unlike(self, other)
}

// Returns and Or node containing a reference to the
// function and other
func (self *CoalesceNode) Or(other interface{}) *GroupingNode {
	return Grouping(Or(self, other))
}

// Returns and And node containing a reference to the
// function and other
func (self *CoalesceNode) And(other interface{}) *GroupingNode {
	return Grouping(And(self, other))
}

// As creates an alias e.g. vor selected cols e.g. COALESCE("products"."id") AS "product_id"
// "product_id" would be the alias here
func (self *CoalesceNode) As(alias interface{}) *AsNode {
	if s, ok := alias.(string); ok {
		alias = Column(s)
	}
	return As(self, alias)
}

// CoalesceNode factory method.
func Coalesce(expr ...interface{}) (coal *CoalesceNode) {
	coal = new(CoalesceNode)
	coal.Expressions = expr
	return
}
