#!/bin/bash

set -e
port="$1"

echo "Waiting for mysql"
#until mysql -h "$host" -u root -p"$password" &> /dev/null
until mysql -u root -h 0.0.0.0 -e "exit" -P $port 2>/dev/null; do
    # >&2
    echo -n "."
    sleep 1
done

echo "MySQL is up - executing command"