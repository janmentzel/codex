// Package managers provides AST managers for the codex package.
package codex

// import ()

// type CreateManager struct {
// 	Tree    *CreateStatementNode
// 	adapter interface{} // The SQL adapter.
// }

// // AddColumn adds a UnexistingColumn from the nodes package to the AST for creation.
// func (self *CreateManager) AddColumn(name interface{}, typ Type) *CreateManager {
// 	if _, ok := name.(string); ok {
// 		name = UnqualifiedColumn(name)
// 	}

// 	self.Tree.UnexistingColumns = append(self.Tree.UnexistingColumns, UnexistingColumn(name, typ))
// 	return self
// }

// // AddColumn adds a ConstraintNode from the nodes package to the AST to apply to a column.
// func (self *CreateManager) AddConstraint(columns []interface{}, kind T_Constraint, options ...interface{}) *CreateManager {
// 	for index, column := range columns {
// 		if _, ok := column.(string); ok {
// 			columns[index] = UnqualifiedColumn(column)
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

// func (self *CreateManager) AddUniqueConstraint(columns []interface{},
// 	name ...interface{}) *CreateManager {

// 	return self.AddConstraint(columns, T_Unique, name...)
// }

// func (self *CreateManager) AddForiegnKeyConstraint(columns []interface{},
// 	reference interface{}, name ...interface{}) *CreateManager {

// 	return self.AddConstraint(columns, T_ForeignKey, append([]interface{}{
// 		reference,
// 	}, name...)...)
// }

// func (self *CreateManager) AddDefaultContraint(columns []interface{},
// 	value interface{}) *CreateManager {

// 	return self.AddConstraint(columns, T_Default, value)
// }

// func (self *CreateManager) AddNotNullConstraint(columns []interface{}) *CreateManager {
// 	return self.AddConstraint(columns, T_NotNull)
// }

// func (self *CreateManager) AddPrimaryKeyConstraint(columns []interface{},
// 	name ...interface{}) *CreateManager {

// 	return self.AddConstraint(columns, T_PrimaryKey, name...)
// }

// // SetEngine sets the AST's Engine field, used for table creation.
// func (self *CreateManager) SetEngine(engine interface{}) *CreateManager {
// 	if _, ok := engine.(*EngineNode); !ok {
// 		engine = Engine(engine)
// 	}

// 	self.Tree.Engine = engine.(*EngineNode)
// 	return self
// }

// // Sets the SQL Adapter.
// func (self *CreateManager) SetAdapter(adapter interface{}) *CreateManager {
// 	self.adapter = adapter
// 	return self
// }

// // ToSql calls a visitor's Accept method based on the manager's SQL adapter.
// func (self *CreateManager) ToSql() (string, error) {
// 	if nil == self.adapter {
// 		self.adapter = "to_sql"
// 	}

// 	return VisitorFor(self.adapter).Accept(self.Tree)
// }

// // SelectManager factory method.
// func Creation(relation *RelationNode) (statement *CreateManager) {
// 	statement = new(CreateManager)
// 	statement.Tree = CreateStatement(relation)
// 	return
// }
