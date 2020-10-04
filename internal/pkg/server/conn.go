package server

import (
	"fmt"
	"net"
)

const (
	nilString     = "nil"
	okString      = "OK"
	newLineString = "\ngodown > "
)

type conn struct {
	conn net.Conn
}

func newClient(c net.Conn) *conn {
	return &conn{c}
}

func (c *conn) Close() {
	c.conn.Close()
}

func (c *conn) write(str string) {
	fmt.Fprintf(c.conn, str)
}

func (c *conn) writeMessage(msg string) {
	fmt.Fprint(c.conn, msg)
	c.writePrompt()
}

func (c *conn) writeNil() {
	fmt.Fprintf(c.conn, "(%s)", nilString)
	c.writePrompt()
}

func (c *conn) writeString(str string) {
	fmt.Fprintf(c.conn, "(string): %s", str)
	c.writePrompt()
}

func (c *conn) writeInt(val int64) {
	fmt.Fprintf(c.conn, "(integer): %d", val)
	c.writePrompt()
}

func (c *conn) writeError(err error) {
	fmt.Fprintf(c.conn, "(error): %s", err.Error())
	c.writePrompt()
}

func (c *conn) writePrompt() {
	fmt.Fprintf(c.conn, newLineString)
}
