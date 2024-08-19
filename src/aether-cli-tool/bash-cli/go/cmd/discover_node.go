package cmd

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	common_errors "go-logic/commons/errors"
	common_helpers "go-logic/commons/helpers"
	"go-logic/model"
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/ssh"
)

// Setting the requirements of the command.
var node_ip, node_name, node_user, node_pass string
var node_details model.Node_Details
var discoverNodeFlags = flag.NewFlagSet(model.Discover_node, flag.ContinueOnError)
var db *sql.DB

func init() {
	discoverNodeFlags.StringVar(&node_ip, "ip", "", "Address of the Node.")
	discoverNodeFlags.StringVar(&node_name, "name", "", "Name of the Node.")
	discoverNodeFlags.StringVar(&node_user, "user", "", "User of the Node.")
	discoverNodeFlags.StringVar(&node_pass, "pass", "", "Pass of the Node.")
}

// Function to validate the flags aren't empty.
func validate_flag_checker_empty(mapvariables map[string]string) {
	for key, value := range mapvariables {
		if value == "" {
			common_errors.CommonError_EmptyValue(key)
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
	session, err := common_helpers.Create_SSH_Session(conn)
	if err != nil {
		log.Fatalf(err.Error())
		os.Exit(1)
	}
	defer session.Close()

	// Extract the details.
	details, err := session.Output("lscpu --json")
	if err != nil {
		log.Fatalf("Failed to get details -- %s", err)
		os.Exit(1)
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
		log.Fatalf("Failed to unmarshal output -- %s", err)
		os.Exit(1)
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
	session, err := common_helpers.Create_SSH_Session(conn)
	if err != nil {
		log.Fatalf(err.Error())
		os.Exit(1)
	}
	defer session.Close()
	// Extract the details.
	details, err := session.Output("lsb_release -d | awk {'print $2,$3'}")
	if err != nil {
		log.Fatalf("Failed to get OS details -- %s", err)
		os.Exit(1)
	}
	os_det_str := strings.Fields(string(details))
	node_details.Node_OS = os_det_str[0]
	node_details.Node_OS_Ver = os_det_str[1]
}

// Function that will confirm the details before inserting them to CDB.
func confirm_validation(node_details map[string]string) {
	// Validate the IP address is legit. We validated its syntax, that's it.
	flag := common_helpers.Validate_IP(node_details["node_ip"])
	if flag {
		conn, err := common_helpers.Create_SSH_Connection(node_details["node_ip"], node_details["node_user"], node_details["node_pass"])
		if err != nil {
			log.Fatalf(err.Error())
			os.Exit(1)
		}
		defer conn.Close()
		// We have got the hardware and software details.
		get_hardware_details(conn)
		get_software_details(conn)
	} else {
		fmt.Println("Invalid IP or Application down!")
		os.Exit(1)
	}
}

func DiscoverNode(args []string) {
	// After parsing the various flags.
	err := discoverNodeFlags.Parse(args)
	if err != nil {
		fmt.Printf("Error parsing flags -- %s", err)
		os.Exit(1)
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
		//fmt.Println("Valid Details!")
		confirm_validation(required_flag_checker)
		// We can now post the details of the struct to the DB.
		db, err := common_helpers.Database_Connection(os.Getenv("MYSQL_DEV_IP"), os.Getenv("MYSQL_DEV_USER"), os.Getenv("MYSQL_DEV_PASS"), "Aether_DB")
		if err != nil {
			log.Fatalf("Error in setting connection -- %s", err)
			os.Exit(1)
		}
		defer db.Close()
		// Marshal the data.
		jsonData, err := json.Marshal(node_details)
		if err != nil {
			log.Fatalf("Error in marshaling data -- %s", err)
			os.Exit(1)
		}
		// Set the DB.
		set_db, err := db.Query("USE Aether_DB")
		if err != nil {
			log.Fatalf("Failed to set DB -- %s", err)
			os.Exit(1)
		}
		defer set_db.Close()

		// Prepare the Insert query.
		// NOTE: Make values Node_Count dynamic.
		insert_stmt := "INSERT INTO Aether_Node VALUES (1,\"" + node_name + "\",\"" + node_ip + "\",\"" + node_user + "\",\"" + node_pass + "\",'" + string(jsonData) + "')"
		//fmt.Println(insert_stmt)
		insert_db, err := db.Query(insert_stmt)
		if err != nil {
			log.Fatalf("Error in running query -- %s", err)
			os.Exit(1)
		}
		defer insert_db.Close()

	} else {
		//fmt.Println("Error detected in flags.")
		os.Exit(1)
	}
}
