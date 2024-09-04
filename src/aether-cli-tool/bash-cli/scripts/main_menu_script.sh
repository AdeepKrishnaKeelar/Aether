#!/bin/bash

# Load the various files as variables.
MAIN_MENU="../menu/main-menu.json"

# Menu function
menu() {
    jq -r '.options[] | "\(.id) \(.description)"' "$MAIN_MENU"
}

# Functions of the main program.
discover_node_service() {
    # The required parameters need to be passed to the program
    printf "Enter Node Name: "
    read NODE_NAME
    printf "Enter Node IP Address: "
    read NODE_IP_ADDR
    printf "Enter Node User: "
    read NODE_USER
    printf "Enter Node Pass: "
    read NODE_PASS
    echo
    cd ../go || exit 1

    # Run the Go program here.
    output=$(go run main.go discover_node  -ip $NODE_IP_ADDR -name $NODE_NAME -pass $NODE_PASS -user $NODE_USER 2>&1)
    exit_code=$?

    if [ $exit_code -ne 0 ]; then
        echo "Program terminated!"
        echo "$output"
    else
        echo "Program success!"
    fi
}

main() {
    # Consider this the start point of main.
    echo "Project Aether"
    while true; do
        menu
        printf "Enter choice: "
        # Read input from the user.
        read CHOICE

        case $CHOICE in
        1)
            discover_node_service
            ;;
        X)
            echo "Bye."
            exit 0
            ;;
        *)
            echo "Error!!"
            ;;
        esac

    done
}

# Call the main function to start the script.
main