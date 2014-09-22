// Package codex provides a Relational Algebra.
package codex

import ()

// ToggleDebugMode toggles debugger variable for managers package.
func ToggleDebugMode() {
	DEBUG = !DEBUG
}

// Table returns an Accessor from the managers package for
// generating SQL to interact with existing tables.
func Table(name string) Accessor {
	relation := Relation(name)
	return func(name interface{}) *AttributeNode {
		if _, ok := name.(string); ok {
			return Attribute(Column(name), relation)
		}

		return Attribute(name, relation)
	}
}

// // CreateTable returns an AlterManager from the managers package
// // for generating SQL to create new tables.
// func CreateTable(name string) *CreateManager {
// 	relation := Relation(name)
// 	return Creation(relation)
// }

// // CreateTable returns an AlterManager from the managers package
// // for generating SQL to alter existing tables.
// func AlterTable(name string) *AlterManager {
// 	relation := Relation(name)
// 	return Alteration(relation)
// }
