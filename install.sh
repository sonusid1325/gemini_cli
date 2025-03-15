#!/bin/bash

# Define the source file and the destination directory
SOURCE_FILE="gemini"
DEST_DIR="/usr/bin"

# Check if the source file exists
if [ ! -f "$SOURCE_FILE" ]; then
  echo "Error: $SOURCE_FILE does not exist."
  exit 1
fi

# Move the file to the destination directory
echo "Installing $SOURCE_FILE to $DEST_DIR..."
sudo cp "$SOURCE_FILE" "$DEST_DIR"


echo "Installation complete."
