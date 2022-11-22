#!/bin/bash

for d in */; do
    if [[ "$d" =~ [0-2][0-9] ]]; then
        echo "day $d"
        cd "$d" || 'not found'
        go run main.go
        cd ..
        echo
    fi
done