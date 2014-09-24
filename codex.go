// Package codex provides a Relational Algebra for PostgreSQL and MySQL. Based on Arel (Ruby on Rails).
package codex

type adapter uint8

const (
	MYSQL adapter = iota + 1
	POSTGRES
)

// ToggleDebugMode toggles debugger variable for managers package.
func ToggleDebugMode() {
	DEBUG = !DEBUG
}

type DbDialect func(string) *AttributeNode

func (db DbDialect) Table(name string) *RelationNode {
	return db(name).Relation
}

func Dialect(adapter adapter) DbDialect {
	return func(tableName string) *AttributeNode {
		table := Relation(tableName)
		table.Adapter = adapter

		return Attribute(tableName, table)
	}
}

// deprecated
// Table returns an Accessor from the managers package for
// generating SQL to interact with existing tables.
func Table(tableName string) Accessor {
	table := Relation(tableName)
	return func(colName interface{}) *AttributeNode {
		if _, ok := colName.(string); ok {
			return Attribute(Column(colName), table)
		}

		return Attribute(colName, table)
	}
}

// Scoper enables scoping support for SelectManager, UpdateManager, DeleteManager
type Scoper interface {
	// Scope wraps Where method and omits the return value so Scope() signature remains generic
	Scope(expr interface{}, args ...interface{})
}

// ScopeFunc is implemented by DB layer 'models'
type ScopeFunc func(Scoper)
