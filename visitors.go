package codex

import (
  "codex/tree/visitors"
)

var VISITORS = map[interface{}]visitors.VisitorInterface{
  "to_sql":   &visitors.ToSqlVisitor{},
  "postgres": &visitors.PostgresVisitor{&visitors.ToSqlVisitor{}},
}
