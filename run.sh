#!/bin/bash

go build -o build/gmailbutler

if [ $? -eq 0 ]; then
    echo "Build successful, running the program..."
    ./build/gmailbutler
else
    echo "Build failed."
fi
