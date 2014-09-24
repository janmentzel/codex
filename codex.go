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

type dbDialect func(string) *AttributeNode

func (db dbDialect) Table(name string) *RelationNode {
	return db(name).Relation
}

func Dialect(adapter adapter) dbDialect {
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
