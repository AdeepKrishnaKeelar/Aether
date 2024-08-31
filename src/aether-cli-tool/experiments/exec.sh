#!/bin/bash

# Run the Go program and capture both the output and the exit code
output=$(go run main.go 2>&1)
exit_code=$?

# Check the exit code. Here, we will realize the error, either 0 or 1.
if [ $exit_code -ne 0 ]; then
    echo "The Go program encountered an error:"
    echo "$output"

    # Now, parsing the error code of the struct defined in the go program.
    # BASH_REMATCH array will have success match in pos 0, and the match in 1.
    if [[ "$output" =~ status\ --\ ([0-9]+) ]]; then
        status_code=${BASH_REMATCH[1]}
        echo "Extracted status code: $status_code"

        # Handling specific status codes
        case $status_code in
            503)
                echo "Handling HTTP 503 Service Unavailable error"
                ;;
            *)
                echo "Handling other errors"
                ;;
        esac
    fi
else
    echo "The Go program ran successfully."
fi
