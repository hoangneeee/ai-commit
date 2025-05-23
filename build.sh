#!/bin/bash

# Build the binary
echo "Building aic..."
go build -o aic .

# Make the binary executable
chmod +x aic

echo "Build complete!"
echo "To use 'aic' command, run one of the following:"
echo "1. For temporary use (current terminal session):"
echo "   alias aic='$(pwd)/aic'"
echo "2. To make it permanent, add to your ~/.zshrc or ~/.bashrc:"
echo "   echo \"alias aic='$(pwd)/aic'\" >> ~/.zshrc"
echo "   source ~/.zshrc"
