// Command Part of AetCLI.
// This command will basically fetch all the Node Details that have been registered in the CDB.
// There will be a seperate root command that will fetch the details as per the provider and system set on.
package aetcli

import (
	"fmt"
	"go-logic/commons/helpers"
	"go-logic/model"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/table"
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
	// fmt.Println("Connection to the DB established...")

	// Task: Make the statement dynamic. No hardcoding.
	statement := "SELECT Node_Name, Node_IP from Aether_Node"

	// Setup DB Connection.
	db, err = helpers.Database_Connection(os.Getenv("MYSQL_DEV_IP"), os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS"), "Aether_DB")
	if err != nil {
		msg := err
		err = model.CallError(model.ErrorMySQLConnectionFail, "Failed to connect to the DB, please check...")
		log.Fatal(msg)
		os.Exit(1)
	}
	defer db.Close()

	// Set the DB.
	query := "USE Aether_DB"
	set_db, err := db.Query(query)
	if err != nil {
		msg := err
		err = model.CallError(model.ErrorMySQLQueryFail, "Failure in running the query...")
		if err != nil {
			fmt.Println(err)
			log.Fatal(msg)
			os.Exit(1)
		}
	}
	defer set_db.Close()

	// Query the db.
	rows, err := db.Query(statement)
	if err != nil {
		msg := err
		err = model.CallError(model.ErrorMySQLQueryFail, "Failure in running the query...")
		if err != nil {
			fmt.Println(err)
			log.Fatal(msg)
			os.Exit(1)
		}
	}
	defer rows.Close()

	var rowsData []table.Row

	// Print the data.
	for rows.Next() {
		var node_name, node_ip string
		err := rows.Scan(&node_name, &node_ip)
		if err != nil {
			msg := err
			err = model.CallError(model.ErrorMySQLQueryFail, "Failure in running query and fetching data...")
			if err != nil {
				fmt.Println(err)
				log.Fatal(msg)
				os.Exit(1)
			}
		}
		rowsData = append(rowsData, table.Row{node_name, node_ip})
	}
	//Append the data to the table view.
	//Setup table writer and render output
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Node Name", "Node IP"})
	t.AppendRows(rowsData)
	t.Render()
}
