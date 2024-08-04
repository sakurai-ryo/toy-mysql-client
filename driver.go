package main

import (
	"database/sql/driver"
	"log/slog"
	"strings"
)

var _ driver.Driver = (*Driver)(nil)

// Driver implements driver.Driver interface
type Driver struct{}

func (d *Driver) Open(dns string) (driver.Conn, error) {
	slog.Info("Open Connection")

	// 簡略化のため`user:pass`の形式のみ
	authData := strings.Split(dns, ":")
	username := authData[0]
	password := authData[1]

	return connect(username, password)
}
