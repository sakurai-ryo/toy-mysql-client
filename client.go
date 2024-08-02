package main

import (
	"log/slog"
	"net"
)

const (
	HOST = "127.0.0.1"
	PORT = "3306"
)

func Client() error {
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		return err
	}

	bug := make([]byte, 1024)
	n, err := conn.Read(bug)
	if err != nil {
		return err
	}
	slog.Info(string(bug[:n]))

	return nil
}
