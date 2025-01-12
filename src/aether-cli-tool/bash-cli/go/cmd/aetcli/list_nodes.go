// Command Part of AetCLI.
// This command will basically fetch all the Node Details that have been registered in the CDB.
// This command will pick the details from the cache layer of Redis. If details are not present, then fetch from the CDB.
// There will be a seperate root command that will fetch the details as per the provider and system set on.
package aetcli

// Algorithm of the function:
// 1. Make Connection to the Redis Layer.
// 2. Check if the Node Details are present in the Cache Layer.
// 3. If Present, fetch only the Node IP, ping it and confirm the status.
// 3.1 Present the details appropriately.
// 4. If not present, fetch the details from CDB, store it in Cache Layer while checking status.

import (
	"fmt"
)

func ListNodes() {
	fmt.Println("List Nodes Command...")
}
