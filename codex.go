// Package codex provides a Relational Algebra for PostgreSQL and MySQL. Based on Arel (Ruby on Rails).
package codex

import (
	"regexp"
)

type adapter uint8

const (
	MYSQL adapter = iota + 1
	POSTGRES
)

var VALID_COL_NAME_PATTERN *regexp.Regexp
var VALID_TABLE_NAME_PATTERN *regexp.Regexp

func init() {
	var err error
	VALID_COL_NAME_PATTERN, err = regexp.Compile(`(?i)^[a-z][a-z0-9_\$]*$`)
	if err != nil {
		panic(err)
	}
	VALID_TABLE_NAME_PATTERN = VALID_COL_NAME_PATTERN
}

// ToggleDebugMode toggles debugger variable for managers package.
func ToggleDebugMode() {
	DEBUG = !DEBUG
}

type DbDialect func(string) *AttributeNode

func (db DbDialect) Table(name string) *TableNode {
	return db(name).Table
}

func Dialect(adapter adapter) DbDialect {
	return func(tableName string) *AttributeNode {
		table := Table(tableName)
		table.Adapter = adapter

		return Attribute(tableName, table)
	}
}

// // deprecated
// // Table returns an Accessor from the managers package for
// // generating SQL to interact with existing tables.
// func Table(tableName string) Accessor {
// 	table := Table(tableName)
// 	return func(colName interface{}) *AttributeNode {
// 		if _, ok := colName.(string); ok {
// 			return Attribute(Column(colName), table)
// 		}

// 		return Attribute(colName, table)
// 	}
// }

// Scoper enables scoping support for SelectManager, UpdateManager, DeleteManager
type Scoper interface {
	// Scope wraps Where method and omits the return value so Scope() signature remains generic
	Scope(expr interface{}, args ...interface{})
}

// ScopeFunc is implemented by DB layer 'models'
type ScopeFunc func(Scoper)
