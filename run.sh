#!/bin/bash

# Build the Go program
go build -o build/gmailbutler

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo "Build successful, running the program..."
    # Run the built program
    ./build/gmailbutler
else
    echo "Build failed."
fi
