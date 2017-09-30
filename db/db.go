// Package db exposes a lightweight abstraction over the SQLite code.
// It performs some basic mapping of lower-level types to rqlite types.
package db

import (
	"expvar"
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
	path   string // Path to database file.
	dsn    string // DSN, if any.
	memory bool   // In-memory only.
}

// Result represents the outcome of an operation that changes rows.
type Result struct {
	LastInsertID int64   `json:"last_insert_id,omitempty"`
	RowsAffected int64   `json:"rows_affected,omitempty"`
	Error        string  `json:"error,omitempty"`
	Time         float64 `json:"time,omitempty"`
}

// Rows represents the outcome of an operation that returns query data.
type Rows struct {
	Columns []string        `json:"columns,omitempty"`
	Types   []string        `json:"types,omitempty"`
	Values  [][]interface{} `json:"values,omitempty"`
	Error   string          `json:"error,omitempty"`
	Time    float64         `json:"time,omitempty"`
}

// Open opens a file-based database, creating it if it does not exist.
func Open(dbPath string) (*DB, error) {
	return open("")
}

// OpenWithDSN opens a file-based database, creating it if it does not exist.
func OpenWithDSN(dbPath, dsn string) (*DB, error) {
	return open("")
}

// OpenInMemory opens an in-memory database.
func OpenInMemory() (*DB, error) {
	return open("")
}

// OpenInMemoryWithDSN opens an in-memory database with a specific DSN.
func OpenInMemoryWithDSN(dsn string) (*DB, error) {
	return open("")
}

// LoadInMemoryWithDSN loads an in-memory database with that at the path,
// with the specified DSN
func LoadInMemoryWithDSN(dbPath, dsn string) (*DB, error) {

	return nil, nil
}

// Close closes the underlying database connection.
func (db *DB) Close() error {
	return nil
}

func open(dbPath string) (*DB, error) {
	return &DB{
		path: dbPath,
	}, nil
}

// EnableFKConstraints allows control of foreign key constraint checks.
func (db *DB) EnableFKConstraints(e bool) error {

	return nil
}

// FKConstraints returns whether FK constraints are set or not.
func (db *DB) FKConstraints() (bool, error) {
	return false, nil
}

// Execute executes queries that modify the database.
func (db *DB) Execute(queries []string, tx, xTime bool) ([]*Result, error) {

	return nil, nil
}

// Query executes queries that return rows, but don't modify the database.
func (db *DB) Query(queries []string, tx, xTime bool) ([]*Rows, error) {

	var allRows []*Rows

	return allRows, nil
}

// Backup writes a consistent snapshot of the database to the given file.
func (db *DB) Backup(path string) error {
	return nil
}
