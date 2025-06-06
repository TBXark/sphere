package sqlite

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"entgo.io/ent/dialect"
	"modernc.org/sqlite"
)

type Driver struct {
	*sqlite.Driver
}

func (d Driver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})
	if _, e := c.Exec("PRAGMA foreign_keys = on;", nil); e != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("failed to enable enable foreign keys: %w", e)
	}
	return conn, nil
}

func init() {
	sql.Register(dialect.SQLite, Driver{Driver: &sqlite.Driver{}})
}
