package cmd

import (
	"flag"
	"fmt"
	"go-logic/model"
	"os"

	_ "github.com/go-playground/validator/v10"
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

func DiscoverNode(args []string) {
	err := discoverNodeFlags.Parse(args)
	if err != nil {
		fmt.Printf("Error parsing flags -- %s", err)
		os.Exit(1)
	}

}
