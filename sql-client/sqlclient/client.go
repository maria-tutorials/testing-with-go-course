package sqlclient

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type client struct {
	db *sql.DB
}

type row struct{}

type SqlClient interface {
	Query(query string, args ...interface{}) (*sqlRows, error)
}

func Open(driverName, dataSourceName string) (SqlClient, error) {
	if driverName == "" {
		return nil, errors.New("invalid driver name")
	}
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	client := &client{
		db: db,
	}
	return client, nil
}

func (c client) Query(query string, args ...interface{}) (*sqlRows, error) {
	rrows, err := c.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	result := sqlRows{
		rows: rrows,
	}

	return &result, nil
}
