// The logic for the Discover Node.
package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"go-logic/commons/helpers"
	"go-logic/model"
	"log"
	"os"
	"strconv"

	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/ssh"
)

/*
	1. Parse the flags.
		Make sure the flags aren't empty, and make sure the flags are valid.
	2. Confirm the validation by checking the IP by pinging the VM, and getting the details.
	3. After getting the details, enter the details to the DB.
*/

// Setting the requirements of the command.
var node_ip, node_name, node_user, node_pass string
var discoverNodeFlags = flag.NewFlagSet(model.Discover_node, flag.ContinueOnError)
var count = 1
var node_details model.Node_Details

// Function to validate the flags aren't empty.
func validate_flag_checker_empty(mapvariables map[string]string) {
	for key, value := range mapvariables {
		if value == "" {
			model.CommonError_EmptyValue(key)
		}
	}
}

// Function to validate the flags are valid.
func validate_flag_checker_valid(flags map[string]string) bool {
	// Define a new interface for the validation and share the values.
	validation_flags := make(map[string]interface{})
	for key, val := range flags {
		validation_flags[key] = val
	}

	// Define the rules for the validation of the flags.
	rules := map[string]interface{}{
		"node_ip":   "required,max=15,min=7,ip4_addr",
		"node_name": "required,max=15",
		"node_user": "required",
		"node_pass": "required",
	}

	// Create a validator object.
	validate := validator.New()
	errs := validate.ValidateMap(validation_flags, rules)
	if len(errs) > 0 {
		fmt.Println(errs)
		return false
	}
	return true
}

// Function to fetch relevant details to be stored for better reference.
func get_hardware_details(conn *ssh.Client) {
	// Create the SSH Session to get the Hardware Details of the device.
	session, err := helpers.Create_SSH_Session(conn)
	if err != nil {
		//log.Fatalf(err.Error())
		msg := err
		err = model.CallError(model.ErrorCreateSSHSession, "Failed to create an SSH Session")
		if err != nil {
			fmt.Println(err)
			log.Fatal(msg)
			os.Exit(1)
		}
	}
	defer session.Close()

	// Extract the details.
	command := "lscpu --json"
	details, err := session.Output(command)
	if err != nil {
		msg := err
		err = model.CallError(model.ErrorFailureToGetDetails, "Failed to get details of the command.")
		if err != nil {
			fmt.Println(err)
			log.Fatal(msg)
			os.Exit(1)
		}
	}

	// temp variable to hold the complete json data.
	/*
		eg. {
			"lscpu": [{
				"field": "Architecture:",
				"data": "x86"
			},{
				...
			}
			]
		}
	*/
	var lscpu_output map[string][]map[string]string
	err = json.Unmarshal(details, &lscpu_output)
	if err != nil {
		msg := err
		err = model.CallError(model.ErrorJSONUnmarshal, "Failed to unmarshal output.")
		if err != nil {
			fmt.Println(err)
			log.Fatal(msg)
			os.Exit(1)
		}
	}
	for _, item := range lscpu_output["lscpu"] {
		switch item["field"] {
		case "Architecture:":
			node_details.Node_Arch = item["data"]
		case "CPU(s):":
			node_details.Node_CPU = item["data"]
		case "Model name:":
			node_details.Node_Model = item["data"]
		}
	}
}

func get_software_details(conn *ssh.Client) {
	// Create the SSH Session to get the Software Details of the device.
	session, err := helpers.Create_SSH_Session(conn)
	if err != nil {
		msg := err
		err = model.CallError(model.ErrorCreateSSHSession, "Failed to create an SSH Session")
		if err != nil {
			fmt.Println(err)
			log.Fatal(msg)
			os.Exit(1)
		}
	}
	defer session.Close()

	// Extract the details.
	command := "lsb_release -d | awk {'print $2,$3'}"
	details, err := session.Output(command)
	if err != nil {
		msg := err
		err = model.CallError(model.ErrorFailureToGetDetails, "Failed to get details of the command.")
		if err != nil {
			fmt.Println(err)
			log.Fatal(msg)
			os.Exit(1)
		}
	}
	os_det_str := strings.Fields(string(details))
	node_details.Node_OS = os_det_str[0]
	node_details.Node_OS_Ver = os_det_str[1]
}

// Function that will confirm the details before inserting them to CDB.
func confirm_validation(node_details map[string]string) {
	// Validate the IP address is legit. We validated its syntax, that's it.
	flag := helpers.Validate_IP(node_details["node_ip"])
	if flag {
		conn, err := helpers.Create_SSH_Connection(node_details["node_ip"], node_details["node_user"], node_details["node_pass"])
		if err != nil {
			//log.Fatalf(err.Error())
			msg := err
			err = model.CallError(model.ErrorCreateSSHSession, "Error in creating a SSH Session")
			if err != nil {
				fmt.Println(err)
				log.Fatal(msg)
				os.Exit(1)
			}
		}
		defer conn.Close()
		// We have got the hardware and software details.
		get_hardware_details(conn)
		get_software_details(conn)
	} else {
		//log.Println("Invalid IP or Application down!")
		err := model.CallError(model.NodeErrStatus, "Invalid IP Address or Node Down!")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

// Initialize to get the values of the flags.
func init() {
	discoverNodeFlags.StringVar(&node_ip, "ip", "", "Address of the Node.")
	discoverNodeFlags.StringVar(&node_name, "name", "", "Name of the Node.")
	discoverNodeFlags.StringVar(&node_user, "user", "", "User of the Node.")
	discoverNodeFlags.StringVar(&node_pass, "pass", "", "Pass of the Node.")
}

func DiscoverNode(args []string) {
	// After parsing the various flags.
	err := discoverNodeFlags.Parse(args)
	if err != nil {
		// Handling the error for the parsing.
		msg := err
		err = model.CallError(model.FlagsNotLoaded, "Error in parsing the flags.")
		if err != nil {
			fmt.Println(err)
			log.Fatal(msg)
			os.Exit(1)
		}
	}

	// Create a map of the flags, and validate them.
	required_flag_checker := map[string]string{
		"node_ip":   node_ip,
		"node_name": node_name,
		"node_user": node_user,
		"node_pass": node_pass,
	}

	// Check if the flags parsed are empty. If passes, then no issues.
	validate_flag_checker_empty(required_flag_checker)

	// Let us validate further.
	flag := validate_flag_checker_valid(required_flag_checker)
	if flag {
		// fmt.Println("Valid flags")
		confirm_validation(required_flag_checker)
		// We can now post the details of the struct to the DB.
		db, err := helpers.Database_Connection(os.Getenv("MYSQL_DEV_IP"), os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS"), "Aether_DB")
		if err != nil {
			msg := err
			err = model.CallError(model.ErrorMySQLConnectionFail, "Failed to connect to the DB, please check...")
			log.Fatal(msg)
			os.Exit(1)
		}
		defer db.Close()
		// Marshal the data.
		jsonData, err := json.Marshal(node_details)
		if err != nil {
			msg := err
			err = model.CallError(model.ErrorJSONMarshal, "Failure in marshaling the data...")
			log.Fatal(msg)
			os.Exit(1)
		}

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

		// Making the Node Count Dynamic.
		count_stmt := "SELECT COUNT(*) FROM Aether_Node"
		err = db.QueryRow(count_stmt).Scan(&count)
		if err != nil {
			msg := err
			err = model.CallError(model.ErrorMySQLQueryFail, "Failed to run the query...")
			if err != nil {
				fmt.Println(err)
				log.Fatal(msg)
				os.Exit(1)
			}
		}

		// If the count is zero, then make it one as the first entry.
		if count == 0 {
			count = 1
		} else {
			count = count + 1
		}

		// Prepare the Insert query.
		// NOTE: Make values Node_Count dynamic.
		insert_stmt := "INSERT INTO Aether_Node VALUES (" + strconv.Itoa(count) + ",\"" + node_name + "\",\"" + node_ip + "\",\"" + node_user + "\",\"" + node_pass + "\",'" + string(jsonData) + "')"
		//fmt.Println(insert_stmt)
		insert_db, err := db.Query(insert_stmt)
		if err != nil {
			msg := err
			err = model.CallError(model.ErrorMySQLQueryFail, "Failure in running the query...")
			if err != nil {
				fmt.Println(err)
				log.Fatal(msg)
				os.Exit(1)
			}
		}
		defer insert_db.Close()

	} else {
		err = model.CallError(model.ErrorInFlags, "Invalid flags passed, please check them.")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
