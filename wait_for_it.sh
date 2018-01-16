#!/bin/bash

set -e
cmd="$1"

echo "Waiting for mysql"
#until mysql -h "$host" -u root -p"$password" &> /dev/null
until mysql -u root -h 0.0.0.0 -e "exit"; do
    >&2 echo -n "."
    sleep 1
done

>&2 echo "MySQL is up - executing command"
echo "Create Event"
exec $cmd