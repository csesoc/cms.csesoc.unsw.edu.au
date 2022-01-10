package contexts

import (
	"log"

	"cms.csesoc.unsw.edu.au/config"
	"cms.csesoc.unsw.edu.au/environment"
)

// Constants regarding database connections
var USER = config.GetDBUser()
var PASSWORD = config.GetDBPassword()
var DATABASE = config.GetDB()

const HOST_AND_PORT = "db:5432"

const TEST_USER = "postgres"
const TEST_PASSWORD = "test"
const TESTING_DB_NAME = "cms_testing_db"

const TEST_DB_EXPIRY_TIME = 180

// DatabaseContext exposes 3 methods to the user, by using an interface
// it allows us to easilly swap what context (and consequently database) a method is actually using.
// Any connection to the database implements the database context interface
type DatabaseContext interface {
	Query(query string, sqlArgs []interface{}, resultOutput ...interface{}) error
	Exec(query string, sqlArgs []interface{}) error
	Close()
}

// returns a new database context based on the current environment
func GetDatabaseContext() DatabaseContext {
	if environment.IsTestingEnvironment() {
		return newTestingContext()
	}

	context, err := newLiveContext()
	if err != nil {
		log.Fatalf("failed to fetch database context: %v", err)
	}

	return context
}
