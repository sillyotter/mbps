// Package db exposes a lightweight abstraction over the SQLite code.
// It performs some basic mapping of lower-level types to rqlite types.
package db

import (
	"database/sql/driver"
	"expvar"
	"fmt"
	"strings"
)

const bkDelay = 250

const (
	fkChecks         = "PRAGMA foreign_keys"
	fkChecksEnabled  = "PRAGMA foreign_keys=ON"
	fkChecksDisabled = "PRAGMA foreign_keys=OFF"

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
	return open(fqdsn(dbPath, ""))
}

// OpenWithDSN opens a file-based database, creating it if it does not exist.
func OpenWithDSN(dbPath, dsn string) (*DB, error) {
	return open(fqdsn(dbPath, dsn))
}

// OpenInMemory opens an in-memory database.
func OpenInMemory() (*DB, error) {
	return open(fqdsn(":memory:", ""))
}

// OpenInMemoryWithDSN opens an in-memory database with a specific DSN.
func OpenInMemoryWithDSN(dsn string) (*DB, error) {
	return open(fqdsn(":memory:", dsn))
}

// LoadInMemoryWithDSN loads an in-memory database with that at the path,
// with the specified DSN
func LoadInMemoryWithDSN(dbPath, dsn string) (*DB, error) {
	db, err := OpenInMemoryWithDSN(dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
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
	stats.Add(numExecutions, int64(len(queries)))
	if tx {
		stats.Add(numETx, 1)
	}

	type Execer interface {
		Exec(query string, args []driver.Value) (driver.Result, error)
	}

	return nil, nil
}

// Query executes queries that return rows, but don't modify the database.
func (db *DB) Query(queries []string, tx, xTime bool) ([]*Rows, error) {
	stats.Add(numQueries, int64(len(queries)))
	if tx {
		stats.Add(numQTx, 1)
	}

	type Queryer interface {
		Query(query string, args []driver.Value) (driver.Rows, error)
	}

	var allRows []*Rows
	err := func() (err error) {
		return nil
	}()

	return allRows, err
}

// Backup writes a consistent snapshot of the database to the given file.
func (db *DB) Backup(path string) error {
	return nil
}

// normalizeRowValues performs some normalization of values in the returned rows.
// Text values come over (from sqlite-go) as []byte instead of strings
// for some reason, so we have explicitly convert (but only when type
// is "text" so we don't affect BLOB types)
func normalizeRowValues(row []driver.Value, types []string) []interface{} {
	values := make([]interface{}, len(types))
	for i, v := range row {
		if isTextType(types[i]) {
			switch val := v.(type) {
			case []byte:
				values[i] = string(val)
			default:
				values[i] = val
			}
		} else {
			values[i] = v
		}
	}
	return values
}

// isTextType returns whether the given type has a SQLite text affinity.
// http://www.sqlite.org/datatype3.html
func isTextType(t string) bool {
	return t == "text" ||
		t == "" ||
		strings.HasPrefix(t, "varchar") ||
		strings.HasPrefix(t, "varying character") ||
		strings.HasPrefix(t, "nchar") ||
		strings.HasPrefix(t, "native character") ||
		strings.HasPrefix(t, "nvarchar") ||
		strings.HasPrefix(t, "clob")
}

// fqdsn returns the fully-qualified datasource name.
func fqdsn(path, dsn string) string {
	if dsn != "" {
		return fmt.Sprintf("file:%s?%s", path, dsn)
	}
	return path
}
