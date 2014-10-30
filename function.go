package codex

// FunctionNode is a Nary node.
type FunctionNode struct {
	Name string        // MAX, COUNT, LOWER, AVG, ...
	Args []interface{} //
	// TODO remove
	Alias interface{} // Alias the function result is reffered to as.
	// TODO make Distinkt a Node without (); strange btw.
	Distinct bool // Function is distinct.
}

// Returns and Equal node containing a reference to the
// function and other
func (f *FunctionNode) Eq(other interface{}) *EqualNode {
	return Equal(f, other)
}

// Returns and NotEqual node containing a reference to the
// function and other
func (f *FunctionNode) Neq(other interface{}) *NotEqualNode {
	return NotEqual(f, other)
}

// Returns and GreaterThan node containing a reference to the
// function and other
func (f *FunctionNode) Gt(other interface{}) *GreaterThanNode {
	return GreaterThan(f, other)
}

// Returns and GreaterThanOrEqual node containing a reference to the
// function and other
func (f *FunctionNode) Gte(other interface{}) *GreaterThanOrEqualNode {
	return GreaterThanOrEqual(f, other)
}

// Returns and LessThan node containing a reference to the
// function and other
func (f *FunctionNode) Lt(other interface{}) *LessThanNode {
	return LessThan(f, other)
}

// Returns and LessThanOrEqual node containing a reference to the
// function and other
func (f *FunctionNode) Lte(other interface{}) *LessThanOrEqualNode {
	return LessThanOrEqual(f, other)
}

// Returns and Like node containing a reference to the
// function and other
func (f *FunctionNode) Like(other interface{}) *LikeNode {
	return Like(f, other)
}

// Returns and Unlike node containing a reference to the
// function and other
func (f *FunctionNode) Unlike(other interface{}) *UnlikeNode {
	return Unlike(f, other)
}

// Returns and Or node containing a reference to the
// function and other
func (f *FunctionNode) Or(other interface{}) *GroupingNode {
	return Grouping(Or(f, other))
}

// Returns and And node containing a reference to the
// function and other
func (f *FunctionNode) And(other interface{}) *GroupingNode {
	return Grouping(And(f, other))
}

// As creates an alias e.g. vor selected cols e.g. "products"."id" AS "product_id"
// "product_id" would be the alias here
func (f *FunctionNode) As(alias interface{}) *AsNode {
	if s, ok := alias.(string); ok {
		alias = Column(s)
	}
	return As(f, alias)
}

// FunctionNode generic factory method.
func Function(name string, args ...interface{}) *FunctionNode {
	return &FunctionNode{
		Name: name,
		Args: args,
	}
}

func Avg(args ...interface{}) *FunctionNode {
	return Function("AVG", args...)
}

func Coalesce(args ...interface{}) *FunctionNode {
	return Function("COALESCE", args...)
}

func Count(args ...interface{}) *FunctionNode {
	return Function("COUNT", args...)
}

func Max(args ...interface{}) *FunctionNode {
	return Function("MAX", args...)
}

func Min(args ...interface{}) *FunctionNode {
	return Function("MIN", args...)
}

func Sum(args ...interface{}) *FunctionNode {
	return Function("SUM", args...)
}

func Lower(args ...interface{}) *FunctionNode {
	return Function("LOWER", args...)
}

func Upper(args ...interface{}) *FunctionNode {
	return Function("UPPER", args...)
}

func Substring(args ...interface{}) *FunctionNode {
	return Function("SUBSTRING", args...)
}
