package codex

//

// // AlterManager manages a tree that compiles to SQL for create and alter statement.
// type AlterManager struct {
// 	Tree    *AlterStatementNode // The AST for the SQL CREATE/ALTER TABLE statements.
// 	adapter interface{}         // The SQL adapter.
// }

// // AddColumn adds a UnexistingColumn from the nodes package to the AST for creation.
// func (self *AlterManager) AddColumn(name interface{}, typ Type) *AlterManager {
// 	if _, ok := name.(string); ok {
// 		name = UnqualifiedColumn(name)
// 	}

// 	self.Tree.UnexistingColumns = append(self.Tree.UnexistingColumns, UnexistingColumn(name, typ))
// 	return self
// }

// // AddColumn adds a UnexistingColumn from the nodes package to the AST for creation.
// func (self *AlterManager) AlterColumn(name interface{}, typ Type) *AlterManager {
// 	if _, ok := name.(string); ok {
// 		name = UnqualifiedColumn(name)
// 	}

// 	self.Tree.ModifiedColumns = append(self.Tree.ModifiedColumns, ExistingColumn(name, typ))
// 	return self
// }

// // AddColumn adds a ConstraintNode from the nodes package to the AST to apply to a column.
// func (self *AlterManager) AddConstraint(columns []interface{}, kind T_Constraint, options ...interface{}) *AlterManager {
// 	for _, column := range columns {
// 		if _, ok := column.(string); ok {
// 			column = UnqualifiedColumn(column)
// 		}
// 	}

// 	var node interface{}

// 	switch kind {
// 	case T_NotNull:
// 		node = NotNull(columns, options...)
// 	case T_Unique:
// 		node = Unique(columns, options...)
// 	case T_PrimaryKey:
// 		node = PrimaryKey(columns, options...)
// 	case T_ForeignKey:
// 		node = ForeignKey(columns, options...)
// 	case T_Check:
// 		node = Check(columns, options...)
// 	case T_Default:
// 		node = Default(columns, options...)
// 	default:
// 		node = Constraint(columns, options...)
// 	}

// 	self.Tree.Constraints = append(self.Tree.Constraints, node)
// 	return self
// }

// func (self *AlterManager) RemoveColumn(column interface{}) *AlterManager {
// 	if _, ok := column.(string); ok {
// 		column = UnqualifiedColumn(column)
// 	}

// 	self.Tree.RemovedColumns = append(self.Tree.RemovedColumns, column)

// 	return self
// }

// func (self *AlterManager) RemoveIndex(name interface{}) *AlterManager {
// 	if _, ok := name.(string); ok {
// 		name = IndexName(name)
// 	}

// 	self.Tree.RemovedIndicies = append(self.Tree.RemovedIndicies, name)

// 	return self
// }

// // Sets the SQL Adapter.
// func (self *AlterManager) SetAdapter(adapter interface{}) *AlterManager {
// 	self.adapter = adapter
// 	return self
// }

// // ToSql calls a visitor's Accept method based on the manager's SQL adapter.
// func (self *AlterManager) ToSql() (string, error) {
// 	if nil == self.adapter {
// 		self.adapter = "to_sql"
// 	}

// 	return VisitorFor(self.adapter).Accept(self.Tree)
// }

// // SelectManager factory method.
// func Alteration(relation *RelationNode) (statement *AlterManager) {
// 	statement = new(AlterManager)
// 	statement.Tree = AlterStatement(relation)
// 	return
// }
