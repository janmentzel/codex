// Package codex provides a Relational Algebra for PostgreSQL and MySQL. Based on Arel (Ruby on Rails).
package codex

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
