package codex

import (
	"bytes"
	"strconv"
	"strings"
)

type CollectorInterface interface {
	AppendSqlStr(s string)
	AppendSqlByte(b byte)
	AppendArg(a interface{})
	String() string
	Args() []interface{}
}

const EXPECTED_SQL_QUERY_LEN = 512

// standard Collector with ? as argument placeholder
type Collector struct {
	sqlBuf bytes.Buffer
	args   []interface{}
}

var _ CollectorInterface = (*Collector)(nil)

func (c *Collector) AppendSqlStr(s string) {
	// n not needed, err always nil according docu
	c.sqlBuf.WriteString(s)
}
func (c *Collector) AppendSqlByte(b byte) {
	// n not needed, err always nil according docu
	c.sqlBuf.WriteByte(b)
}
func (c *Collector) AppendArg(a interface{}) {
	c.args = append(c.args, a)
}

func (c *Collector) String() string {
	return c.sqlBuf.String()
}

func (c *Collector) Args() []interface{} {
	return c.args
}

// creates a Collector with 512 bytes buffer capacity
func NewCollector() *Collector {
	return &Collector{
		sqlBuf: *bytes.NewBuffer(make([]byte, 0, EXPECTED_SQL_QUERY_LEN)),
	}
}

// Postgres speciffic Collector with $1, $2 ... $n as argument placeholder
type PostgresCollector struct {
	sqlBuf bytes.Buffer
	args   []interface{}
	iArg   int
}

var _ CollectorInterface = (*PostgresCollector)(nil)

func (c *PostgresCollector) AppendSqlStr(s string) {
	if strings.ContainsRune(s, QUESTION) {

		// pregrow buffer to avoid iterating reallocation
		n := strings.Count(s, string(QUESTION))

		factor := 1
		if c.iArg+n > 9 {
			factor++
		}
		if c.iArg+n > 99 {
			factor++
		}

		c.sqlBuf.Grow(len(s) + n*factor)

		for _, r := range s {
			if r == QUESTION {
				c.iArg++
				c.sqlBuf.WriteRune('$')
				c.sqlBuf.WriteString(strconv.Itoa(c.iArg))
			} else {
				c.sqlBuf.WriteRune(r)
			}
		}
	} else {
		c.sqlBuf.WriteString(s)
	}
}

func (c *PostgresCollector) AppendSqlByte(b byte) {
	if b == '?' {
		c.iArg++
		c.sqlBuf.WriteRune('$')
		c.sqlBuf.WriteString(strconv.Itoa(c.iArg))
	} else {
		c.sqlBuf.WriteByte(b)
	}
}
func (c *PostgresCollector) AppendArg(a interface{}) {
	c.args = append(c.args, a)
}

func (c *PostgresCollector) String() string {
	return c.sqlBuf.String()
}

func (c *PostgresCollector) Args() []interface{} {
	return c.args
}

// creates a PostgresCollector with 512 bytes buffer capacity
func NewPostgresCollector() *PostgresCollector {
	return &PostgresCollector{
		sqlBuf: *bytes.NewBuffer(make([]byte, 0, EXPECTED_SQL_QUERY_LEN)),
	}
}
