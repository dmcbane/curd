#!/bin/bash

# Test script for CURD
# Handles Go environment conflicts by unsetting GOROOT

echo "Running CURD tests..."
unset GOROOT
go test ./...

if [ $? -eq 0 ]; then
    echo "All tests passed!"
else
    echo "Some tests failed!"
    exit 1
fi