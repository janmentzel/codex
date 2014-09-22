package codex

import (
	"bytes"
)

type Collector struct {
	sqlBuf bytes.Buffer
	args   []interface{}
}

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
