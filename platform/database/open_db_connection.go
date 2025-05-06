package database

import "authservice/app/queries"

// Queries struct for collect all app queries.
type Queries struct {
	*queries.TokenQueries // load queries from Book model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		TokenQueries: &queries.TokenQueries{DB: db}, // from Book model
	}, nil
}
