#!/bin/bash
# Script to automate the generation of the inventory for Ansible.
# Steps
# 1. Detect the ssh-key used in the system, and appropriately set the path of the private key.
# 2. Assign the correct IP and root credentials to the host.
# 3. Also generate the group_vars folder for ease of hosts.

ssh_key_detection() {
    # Define the path where the ssh key lies.
    SSH_KEY_PATH="$HOME/.ssh"
    SSH_POSSIBILITY=("id_rsa" "id_ed25519" "id_ecdsa")

    # Loop through the possibilities
    for key in "${SSH_POSSIBILITY[@]}"; do
        if [ -f "$SSH_KEY_PATH/$key" ]; then
            echo "$SSH_KEY_PATH/$key"
            return
        fi
    done

    # If Private key does not exist, break, else the key is detected.
    echo "No SSH Key detected, error!"
    exit 1
}

inventory_generation() {
    CURRENT_PATH=$(pwd)

    # If the directory exists in the current location, then clean it up.
    if [ -d "$CURRENT_PATH/inventory" ]; then
        rm -rf "$CURRENT_PATH/inventory"
    fi

    mkdir -p "$CURRENT_PATH/inventory"
    CURRENT_PATH="$CURRENT_PATH/inventory"

    # Setting up the host file and group_vars directory.
    if [ -d "$CURRENT_PATH/group_vars" ]; then
        rm -rf "$CURRENT_PATH/group_vars"
    fi

    mkdir -p "$CURRENT_PATH/group_vars" 
    touch "$CURRENT_PATH/group_vars/master.yml"
    echo "node_type: \"Master\"" > "$CURRENT_PATH/group_vars/master.yml"
    touch "$CURRENT_PATH/group_vars/worker.yml"
    echo "node_type: \"Worker\"" > "$CURRENT_PATH/group_vars/worker.yml"
    touch "$CURRENT_PATH/hosts"

    echo "$CURRENT_PATH"
}

host_file_creation() {
    local PRIVATE_KEY
    local CURRENT_PATH
    PRIVATE_KEY=$(ssh_key_detection)
    CURRENT_PATH=$(inventory_generation)

    # echo "Private Key: $PRIVATE_KEY"
    # echo "Current Path: $CURRENT_PATH"

    # Ensure the file is created before attempting to write
    if [ ! -f "$CURRENT_PATH/hosts" ]; then
        echo "Error: Hosts file not created."
        exit 1
    fi

    # Write into the hosts file
    {
        echo "[master]"
        echo "$1 ansible_user=$2 ansible_ssh_private_key_file=$PRIVATE_KEY ansible_become_password=$3"
        echo ""
        echo "[worker]"

        shift 3 # Skip the first three values of the master node.
        while [[ $# -ge 3 ]]; do
            echo "$1 ansible_user=$2 ansible_ssh_private_key_file=$PRIVATE_KEY ansible_become_password=$3"
            shift 3 # Move to the next set of parameters
        done
    } > "$CURRENT_PATH/hosts" # Use `>` to overwrite the file, or `>>` to append if you call it multiple times, I was messing here lol.

    echo "Hosts file creation completed."
}


host_file_creation "192.168.1.1" "root" "password" \
                   "192.168.1.2" "worker_user" "worker_password" \
                   "192.168.1.3" "worker_user" "worker_password"