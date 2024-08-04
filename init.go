package main

import (
	"database/sql"
)

const driverName = "mysql"

func init() {
	sql.Register(driverName, &Driver{})
}
