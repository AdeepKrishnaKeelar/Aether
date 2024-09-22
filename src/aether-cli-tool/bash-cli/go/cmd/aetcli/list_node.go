// Command Part of AetCLI.
// This command will basically fetch all the Node Details that have been registered in the CDB.
// There will be a seperate root command that will fetch the details as per the provider and system set on.
package aetcli

import (
	"go-logic/commons/helpers"
	"go-logic/model"
	"log"
	"os"
)

func ListNode() {
	// Make Connection to the CDB.
	// Parse the information of the flags.
	// Appropriately fetch the data from the CDB.
	// Print the data in a tabular, readable format.

	// Variables
	db, err := helpers.Database_Connection(os.Getenv("MYSQL_DEV_IP"), os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS"), "Aether_DB")
	if err != nil {
		msg := err
		err = model.CallError(model.ErrorMySQLConnectionFail, "Failed to connect to the DB, please check...")
		log.Fatal(msg)
		os.Exit(1)
	}
	defer db.Close()
	//fmt.Println("Connection to the DB established...")

	statement := "SELECT Node_Name, Node_IP from Aether_Node"
}
