package codex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToSqlVisitorAccept(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(nil)
	assert.Nil(t, err)
	assert.Equal(t, "?", sql)
	assert.Equal(t, []interface{}{nil}, args)
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
	sql, args, err := NewToSqlVisitor().Accept(Equal(Column("a"), 2))
	assert.Nil(t, err)
	assert.Equal(t, `"a"=?`, sql)
	assert.Equal(t, []interface{}{2}, args)
}

func TestToSqlVisitorEqualNil(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Equal(Column("a"), nil))
	assert.Nil(t, err)
	assert.Equal(t, `"a" IS NULL`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorNotEqual(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(NotEqual(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "?!=?", sql)
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestToSqlVisitorNotEqualNil(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(NotEqual(Column("a"), nil))
	assert.Nil(t, err)
	assert.Equal(t, `"a" IS NOT NULL`, sql)
	assert.Empty(t, args)
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

func TestToSqlVisitorInEmpty(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(In(1))
	assert.Nil(t, err)
	assert.Equal(t, "? IN()", sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestToSqlVisitorInOne(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(In(1, 2))
	assert.Nil(t, err)
	assert.Equal(t, "? IN(?)", sql)
	assert.Equal(t, []interface{}{1, 2}, args)
}

func TestToSqlVisitorInTwo(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(In(1, 2, 3))
	assert.Nil(t, err)
	assert.Equal(t, "? IN(?,?)", sql)
	assert.Equal(t, []interface{}{1, 2, 3}, args)
}

func TestToSqlVisitorInThree(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(In(1, 2, 3, 4))
	assert.Nil(t, err)
	assert.Equal(t, "? IN(?,?,?)", sql)
	assert.Equal(t, []interface{}{1, 2, 3, 4}, args)
}

func TestToSqlVisitorInColums(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(In(Column("x"), Column("a"), Column("b")))
	assert.Nil(t, err)
	assert.Equal(t, `"x" IN("a","b")`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorInLeftError(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(In(Column(".raises error"), Column("a")))
	assert.NotNil(t, err)
	assert.Equal(t, `invalid column name: '.raises error'`, err.Error())
	assert.Equal(t, `-- ERROR --`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorInRightError(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(In(Column("x"), Column(".raises error")))
	assert.NotNil(t, err)
	assert.Equal(t, `invalid column name: '.raises error'`, err.Error())
	assert.Equal(t, `"x" IN(-- ERROR --`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorInError(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(In(1, "word"))
	assert.Nil(t, err)
	assert.Equal(t, "? IN(?)", sql)
	assert.Equal(t, []interface{}{1, "word"}, args)
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

func TestToSqlVisitorUnaliasedTable(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Table("table"))
	assert.Nil(t, err)
	assert.Equal(t, `"table"`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorColumn(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Column("column"))
	assert.Nil(t, err)
	assert.Equal(t, `"column"`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorUnaliasedAttribute(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Attribute(Column("column"), Table("table")))
	assert.Nil(t, err)
	assert.Equal(t, `"table"."column"`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorAliasedAttributeString(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Attribute(Column("column"), Table("table")).As("table_column"))
	assert.Nil(t, err)
	assert.Equal(t, `"table"."column" AS "table_column"`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorAliasedAttributeCol(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Attribute(Column("column"), Table("table")).As(Column("table_column")))
	assert.Nil(t, err)
	assert.Equal(t, `"table"."column" AS "table_column"`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorJoinSource(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(JoinSource(Table("table")))
	assert.Nil(t, err)
	assert.Equal(t, `"table"`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorExtensiveJoinSource(t *testing.T) {
	relation := Table("table")
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
// 	assert.Empty(t, args)
// }

func TestToSqlVisitorCountCol(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Count(Column("id")))
	assert.Nil(t, err)
	assert.Equal(t, `COUNT("id")`, sql)
	assert.Empty(t, args)
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
	assert.Empty(t, args)
}

func TestToSqlVisitorCountEmptyAlias(t *testing.T) {
	c := Count()
	c.Alias = "n"
	sql, args, err := NewToSqlVisitor().Accept(c)
	assert.Nil(t, err)
	assert.Equal(t, `COUNT(*) AS "n"`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorSumInt(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Sum(1))
	assert.Nil(t, err)
	assert.Equal(t, `SUM(?)`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestToSqlVisitorSumCol(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Sum(Column("amount")))
	assert.Nil(t, err)
	assert.Equal(t, `SUM("amount")`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorAvgInt(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Avg(1))
	assert.Nil(t, err)
	assert.Equal(t, `AVG(?)`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestToSqlVisitorAvgCol(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Avg(Column("amount")))
	assert.Nil(t, err)
	assert.Equal(t, `AVG("amount")`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorMinInt(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Min(1))
	assert.Nil(t, err)
	assert.Equal(t, `MIN(?)`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestToSqlVisitorMinCol(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Min(Column("amount")))
	assert.Nil(t, err)
	assert.Equal(t, `MIN("amount")`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorMaxInt(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Max(1))
	assert.Nil(t, err)
	assert.Equal(t, `MAX(?)`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestToSqlVisitorMaxCol(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Max(Column("amount")))
	assert.Nil(t, err)
	assert.Equal(t, `MAX("amount")`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorMaxColAlias(t *testing.T) {
	max := Max(Column("amount"))
	max.Alias = "max_amount"
	sql, args, err := NewToSqlVisitor().Accept(max)
	assert.Nil(t, err)
	assert.Equal(t, `MAX("amount") AS "max_amount"`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorCoalesceCol(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Coalesce(Column("amount"), 0))
	assert.Nil(t, err)
	assert.Equal(t, `COALESCE("amount",?)`, sql)
	assert.Equal(t, []interface{}{0}, args)
}

func TestToSqlVisitorExtensiveFunctionInt(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Sum(1).Or(Count(2).And(Avg(3))))
	assert.Nil(t, err)
	assert.Equal(t, `(SUM(?) OR (COUNT(?) AND AVG(?)))`, sql)
	assert.Equal(t, []interface{}{1, 2, 3}, args)
}

func TestToSqlVisitorExtensiveFunctionCol(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Sum(Column("amount")).Or(Count(Column("id")).And(Avg(Column("volume")))))
	assert.Nil(t, err)
	assert.Equal(t, `(SUM("amount") OR (COUNT("id") AND AVG("volume")))`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorExtensiveFunctionLiteral(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Sum(Literal(`COALESCE("name", 1, 0)`)))
	assert.Nil(t, err)
	assert.Equal(t, `SUM(COALESCE("name", 1, 0))`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorLimit(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Limit(10))
	assert.Nil(t, err)
	assert.Equal(t, `LIMIT ?`, sql)
	assert.Equal(t, []interface{}{10}, args)
}

func TestToSqlVisitorOffset(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Offset(1000))
	assert.Nil(t, err)
	assert.Equal(t, `OFFSET ?`, sql)
	assert.Equal(t, []interface{}{1000}, args)
}

func TestToSqlVisitorHaving(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Having(1))
	assert.Nil(t, err)
	assert.Equal(t, ` HAVING ?`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestToSqlVisitorOnInt(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(On(1))
	assert.Nil(t, err)
	assert.Equal(t, `ON ?`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestToSqlVisitorOnLiteral(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(On(Literal(`users.id = projects.user_id`)))
	assert.Nil(t, err)
	assert.Equal(t, `ON users.id = projects.user_id`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorAscending(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Ascending(Column("date")))
	assert.Nil(t, err)
	assert.Equal(t, `"date" ASC`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorDescending(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Descending(Column("date")))
	assert.Nil(t, err)
	assert.Equal(t, `"date" DESC`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorInnerJoin(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(InnerJoin(Table("foo"), nil))
	assert.Nil(t, err)
	assert.Equal(t, `INNER JOIN "foo"`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorInnerJoinWithOn(t *testing.T) {
	foo := Table("foo")
	bar := Table("bar")
	sql, args, err := NewToSqlVisitor().Accept(InnerJoin(foo, On(foo.Col("id").Eq(bar.Col("foo_id")))))
	assert.Nil(t, err)
	assert.Equal(t, `INNER JOIN "foo" ON "foo"."id"="bar"."foo_id"`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorOuterJoinWithLiteral(t *testing.T) {
	foo := Table("foo")
	sql, args, err := NewToSqlVisitor().Accept(OuterJoin(foo, On(Literal("foo.id = bar.foo_id"))))
	assert.Nil(t, err)
	assert.Equal(t, `LEFT OUTER JOIN "foo" ON foo.id = bar.foo_id`, sql)
	assert.Empty(t, args)
}

func TestToSqlTableSelect(t *testing.T) {
	bar := Table("bar")
	sql, args, err := Table("foo").Select("id", Column("company"), bar.Col("name")).ToSql()
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "foo"."id","company","bar"."name" FROM "foo"`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorSelectStatement(t *testing.T) {
	relation := Table("table")
	stmt := SelectStatement(relation)

	sql, args, err := NewToSqlVisitor().Accept(stmt)
	assert.Nil(t, err)
	assert.Equal(t, `SELECT * FROM "table"`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorSelectStatementExtensive(t *testing.T) {
	foo := Table("foo")
	bar := Table("bar")
	stm := SelectStatement(foo)
	stm.Cols = append(stm.Cols, Column("id"), Column("name"))
	stm.Wheres = append(stm.Wheres, Equal(foo.Col("id"), 1), NotEqual(foo.Col("name"), nil))
	stm.Source.Right = append(stm.Source.Right, InnerJoin(bar, Literal("ON bar.id=foo.bar_id")))

	sql, args, err := NewToSqlVisitor().Accept(stm)
	assert.Nil(t, err)
	assert.Equal(t, `SELECT "id","name" FROM "foo" INNER JOIN "bar" ON bar.id=foo.bar_id WHERE "foo"."id"=? AND "foo"."name" IS NOT NULL`, sql)
	assert.Equal(t, []interface{}{1}, args)
}

func TestToSqlVisitorUnion(t *testing.T) {
	relationOne := Table("table_one")
	relationTwo := Table("table_two")
	relationThree := Table("table_three")
	one := SelectStatement(relationOne)
	two := SelectStatement(relationTwo)
	three := SelectStatement(relationThree)
	one.Combinator = Union(one, two)
	two.Combinator = Union(two, three)

	sql, args, err := NewToSqlVisitor().Accept(one)
	assert.Nil(t, err)
	assert.Equal(t, `(SELECT * FROM "table_one" UNION (SELECT * FROM "table_two" UNION SELECT * FROM "table_three"))`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorIntersect(t *testing.T) {
	relationOne := Table("table_one")
	relationTwo := Table("table_two")
	relationThree := Table("table_three")
	one := SelectStatement(relationOne)
	two := SelectStatement(relationTwo)
	three := SelectStatement(relationThree)
	one.Combinator = Intersect(one, two)
	two.Combinator = Intersect(two, three)

	sql, args, err := NewToSqlVisitor().Accept(one)
	assert.Nil(t, err)
	assert.Equal(t, `(SELECT * FROM "table_one" INTERSECT (SELECT * FROM "table_two" INTERSECT SELECT * FROM "table_three"))`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorExcept(t *testing.T) {
	relationOne := Table("table_one")
	relationTwo := Table("table_two")
	relationThree := Table("table_three")
	one := SelectStatement(relationOne)
	two := SelectStatement(relationTwo)
	three := SelectStatement(relationThree)
	one.Combinator = Except(one, two)
	two.Combinator = Except(two, three)

	sql, args, err := NewToSqlVisitor().Accept(one)
	assert.Nil(t, err)
	assert.Equal(t, `(SELECT * FROM "table_one" EXCEPT (SELECT * FROM "table_two" EXCEPT SELECT * FROM "table_three"))`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorSubSelect(t *testing.T) {
	subTab := Table("sub_table")
	sub := subTab.Select(Sum(Column("s"))).Group(Column("id"))
	one := Table("").Select(Coalesce(sub, Literal("0")).As("total"))

	sql, args, err := NewToSqlVisitor().Accept(one)
	assert.Nil(t, err)
	assert.Equal(t, `(SELECT COALESCE((SELECT SUM("s") FROM "sub_table" GROUP BY "id"),0) AS "total")`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorInsertStatement(t *testing.T) {
	relation := Table("table")
	stmt := InsertStatement(relation)

	sql, args, err := NewToSqlVisitor().Accept(stmt)
	assert.Nil(t, err)
	assert.Equal(t, `INSERT INTO "table" `, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorUpdateStatement(t *testing.T) {
	relation := Table("table")
	stmt := UpdateStatement(relation)
	stmt.Values = []interface{}{Assignment(Column("name"), "Undo")}
	stmt.Limit = Limit(1)

	sql, args, err := NewToSqlVisitor().Accept(stmt)
	assert.Nil(t, err)
	assert.Equal(t, `UPDATE "table" SET "name"=?  LIMIT ?`, sql)
	assert.Equal(t, []interface{}{"Undo", 1}, args)
}

func TestToSqlVisitorUpdateStatementTwoCols(t *testing.T) {
	relation := Table("table")
	stmt := UpdateStatement(relation)
	stmt.Values = []interface{}{Assignment(Column("name"), "Undo"), Assignment(Column("enabled"), true)}
	stmt.Limit = Limit(1)

	sql, args, err := NewToSqlVisitor().Accept(stmt)
	assert.Nil(t, err)
	assert.Equal(t, `UPDATE "table" SET "name"=?,"enabled"=?  LIMIT ?`, sql)
	assert.Equal(t, []interface{}{"Undo", true, 1}, args)
}

func TestToSqlVisitorDeleteStatement(t *testing.T) {
	relation := Table("table")
	stmt := DeleteStatement(relation)

	sql, args, err := NewToSqlVisitor().Accept(stmt)
	assert.Nil(t, err)
	assert.Equal(t, `DELETE FROM "table" `, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorAssignment(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Assignment(Column("a"), 2))
	assert.Nil(t, err)
	assert.Equal(t, `"a"=?`, sql)
	assert.Equal(t, []interface{}{2}, args)
}

func TestToSqlVisitorNot(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Not(Column("b")))
	assert.Nil(t, err)
	assert.Equal(t, `NOT("b")`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorLiteralNoArgs(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Literal("id IS NOT NULL"))
	assert.Nil(t, err)
	assert.Equal(t, `id IS NOT NULL`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorLiteralTwoArgs(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Literal("id = ? AND name like ?", 1, "Hans%"))
	assert.Nil(t, err)
	assert.Equal(t, `id = ? AND name like ?`, sql)
	assert.Equal(t, []interface{}{1, "Hans%"}, args)
}

func TestToSqlVisitorLiteralExanding(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Literal("id IN(?...)", 1, 2, 3))
	assert.Nil(t, err)
	assert.Equal(t, `id IN(?,?,?)`, sql)
	assert.Equal(t, []interface{}{1, 2, 3}, args)
}

func TestToSqlVisitorStar(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Star())
	assert.Nil(t, err)
	assert.Equal(t, `*`, sql)
	assert.Empty(t, args)
}

// TODO maybe useless?
func TestToSqlVisitorBinding(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(Binding())
	assert.Nil(t, err)
	assert.Equal(t, `?`, sql)
	assert.Empty(t, args)
}

func TestToSqlVisitorVisitString(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept("test")
	assert.Nil(t, err)
	assert.Equal(t, `?`, sql)
	assert.Equal(t, []interface{}{"test"}, args)
}

func TestToSqlVisitorVisitInteger(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(0)
	assert.Nil(t, err)
	assert.Equal(t, `?`, sql)
	assert.Equal(t, []interface{}{0}, args)
}

func TestToSqlVisitorVisitFloat(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(0.25)
	assert.Nil(t, err)
	assert.Equal(t, `?`, sql)
	assert.Equal(t, []interface{}{0.25}, args)
}

func TestToSqlVisitorVisitBool(t *testing.T) {
	sql, args, err := NewToSqlVisitor().Accept(true)
	assert.Nil(t, err)
	assert.Equal(t, `?`, sql)
	assert.Equal(t, []interface{}{true}, args)
}

func TestToSqlVisitorQuoteColumnName(t *testing.T) {
	v := NewToSqlVisitor()
	err := v.QuoteColumnName("foo_Bar1", v)
	assert.Nil(t, err)
	assert.Equal(t, `"foo_Bar1"`, v.String())
}

func TestToSqlVisitorQuoteColumnNameDolar(t *testing.T) {
	v := NewToSqlVisitor()
	err := v.QuoteColumnName("foo$", v)
	assert.Nil(t, err)
	assert.Equal(t, `"foo$"`, v.String())
}

func TestToSqlVisitorQuoteColumnNameWithQuoteReturnsError(t *testing.T) {
	v := NewToSqlVisitor()
	err := v.QuoteColumnName(`id" baaaam`, v)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid column name: 'id\" baaaam'", err.Error())
	assert.Equal(t, `-- ERROR --`, v.String())
}

func TestToSqlVisitorQuoteColumnNameLeadingNumReturnsError(t *testing.T) {
	v := NewToSqlVisitor()
	err := v.QuoteColumnName(`1foo`, v)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid column name: '1foo'", err.Error())
	assert.Equal(t, `-- ERROR --`, v.String())
}

func TestToSqlVisitorQuoteTableName(t *testing.T) {
	v := NewToSqlVisitor()
	err := v.QuoteTableName("foo_Bar2", v)
	assert.Nil(t, err)
	assert.Equal(t, `"foo_Bar2"`, v.String())
}

func TestToSqlVisitorQuoteTableNameDolar(t *testing.T) {
	v := NewToSqlVisitor()
	err := v.QuoteTableName("foo$", v)
	assert.Nil(t, err)
	assert.Equal(t, `"foo$"`, v.String())
}

func TestToSqlVisitorQuoteTableNameEmptyReturnsError(t *testing.T) {
	v := NewToSqlVisitor()
	err := v.QuoteTableName(``, v)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid table name: ''", err.Error())
	assert.Equal(t, `-- ERROR --`, v.String())
}

func TestToSqlVisitorQuoteTableNameWithQuoteReturnsError(t *testing.T) {
	v := NewToSqlVisitor()
	err := v.QuoteTableName(`foo" baaam`, v)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid table name: 'foo\" baaam'", err.Error())
	assert.Equal(t, `-- ERROR --`, v.String())
}

func TestToSqlVisitorQuoteTableNameWithLeadingNumReturnsError(t *testing.T) {
	v := NewToSqlVisitor()
	err := v.QuoteTableName(`124foo`, v)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid table name: '124foo'", err.Error())
	assert.Equal(t, `-- ERROR --`, v.String())
}

//         map: 5800 ns/op
// type switch: 4700 ns/op   ... seems to be quite optimized
func BenchmarkVisit(b *testing.B) {
	foo := Table("foo")
	bar := Table("bar")
	stm := SelectStatement(foo)
	stm.Cols = append(stm.Cols, Column("id"), Column("name"))
	stm.Wheres = append(stm.Wheres, Equal(foo.Col("id"), 1), NotEqual(foo.Col("name"), nil))
	stm.Source.Right = append(stm.Source.Right, InnerJoin(bar, Literal("ON bar.id=foo.bar_id")))

	for n := 0; n < b.N; n++ {
		NewToSqlVisitor().Accept(stm)
	}
}
