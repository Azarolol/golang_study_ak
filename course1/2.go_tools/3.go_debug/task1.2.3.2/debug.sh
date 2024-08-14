#!/bin/bash
echo "Debug started..."
dlv debug "$1"
go build -o "myprogram" $1
dlv exec myprogram
echo "Debug ended."