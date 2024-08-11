package cmd

import (
	"flag"
	"fmt"
	"go-logic/model"
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
	discoverNodeFlags.Parse(args)
	fmt.Println(model.Discover_node)
	fmt.Println(node_ip)
	fmt.Println(node_name)
	fmt.Println(node_user)
	fmt.Println(node_pass)
}
