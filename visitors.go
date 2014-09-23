package codex

// VisitorFor returns a AST visitor for the adapter argument.
func VisitorFor(adapter adapter) VisitorInterface {
	switch adapter {
	case MYSQL:
		return NewMySqlVisitor()
	case POSTGRES:
		return NewPostgresVisitor()
	default:
		return NewToSqlVisitor()
	}
}
