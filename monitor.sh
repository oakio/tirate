#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Usage: $0 \"USDRUB\" 5s"
    exit 1
fi

clear
while /bin/true; do
    ./tirate.exe | grep $1
    sleep $2
done