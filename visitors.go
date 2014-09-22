// Package managers provides AST managers for the codex package.
package codex

import ()

// VisitorFor returns a AST visitor for the adapter argument.
func VisitorFor(adapter interface{}) VisitorInterface {
	switch adapter {
	case "mysql":
		return NewMySqlVisitor()
	case "postgres":
		return NewPostgresVisitor()
	default:
		return NewToSqlVisitor()
	}
}
