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
func On(expr interface{}) *OnNode {
	return &OnNode{expr}
}

// LimitNode factory method.
func Limit(expr interface{}) *LimitNode {
	return &LimitNode{expr}
}

// OffsetNode factory method.
func Offset(expr interface{}) *OffsetNode {
	return &OffsetNode{expr}
}

// HavingNode factory method.
func Having(expr interface{}) *HavingNode {
	return &HavingNode{expr}
}

// ColumnNode factory method.
func Column(expr interface{}) *ColumnNode {
	return &ColumnNode{expr}
}

// StarNode factory method.
func Star() *StarNode {
	return &StarNode{}
}

// deprecated
// BindingNode factory method.
func Binding() *BindingNode {
	return &BindingNode{}
}
