package postgres

import (
	"database/sql"
	"fmt"
	"time"

	// postgres driver
	_ "github.com/lib/pq"
)

// Db is our database struct used for interacting with the database
type Db struct {
	*sql.DB
}

// New makes a new database using the connection string and
// returns it, otherwise returns the error
func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// Check that our connection is good
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

// ConnString returns a connection string based on the parameters it's given
// This would normally also contain the password, however we're not using one
func ConnString(host string, port int, user string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbName,
	)
}

// Request shape
type Request struct {
	Request_id   int
	User_email   string
	Request_time time.Time
	Request_type string
}

// GetRequestsByEmail is called within our request query for graphql
func (d *Db) GetRequestsByEmail(user_email string) []Request {
	// Prepare query, takes a user_email argument, protects from sql injection
	stmt, err := d.Prepare("SELECT * FROM Request WHERE user_email=$1")
	if err != nil {
		fmt.Println("GetRequestsByEmail Preperation Err: ", err)
	}

	// Make query with our stmt, passing in user_email argument
	rows, err := stmt.Query(user_email)
	if err != nil {
		fmt.Println("GetRequestsByEmail Query Err: ", err)
	}

	// Create Request struct for holding each row's data
	var r Request

	// Create slice of Requests for our response
	requests := []Request{}

	// Copy the columns from row into the values pointed at by r (Request)
	for rows.Next() {
		err = rows.Scan(
			&r.Request_id,
			&r.User_email,
			&r.Request_time,
			&r.Request_type,
		)

		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}

		requests = append(requests, r)
	}

	return requests
}
