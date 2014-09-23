package codex

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
