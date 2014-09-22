package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToSqlVisitorAccept(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "", sql)
	assert.Equal(t, []interface{}(nil), args)
}

func TestToSqlVisitorGrouping(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Grouping(1))
	assert.Nil(t, err)
	assert.Equal(t, "(?)", sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestToSqlVisitorAnd(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(And(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "? AND ?", sql)
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestToSqlVisitorOr(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Or(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "? OR ?", sql)
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestToSqlVisitorExtensiveGrouping(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(And(1, 2).Or(3))
	assert.Nil(t, err)
	assert.Equal(t, "(? AND ? OR ?)", sql)
	assert.Equal(t, []interface{}{1, 2, 3}, args)
}

func TestToSqlVisitorEqual(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Equal(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "?=?", sql)
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestToSqlVisitorNotEqual(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(NotEqual(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "?!=?", sql)
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestToSqlVisitorGreaterThan(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(GreaterThan(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "?>?", sql)
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestToSqlVisitorGreaterThanOrEqual(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(GreaterThanOrEqual(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "?>=?", sql)
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestToSqlVisitorLessThan(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(LessThan(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "?<?", sql)
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestToSqlVisitorLessThanOrEqual(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(LessThanOrEqual(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "?<=?", sql)
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestToSqlVisitorLike(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Like(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "? LIKE ?", sql)
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestToSqlVisitorUnlike(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Unlike(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "? NOT LIKE ?", sql)
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestToSqlVisitorExtensiveComparison(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Equal(1, 2).Or(NotEqual(3, 4).And(5)))
	assert.Nil(t, err)
	assert.Equal(t, "(?=? OR (?!=? AND ?))", sql)
	assert.Equal(t, []interface{}{1, 2, 3, 4, 5}, args)
}

func TestToSqlVisitorUnaliasedRelation(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Relation("table"))
	assert.Nil(t, err)
	assert.Equal(t, `"table"`, sql)
	assert.Equal(t, []interface{}(nil), args)
}

func TestToSqlVisitorColumn(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Column("column"))
	assert.Nil(t, err)
	assert.Equal(t, `"column"`, sql)
	assert.Equal(t, []interface{}(nil), args)
}

func TestToSqlVisitorUnaliasedAttribute(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Attribute(Column("column"), Relation("table")))
	assert.Nil(t, err)
	assert.Equal(t, `"table"."column"`, sql)
	assert.Equal(t, []interface{}(nil), args)
}

func TestToSqlVisitorJoinSource(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(JoinSource(Relation("table")))
	assert.Nil(t, err)
	assert.Equal(t, `"table"`, sql)
	assert.Equal(t, []interface{}(nil), args)
}

func TestToSqlVisitorExtensiveJoinSource(t *testing.T) {
	relation := Relation("table")
	source := JoinSource(relation)

	// TODO append usefull stuff! not integers
	source.Right = append(source.Right, []interface{}{1, 2, 3}...)

	sql, args, err := NewToSqlVisitor().Accept(source)
	assert.Nil(t, err)
	assert.Equal(t, `"table" ? ? ?`, sql)
	assert.Equal(t, []interface{}{1, 2, 3}, args)
}

// TODO convenience
// func TestToSqlVisitorCountStr(t *testing.T) {
// 	sql, args, err := NewToSqlVisitor().Accept(Count("id"))
// 	assert.Nil(t, err)
// 	assert.Equal(t, `COUNT("id")`, sql)
// 	assert.Equal(t, []interface{}(nil), args)
// }

func TestToSqlVisitorCountCol(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Count(Column("id")))
	assert.Nil(t, err)
	assert.Equal(t, `COUNT("id")`, sql)
	assert.Equal(t, []interface{}(nil), args)
}

func TestToSqlVisitorCountInt(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Count(1))
	assert.Nil(t, err)
	assert.Equal(t, `COUNT(?)`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestToSqlVisitorCountEmpty(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Count())
	assert.Nil(t, err)
	assert.Equal(t, `COUNT(*)`, sql)
	assert.Equal(t, []interface{}(nil), args)
}

// func TestToSqlVisitorSum(t *testing.T) {
// 	sum := Sum(1)
// 	expected := "SUM(1)"
// 	if got, _ := NewToSqlVisitor().Accept(sum); expected != got {
// 		t.Errorf("TestSum was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorAverage(t *testing.T) {
// 	avg := Average(1)
// 	expected := "AVG(1)"
// 	if got, _ := NewToSqlVisitor().Accept(avg); expected != got {
// 		t.Errorf("TestAverage was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorMinimum(t *testing.T) {
// 	min := Minimum(1)
// 	expected := "MIN(1)"
// 	if got, _ := NewToSqlVisitor().Accept(min); expected != got {
// 		t.Errorf("TestMinimum was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorMaximum(t *testing.T) {
// 	max := Maximum(1)
// 	expected := "MAX(1)"
// 	if got, _ := NewToSqlVisitor().Accept(max); expected != got {
// 		t.Errorf("TestMaximum was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorExtensiveFunction(t *testing.T) {
// 	function := Sum(1).Or(Count(2).And(Average(3)))
// 	expected := "(SUM(1) OR (COUNT(2) AND AVG(3)))"
// 	if got, _ := NewToSqlVisitor().Accept(function); expected != got {
// 		t.Errorf("TestExtensiveFunction was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorLimit(t *testing.T) {
// 	limit := Limit(1)
// 	expected := "LIMIT 1"
// 	if got, _ := NewToSqlVisitor().Accept(limit); expected != got {
// 		t.Errorf("TestLimit was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorOffset(t *testing.T) {
// 	offset := Offset(1)
// 	expected := "OFFSET 1"
// 	if got, _ := NewToSqlVisitor().Accept(offset); expected != got {
// 		t.Errorf("TestOffset was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorHaving(t *testing.T) {
// 	having := Having(1)
// 	expected := "HAVING 1"
// 	if got, _ := NewToSqlVisitor().Accept(having); expected != got {
// 		t.Errorf("TestHaving was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorOn(t *testing.T) {
// 	on := On(1)
// 	expected := "ON 1"
// 	if got, _ := NewToSqlVisitor().Accept(on); expected != got {
// 		t.Errorf("TestOn was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorAscending(t *testing.T) {
// 	asc := Ascending(1)
// 	expected := "1 ASC"
// 	if got, _ := NewToSqlVisitor().Accept(asc); expected != got {
// 		t.Errorf("TestAscending was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorDescending(t *testing.T) {
// 	asc := Descending(1)
// 	expected := "1 DESC"
// 	if got, _ := NewToSqlVisitor().Accept(asc); expected != got {
// 		t.Errorf("TestDescending was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorInnerJoin(t *testing.T) {
// 	join := InnerJoin(1, nil)
// 	expected := "INNER JOIN 1"
// 	if got, _ := NewToSqlVisitor().Accept(join); expected != got {
// 		t.Errorf("TestInnerJoin was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorOuterJoin(t *testing.T) {
// 	join := OuterJoin(1, 2)
// 	expected := "LEFT OUTER JOIN 1 2"
// 	if got, _ := NewToSqlVisitor().Accept(join); expected != got {
// 		t.Errorf("TestOuterJoin was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorSelectCore(t *testing.T) {
// 	relation := Relation("table")
// 	core := SelectCore(relation)
// 	expected := `SELECT FROM "table"`
// 	if got, _ := NewToSqlVisitor().Accept(core); expected != got {
// 		t.Errorf("TestSelectCore was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorSelectCoreExtensive(t *testing.T) {
// 	relation := Relation("table")
// 	core := SelectCore(relation)
// 	core.Cols = append(core.Cols, 1, 2)
// 	core.Wheres = append(core.Wheres, 3, 4)
// 	core.Source.Right = append(core.Source.Right, InnerJoin(5, nil))
// 	expected := `SELECT 1, 2 FROM "table" INNER JOIN 5 WHERE 3 AND 4`
// 	if got, _ := NewToSqlVisitor().Accept(core); expected != got {
// 		t.Errorf("TestSelectCoreExtensive was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorSelectStatement(t *testing.T) {
// 	relation := Relation("table")
// 	stmt := SelectStatement(relation)
// 	expected := `SELECT FROM "table"`
// 	if got, _ := NewToSqlVisitor().Accept(stmt); expected != got {
// 		t.Errorf("TestSelectStatement was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorUnion(t *testing.T) {
// 	relationOne := Relation("table_one")
// 	relationTwo := Relation("table_two")
// 	relationThree := Relation("table_three")
// 	one := SelectStatement(relationOne)
// 	two := SelectStatement(relationTwo)
// 	three := SelectStatement(relationThree)
// 	one.Combinator = Union(one, two)
// 	two.Combinator = Union(two, three)
// 	expected := `(SELECT FROM "table_one" UNION (SELECT FROM "table_two" UNION SELECT FROM "table_three"))`
// 	if got, _ := NewToSqlVisitor().Accept(one); expected != got {
// 		t.Errorf("TestUnion was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorIntersect(t *testing.T) {
// 	relationOne := Relation("table_one")
// 	relationTwo := Relation("table_two")
// 	relationThree := Relation("table_three")
// 	one := SelectStatement(relationOne)
// 	two := SelectStatement(relationTwo)
// 	three := SelectStatement(relationThree)
// 	one.Combinator = Intersect(one, two)
// 	two.Combinator = Intersect(two, three)
// 	expected := `(SELECT FROM "table_one" INTERSECT (SELECT FROM "table_two" INTERSECT SELECT FROM "table_three"))`
// 	if got, _ := NewToSqlVisitor().Accept(one); expected != got {
// 		t.Errorf("TestUnion was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorExcept(t *testing.T) {
// 	relationOne := Relation("table_one")
// 	relationTwo := Relation("table_two")
// 	relationThree := Relation("table_three")
// 	one := SelectStatement(relationOne)
// 	two := SelectStatement(relationTwo)
// 	three := SelectStatement(relationThree)
// 	one.Combinator = Except(one, two)
// 	two.Combinator = Except(two, three)
// 	expected := `(SELECT FROM "table_one" EXCEPT (SELECT FROM "table_two" EXCEPT SELECT FROM "table_three"))`
// 	if got, _ := NewToSqlVisitor().Accept(one); expected != got {
// 		t.Errorf("TestUnion was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorInsertStatement(t *testing.T) {
// 	relation := Relation("table")
// 	stmt := InsertStatement(relation)
// 	expected := `INSERT INTO "table" `
// 	if got, _ := NewToSqlVisitor().Accept(stmt); expected != got {
// 		t.Errorf("TestInsertStatement was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorUpdateStatement(t *testing.T) {
// 	relation := Relation("table")
// 	stmt := UpdateStatement(relation)
// 	expected := `UPDATE "table" `
// 	if got, _ := NewToSqlVisitor().Accept(stmt); expected != got {
// 		t.Errorf("TestUpdateStatement was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorDeleteStatement(t *testing.T) {
// 	relation := Relation("table")
// 	stmt := DeleteStatement(relation)
// 	expected := `DELETE FROM "table" `
// 	if got, _ := NewToSqlVisitor().Accept(stmt); expected != got {
// 		t.Errorf("TestDeleteStatement was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorAssignment(t *testing.T) {
// 	assignment := Assignment(1, 2)
// 	expected := "1 = 2"
// 	if got, _ := NewToSqlVisitor().Accept(assignment); expected != got {
// 		t.Errorf("TestAssignment was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorLiteral(t *testing.T) {
// 	number := Literal(1)
// 	str := Literal("1")
// 	n, _ := NewToSqlVisitor().Accept(number)
// 	s, _ := NewToSqlVisitor().Accept(str)
// 	if n != s {
// 		t.Errorf("TestLiteral was expected to return the same result, got %s and %s", n, s)
// 	}
// }

// func TestToSqlVisitorStar(t *testing.T) {
// 	star := Star()
// 	result, _ := NewToSqlVisitor().Accept(star)
// 	if "*" != result {
// 		t.Errorf("TestStar was expected to return *, got %s", result)
// 	}
// }

// func TestToSqlVisitorBinding(t *testing.T) {
// 	binding := Binding()
// 	result, _ := NewToSqlVisitor().Accept(binding)
// 	if "?" != result {
// 		t.Errorf("TestStar was expected to return ?, got %s", result)
// 	}
// }

// func TestToSqlVisitorUnqualifiedColumn(t *testing.T) {
// 	column := UnqualifiedColumn("column")
// 	expected := `"column"`
// 	if got, _ := NewToSqlVisitor().Accept(column); expected != got {
// 		t.Errorf("TestUnqualifiedColumn was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorVisitNotNull(t *testing.T) {
// 	nnull := NotNull([]interface{}{"column"})
// 	expected := `ALTER 'column' SET NOT NULL`
// 	if got, _ := NewToSqlVisitor().Accept(nnull); expected != got {
// 		t.Errorf("TestVisitNotNull was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorVisitUnique(t *testing.T) {
// 	unique := Unique([]interface{}{"column"})
// 	expected := `ADD UNIQUE('column')`
// 	if got, _ := NewToSqlVisitor().Accept(unique); expected != got {
// 		t.Errorf("TestVisitUnique was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorVisitPrimaryKey(t *testing.T) {
// 	pkey := PrimaryKey([]interface{}{"column"})
// 	expected := `ADD PRIMARY KEY('column')`
// 	if got, _ := NewToSqlVisitor().Accept(pkey); expected != got {
// 		t.Errorf("TestVisitPrimaryKey was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorVisitForeignKey(t *testing.T) {
// 	fkey := ForeignKey([]interface{}{UnqualifiedColumn("column")})
// 	fkey.Options = append(fkey.Options, Relation("table"))
// 	expected := `ADD FOREIGN KEY("column") REFERENCES "table"`
// 	if got, _ := NewToSqlVisitor().Accept(fkey); expected != got {
// 		t.Errorf("TestVisitForeignKey was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorVisitDefault(t *testing.T) {
// 	def := Default([]interface{}{"column"})
// 	expected := `ALTER 'column' SET DEFAULT`
// 	if got, _ := NewToSqlVisitor().Accept(def); expected != got {
// 		t.Errorf("TestVisitDefault was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorVisitString(t *testing.T) {
// 	value, expected := `test`, `'test'`
// 	if got, _ := NewToSqlVisitor().Accept(value); expected != got {
// 		t.Errorf("TestVisitString was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorVisitInteger(t *testing.T) {
// 	value, expected := 0, `0`
// 	if got, _ := NewToSqlVisitor().Accept(value); expected != got {
// 		t.Errorf("TestVisitInteger was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorVisitFloat(t *testing.T) {
// 	value, expected := 0.25, `0.25`
// 	if got, _ := NewToSqlVisitor().Accept(value); expected != got {
// 		t.Errorf("TestVisitFloat was expected to return %s, got %s", expected, got)
// 	}
// }

// func TestToSqlVisitorVisitBool(t *testing.T) {
// 	value, expected := true, `'true'`
// 	if got, _ := NewToSqlVisitor().Accept(value); expected != got {
// 		t.Errorf("TestVisitBool was expected to return %s, got %s", expected, got)
// 	}
// }
