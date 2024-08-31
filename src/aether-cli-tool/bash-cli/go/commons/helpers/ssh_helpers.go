// All the helper functions will be housed here.

package helpers

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

// Helper function to validate the IP address.
func Validate_IP(ip_address string) bool {
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

// Create the config variable that can be used to ssh into the VM for details.
func Create_SSH_Connection(node_ip, node_user, node_pass string) (*ssh.Client, error) {
	// Create the config where we shall have an SSH session with the node to extract details.
	// There is no need to utilize the SSH tokens, simple keys would suffice.
	config := &ssh.ClientConfig{
		User: node_user,
		Auth: []ssh.AuthMethod{
			ssh.Password(node_pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	// Establishing connection with the node.
	conn, err := ssh.Dial("tcp", node_ip+":22", config)
	if err != nil {
		return nil, fmt.Errorf("failed to estabish connection with node -- %w", err)
	}
	return conn, nil
}

func Create_SSH_Session(conn *ssh.Client) (*ssh.Session, error) {
	// Create the SSH Session.
	session, err := conn.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create ssh session -- %w", err)
	}
	return session, nil
}
