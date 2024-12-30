// All the common data structures used in the application will be stored here.
package model

// Enumerated list of the commands.
const (
	Discover_node = "discover_node" // command for the node discovery.
	AetCLI        = "aetcli"        // Aether CLI for Interactive Mode.
	AetCLI_Admin  = "aetcli-admin"  // Aether CLI in Admin Mode.
)

// Sub Consts for AetCLI Command only.
const (
	ListNode       = "list_node"       // List all Nodes that are discovered and in the CDB.
	SystemOverview = "system_overview" // Get the System Overview of the Platform.
)

// Sub Consts for AetCLI_Admin Command only.
const ()

// Struct for the Discovery Service
// Sub struct for the Discovery Service
type Node_Details struct {
	Node_Arch   string `json:"Architecture"`
	Node_CPU    string `json:"CPU"`
	Node_Model  string `json:"Model"`
	Node_OS     string `json:"OS"`
	Node_OS_Ver string `json:"Version"`
}

type Node_Info struct {
	Node_Name    string
	Node_IP      string
	Node_User    string
	Node_Pass    string
	Node_Details Node_Details
}
