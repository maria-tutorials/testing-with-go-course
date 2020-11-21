package sqlclient

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	goEnv      = "GO_ENVIRONMENT"
	production = "production"
)

var (
	isMocked bool
	dbClient SqlClient
)

func StartMockServer() {
	isMocked = true
}

func StopMockServer() {
	isMocked = false
}

type client struct {
	db *sql.DB
}

type row struct{}

type SqlClient interface {
	Query(query string, args ...interface{}) (rows, error)
}

func isProduction() bool {
	return os.Getenv(goEnv) == production
}

func Open(driverName, dataSourceName string) (SqlClient, error) {

	if isMocked || !isProduction() {
		dbClient := &clientMock{}
		return dbClient, nil
	}

	if driverName == "" {
		return nil, errors.New("invalid driver name")
	}
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	dbClient := &client{
		db: db,
	}
	return dbClient, nil
}

func (c client) Query(query string, args ...interface{}) (rows, error) {
	rrows, err := c.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	result := sqlRows{
		rows: rrows,
	}

	return &result, nil
}
