package codex

// UnaryNode struct.
type UnaryNode struct {
	Expr interface{} // Single leaf for the Unary node.
}

// BindingNode is deprecated!
type BindingNode UnaryNode // BindingNode is a UnaryNode struct.

type OnNode UnaryNode     // OnNode is a UnaryNode struct.
type LimitNode UnaryNode  // LimitNode is a UnaryNode struct.
type OffsetNode UnaryNode // OffsetNode is a UnaryNode struct.
type HavingNode UnaryNode // HavingNode is a UnaryNode struct.
type ColumnNode UnaryNode // ColumnNode is a UnaryNode struct.
type StarNode UnaryNode   // StarNode is a Unary node struct.

// OnNode factory method.
func On(expr interface{}) (on *OnNode) {
	on = new(OnNode)
	on.Expr = expr
	return
}

// LimitNode factory method.
func Limit(expr interface{}) (limit *LimitNode) {
	limit = new(LimitNode)
	limit.Expr = expr
	return
}

// OffsetNode factory method.
func Offset(expr interface{}) (offset *OffsetNode) {
	offset = new(OffsetNode)
	offset.Expr = expr
	return
}

// HavingNode factory method.
func Having(expr interface{}) (having *HavingNode) {
	having = new(HavingNode)
	having.Expr = expr
	return
}

// ColumnNode factory method.
func Column(expr interface{}) (column *ColumnNode) {
	column = new(ColumnNode)
	column.Expr = expr
	return
}

// StarNode factory method.
func Star() *StarNode {
	return new(StarNode)
}

// deprecated
// BindingNode factory method.
func Binding() *BindingNode {
	return new(BindingNode)
}
