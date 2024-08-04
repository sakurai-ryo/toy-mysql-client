package main

import (
	"database/sql/driver"
	"net"
)

const (
	HOST = "127.0.0.1"
	PORT = "3306"
)

func connect(user string, password string) (driver.Conn, error) {
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		return nil, err
	}

	// TODO: auth
	if err := readHandshakePacket(conn); err != nil {
		return nil, err
	}

	return &Connection{
		conn: conn,
	}, nil
}

var _ driver.Conn = (*Connection)(nil)

// Connection implements driver.Conn interface
type Connection struct {
	conn net.Conn
}

func (c *Connection) Close() error {
	if c.conn == nil {
		return nil
	}

	return c.conn.Close()
}

func (c *Connection) Prepare(query string) (driver.Stmt, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Connection) Begin() (driver.Tx, error) {
	//TODO implement me
	panic("implement me")
}
