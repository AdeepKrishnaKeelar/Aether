// Command Part of AetCLI.
// TODO: Remove the Docker check as it is not required.
// This command will present the layout of the platform through the CLI.
package aetcli

import (
	"context"
	"fmt"
	"go-logic/model"
	"log"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// Algorithm of the Command:
// 1. Present System Details such as -- Hostname, OS details, CPU, Memory, Disk, Network, etcetera.
// 2. Present Containerization Software Check -- Docker and containers.

func SystemOverview() {
	ctx := context.Background()

	// Get the Host Details.
	hostinfo, err := host.Info()
	if err != nil {
		// Failed to get host details.
		msg := err
		err = model.CallError(model.ErrorFailureToGetDetails, "Failed to get host details...")
		if err != nil {
			fmt.Println(err)
			log.Fatal(msg)
			os.Exit(1)
		}
	}

	// Get the Memory Details.
	meminfo, err := mem.VirtualMemory()
	if err != nil {
		// Failed to get memory details.
		msg := err
		err = model.CallError(model.ErrorFailureToGetDetails, "Failed to get memory details...")
		if err != nil {
			fmt.Println(err)
			log.Fatal(msg)
			os.Exit(1)
		}
	}

	// Get the Disk Details.
	diskinfo, err := disk.Usage("/")
	if err != nil {
		// Failed to get Disk details.
		msg := err
		err = model.CallError(model.ErrorFailureToGetDetails, "Failed to get Disk details...")
		if err != nil {
			fmt.Println(err)
			log.Fatal(msg)
			os.Exit(1)
		}
	}

	fmt.Println("Hostname: ", hostinfo.Hostname)
	fmt.Println("OS: ", hostinfo.OS)
	fmt.Println("Total RAM: ", meminfo.Total)
	fmt.Println("Free RAM: ", meminfo.Free)
	fmt.Println("Total Disk Space: ", diskinfo.Total)
	fmt.Println("Free Disk Space: ", diskinfo.Free)
	fmt.Println("Disk Free Percent: ", (100 - diskinfo.UsedPercent))
	fmt.Println("Used Disk Space: ", diskinfo.Used)
	fmt.Println("Disk Usage: ", diskinfo.UsedPercent)

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer cli.Close()

	// List the containers.
	containers, err := cli.ContainerList(ctx, container.ListOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Containers that are active --- ")
	for _, container := range containers {
		fmt.Println(container.Names)
	}
}
