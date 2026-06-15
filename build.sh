#!/bin/bash

# Build script for CURD
# Handles Go environment conflicts by unsetting GOROOT

echo "Building CURD..."
unset GOROOT
go build -o curd

if [ $? -eq 0 ]; then
    echo "Build successful! Binary created: ./curd"
else
    echo "Build failed!"
    exit 1
fi