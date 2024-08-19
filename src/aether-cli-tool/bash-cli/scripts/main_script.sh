#!/bin/bash

# Load the JSON data into variables.
MAIN_MENU_FILE="../menu/main-menu.json"
CLI_ERROR_FILE="../errors/cli_errors.json"
PROGRAM_ERROR_CODES="../errors/functionality_errors.json"

# Load the CLI_ERRORS
CLI_INVALID_OPTION=$(jq -r '.invalid_choice' "$CLI_ERROR_FILE")

# Function to load the main-menu
main_menu() {
    jq -r '.options[] | "\(.id) \(.description)"' "$MAIN_MENU_FILE"
}


#1.1 Helper for Discover Node -- handle_error
# handle_error() {
#     local exit_code=$1
#     echo "Handling error code: $exit_code"

#     # Extract the error message by iterating through the array and matching the code
#     error_message=$(jq -r --arg code "$exit_code" '.exit_code[] | select(keys[] == $code) | .[$code]' "$PROGRAM_ERROR_CODES")

#     echo "Error message from JSON: $error_message"

#     if [ -n "$error_message" ]; then
#         echo "Error: $error_message"
#     else
#         echo "Undetected error"
#     fi
# }



# 1. Discover Node Service
discover_node_service() {
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
    RESULT=$(./main discover_node  -name=$NODE_NAME -ip=$NODE_IP_ADDR -user=$NODE_USER -pass=$NODE_PASS)
    # EXIT_CODE=$?

    # if [ $EXIT_CODE -ne 0 ]; then
    #     handle_error "$EXIT_CODE"
    #     exit $EXIT_CODE
    # else
    #     echo "Discovery of $NODE_NAME successful!"
    # fi

}

main() {
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
}

main