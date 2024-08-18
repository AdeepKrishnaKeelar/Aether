package model

// Enumerated list of the commands.
const (
	Discover_node = "discover_node" // command for the node discovery.
)

// Struct for the Discovery Service
// Sub struct for the Discovery Service
type Node_Details struct {
	Node_Arch   string `json:"Architecture:"`
	Node_CPU    string
	Node_Model  string
	Node_OS     string
	Node_OS_Ver string
}

type Node_Info struct {
	Node_Name    string
	Node_IP      string
	Node_User    string
	Node_Pass    string
	Node_Details Node_Details
}
