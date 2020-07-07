#!/bin/sh
# waitforaccountapi.sh

host="$1"
port="$2"
shift
shift
cmd="$@"
until nc -vzw 2 $host $port; do 
    >&2 echo "Account API service is unavailable - sleeping"
    sleep 1; 
done


# accountapi service for some reason when service gets build 

exec $cmd
