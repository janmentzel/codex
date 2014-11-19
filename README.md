# codex

Codex is **NOT** an ORM, but a relation algebra inspired by [Arel](http://www.github.com/rails/arel) for generating SQL. Project is still early in development but stable.
We made many changes to the original codex code. Our codex is not compatible with the original one anymore.

[![Build Status](https://drone.io/github.com/janmentzel/codex/status.png)](https://drone.io/github.com/janmentzel/codex/latest)
[![Coverage Status](https://coveralls.io/repos/janmentzel/codex/badge.png)](https://coveralls.io/r/janmentzel/codex)

## Installation

With Google's [Go](http://www.golang.org) installed on your machine:

    $ go get -u github.com/janmentzel/codex

## Usage

To show a little sample of what using Codex looks like, lets assume you have a table in your database (we'll call it `users`) and you want to select all the records it contains.  The SQL looks a little like this:

```sql
SELECT * FROM "users"
```

Now using Codex:

```go
import (
  "github.com/janmentzel/codex"
)

users := codex.Table("users")
sql, _, err := users.ToSql()
// sql = SELECT * FROM "users"
```

Now that wasn't too bad, was it?

## SELECT

#### Explicit Columns

```sql
SELECT id, email, first_name, last_name FROM users
```

```go
// ...

users := codex.Table("users")
sql, _, err := users.Select("id", "email", "first_name", "last_name").ToSql()

// sql = SELECT id, email, first_name, last_name FROM users
```

#### WHERE Clause

```sql
SELECT * FROM users WHERE users.id = 123
```

Inlining arguments in SQL queries is SQL-injection prone and not working with prepared queries.
Pass arguments seperately into the DB driver.

```sql
SELECT * FROM users WHERE users.id = ?
```
arguments is an array (`[]interface{}`) with one `int` value here `[123]`

```go
users := codex.Table("users")
sql, args, err := users.Where(users.Col("id").Eq(123)).ToSql()

// sql = `SELECT * FROM users WHERE users.id = ?`
// args = [123]
```

With an `OR`

```sql
SELECT * FROM users WHERE users.id = 123 OR users.email = "test@example.com"
```

arguments separated `[123, "test@example.com"]`
```sql
SELECT * FROM users WHERE users.id = ? OR users.email = ?
```


```go
users := codex.Table("users")
sql, args, err := users.Where(users.Col("id").Eq(1).Or(users.Col("email").Eq("test@example.com"))).ToSql()

// sql = SELECT * FROM users WHERE users.id = ? OR users.email = ?
// args = [123, "test@example.com"]
```

`IN()`
```go
users := codex.Table("users")
sql, args, err := users.Where(users.Col("id").In(1,2,3,4,5)).ToSql()

// sql = SELECT * FROM users WHERE "users"."id" IN(?,?,?,?,?)
// args = [1,2,3,4,5]
```

Or with literal and argument expanding `?...`
```go
users := codex.Table("users")
sql, args, err := users.Where("id IN(?...)", 1, 2, 3, 4, 5).ToSql()

// sql = SELECT * FROM users WHERE id IN(?,?,?,?,?)
// args = [1,2,3,4,5]
```

PostgreSQL array operators with literal and argument expanding
```go
psql := Dialect(POSTGRES)
sql, args, err := psql.Table("products").Where("tags @> ARRAY[?...]", "fancy", "cheap", "retro").ToSql()
// sql = SELECT * FROM products WHERE tags @> ARRAY[$1,$2,$3]
// args = ["fancy","cheap","retro"]
```

The same with expicit table:
```go
psql := Dialect(POSTGRES)
products := psql.Table("products")
sql, args, err := products.Where(products.Col("tags").Literal("@> ARRAY[?...]", "fancy", "cheap", "retro")).ToSql()
// sql = SELECT "products".* FROM "products" WHERE ("products"."tags" @> ARRAY[$1,$2,$3])
// args = ["fancy","cheap","retro"]
```

#### JOIN

```go
users := codex.Table("users")
orders := codex.Table("orders")
sql, args, err := users.Select(orders.Star()).InnerJoin(orders).On(orders("user_id").Eq(users("id"))).ToSql()

// sql = SELECT "users".*, "orders".*
//       FROM "users"
//       INNER JOIN "orders" ON "orders"."user_id" = "users"."id"
// args = []
```

#### Column Alias

```go
companies := codex.Table("companies")
users := codex.Table("users")
q := users.Select(users.Star(), companies.Col("name").As("company_name")).
  InnerJoin(companies).On(companies("id").Eq(users("company_id")))

sql, args, err := q.ToSql()

// sql = SELECT "users".*, "companies"."name" AS "company_name"
//       FROM "users"
//       INNER JOIN "companies" ON "companies"."id" = "users"."company_id"
// args = []
```


## INSERT

```go
sql, args, err := users.Insert("Jon", "Doe", "jon@example.com").
    Into("first_name", "last_name", "email").ToSql()

// sql = `INSERT INTO "users" ("first_name", "last_name", "email") VALUES (?, ?, ?)`
// args = ["Jon", "Doe", "jon@example.com"]
```

## UPDATE

```go
sql, args, err := users.Set("first_name", "last_name", "email").
    To("Jon", "Doe", "jon@example.com").
    Where(users("id").Eq(1)).ToSql()

// sql = UPDATE "users" SET "first_name" = ?, "last_name" = ?, "email" = ?
//       WHERE "users"."id" = ?
// args = ["Jon", "Doe", "jon@example.com", 1]
```

## DELETE

```go
sql, args, err := users.Delete(users("id").Eq(123)).ToSql()

// sql = DELETE FROM "users" WHERE "users"."id" = ?
// args = [123]
```

## CREATE TABLE / ALTER TABLE

DB schema CREATE and ALTER statements are not supported by codex.

codex focus is read/writes queries.

For database migrations you might like to use [goose](https://bitbucket.org/liamstask/goose)



## Documentation

View godoc or visit [godoc.org](http://godoc.org/github.com/janmentzel/codex).

    $ godoc codex

## License

> The MIT License (MIT)

> Copyright (c) 2013 Chuck Preslar

> Copyright (c) 2014 Jan Mentzel, Luzifer Altenberg

> Permission is hereby granted, free of charge, to any person obtaining a copy
> of this software and associated documentation files (the "Software"), to deal
> in the Software without restriction, including without limitation the rights
> to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
> copies of the Software, and to permit persons to whom the Software is
> furnished to do so, subject to the following conditions:

> The above copyright notice and this permission notice shall be included in
> all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
> THE SOFTWARE.
