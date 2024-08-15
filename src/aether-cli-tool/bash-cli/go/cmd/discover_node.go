package cmd

import (
	"flag"
	"fmt"
	"go-logic/commons"
	"go-logic/model"
	"os"

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
		fmt.Println("Valid Details!")
	} else {
		fmt.Println("Error detected in flags.")
		os.Exit(1)
	}
}
