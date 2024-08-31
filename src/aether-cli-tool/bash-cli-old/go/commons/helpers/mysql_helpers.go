package commons

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var PORT = "3306"

// Function to setup DB connection.
func Database_Connection(db_addr, db_user, db_pass, db_name string) (*sql.DB, error) {
	// DB format --> mysql,db_user:db_pass@tcp(db_addr:PORT)/db_name
	db, err := sql.Open("mysql", db_user+":"+db_pass+"@tcp("+db_addr+":"+PORT+")/"+db_name)
	if err != nil {
		log.Printf("Error in creating connection to DB -- %s", err)
		return nil, err
	}
	//fmt.Println("Trying to establish connection to DB...")
	err = db.Ping()
	if err != nil {
		log.Printf("Error in establishing connection to DB -- %s", err)
		return nil, err
	}
	return db, nil
}
