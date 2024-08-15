#!/bin/bash

# Load the JSON data into variables.
MAIN_MENU_FILE="menu/main-menu.json"
CLI_ERROR_FILE="errors/cli_errors.json"

# Load the CLI_ERRORS
CLI_INVALID_OPTION=$(jq -r '.invalid_choice' "$CLI_ERROR_FILE")

# Function to load the main-menu
main_menu() {
    jq -r '.options[] | "\(.id) \(.description)"' "$MAIN_MENU_FILE"
}

# 1. Discover Node Service
discover_node_service() {
    printf "Enter Node IP Address: "
    read NODE_IP_ADDR
    printf "Enter Node User: "
    read NODE_USER
    printf "Enter Node Pass: "
    read NODE_PASS
    echo "./main discover_node -ip=$NODE_IP_ADDR -user=$NODE_USER -pass=$NODE_PASS"
}

# Consider this the start point of main.
echo "Project Aether"
while true; do
    main_menu
    printf "Enter choice: "
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
        echo $CLI_INVALID_OPTION
        ;;
    esac
done