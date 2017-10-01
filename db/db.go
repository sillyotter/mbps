// Package db exposes a lightweight abstraction over the SQLite code.
// It performs some basic mapping of lower-level types to rqlite types.
package db

import (
	"expvar"

	"github.com/boltdb/bolt"
)

const bkDelay = 250

const (
	numExecutions      = "executions"
	numExecutionErrors = "execution_errors"
	numQueries         = "queries"
	numETx             = "execute_transactions"
	numQTx             = "query_transactions"
)

// DBVersion is the SQLite version.
var DBVersion string

// stats captures stats for the DB layer.
var stats *expvar.Map

func init() {
	stats = expvar.NewMap("db")
	stats.Add(numExecutions, 0)
	stats.Add(numExecutionErrors, 0)
	stats.Add(numQueries, 0)
	stats.Add(numETx, 0)
	stats.Add(numQTx, 0)
}

// DB is the SQL database.
type DB struct {
	path string // Path to database file.
	db   *bolt.DB
}

// Rows asdf
type Rows struct {
	data []byte
}

// Result asdf
type Result struct {
	data []*Rows
}

// Open opens a file-based database, creating it if it does not exist.
func Open(dbPath string) (*DB, error) {
	xdb, err := bolt.Open(dbPath, 0600, nil)

	if err != nil {
		return nil, err
	}

	return &DB{
		path: dbPath,
		db:   xdb,
	}, nil
}

// Close closes the underlying database connection.
func (db *DB) Close() error {
	return db.Close()
}

// Backup writes a consistent snapshot of the database to the given file.
func (db *DB) Backup(path string) error {
	return nil
}

/*
All I have to do is make the db bit work right, and expose the needed operations
then modify store to handle the right operations
then handle the http server to accept the commands
and make sure they get here
and make the client invoke them for testing purposes
*/
