package codex

// LiteralNode has raw SQL string and args array
type LiteralNode struct {
	Sql  string
	Args []interface{}
}

// LiteralNode factory method.
func Literal(sql string, args ...interface{}) *LiteralNode {
	return &LiteralNode{sql, args}
}

// TODO klimbim methoden

// count
// sum
// max
// min
// avg
// extract(field)
// equal
// notequal
// ...
// matches
// ...
