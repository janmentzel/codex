package codex

import (
	"strings"
)

// LiteralNode has raw SQL string and args array
type LiteralNode struct {
	Sql  string
	Args []interface{}
}

// LiteralNode factory method.
// "?..." in sql and args = ["a","b","c"]  expands to "?,?,?" in sql
// only one "?..." per Literal supported atm.
// something like that does NOT work:  Literal("foo IN(?...) OR bar IN(?...)")
func Literal(sql string, args ...interface{}) *LiteralNode {

	if strings.Contains(sql, "?...") {
		parts := strings.Split(sql, "?...")
		if len(parts) > 2 {
			sql = parts[0] + "?..." + parts[1] + "-- ERROR supports ?... only once per Literal --"
		} else {
			nargs := len(args)
			if nargs > 0 {
				qs := make([]rune, nargs*2-1)

				for i := 0; i < len(qs); i++ {
					if i > 0 {
						qs[i] = ','
						i++
					}
					qs[i] = '?'
				}
				sql = parts[0] + string(qs) + parts[1]
			} else {
				sql = parts[0] + parts[1]
			}
		}
	}

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
