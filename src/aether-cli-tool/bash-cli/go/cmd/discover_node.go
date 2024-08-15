package cmd

import (
	"flag"
	"fmt"
	"go-logic/commons"
	"go-logic/model"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/go-playground/validator/v10"
)

// Setting the requirements of the command.
var node_ip, node_name, node_user, node_pass string
var discoverNodeFlags = flag.NewFlagSet(model.Discover_node, flag.ContinueOnError)

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
			commons.CommonError_EmptyValue(key)
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

// Helper function to validate the IP address.
func validate_ip(ip_address string) bool {
	// Here, we attempt to establish a TCP connection.
	// If it fails, then either the connection to be established with is shot or the wrong IP is passed.
	// If the nut passes something that can reach the public domain, like IP of google, then it's his problem.
	port := "22"
	timeout := 2 * time.Second

	conn, err := net.DialTimeout("tcp", ip_address+":"+port, timeout)
	if err != nil {
		// This is a connection failure.
		return false
	}
	defer conn.Close()
	return true
}

// Function to fetch relevant details to be stored for better reference.
func get_details(ip_address, user, pass string) {
	// Create the config where we shall have an SSH session with the node to extract details.
	// There is no need to utilize the SSH tokens, simple keys would suffice.
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}
	// Establishing connection with the node.
	conn, err := ssh.Dial("tcp", ip_address+":22", config)
	if err != nil {
		fmt.Printf("Failed to establish connection -- %s", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Create a session.
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create a session -- %s", err)
		os.Exit(1)
	}
	defer session.Close()

	// Extract the details.
	details, err := session.Output("lscpu --json")
	if err != nil {
		log.Fatalf("Failed to get details -- %s", err)
		os.Exit(1)
	}
	fmt.Println(string(details))

}

// Function that will confirm the details before inserting them to CDB.
func confirm_validation(node_details map[string]string) {
	// Validate the IP address is legit. We validated its syntax, that's it.
	flag := validate_ip(node_details["node_ip"])
	if flag {
		fmt.Println("Success")
		get_details(node_details["node_ip"], node_details["node_user"], node_details["node_pass"])
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
	} else {
		//fmt.Println("Error detected in flags.")
		os.Exit(1)
	}
}
