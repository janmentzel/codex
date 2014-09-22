// Package sql provides common SQL constants for the codex package.
package codex

type T_Constraint uint8

// SQL T_Constraint constants.
const (
	T_NotNull T_Constraint = iota
	T_Unique
	T_PrimaryKey
	T_ForeignKey
	T_Check
	T_Default
)

type Type uint8

// SQL Type constants.
const (
	String Type = iota
	Text
	Boolean
	Integer
	Float
	Decimal
	Date
	Time
	Datetime
	Timestamp
)
