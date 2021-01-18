package utils

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

// NewDbConnection is function to create New DB Connection
func NewDbConnection(host, port, user, password, db string) (*sql.DB, *context.Context, error) {

	oracleSQLConnectionString := fmt.Sprintf(
		`user="%s" password="%s" connectString=%s:%s/%s`,
		user,
		password,
		host,
		port,
		db,
	)

	client, err := sql.Open("godror", oracleSQLConnectionString)

	if err != nil {
		return nil, nil, err
	}

	ctx := context.Background()

	return client, &ctx, nil
}
