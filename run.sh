#!/usr/bin/env bash

  # resove program dependecie as a creature comfort for the end user


  # Read the key-value pairs from the file

while read -r line; do

  # Split the line into key and value

  key=$(echo "$line" | cut -d '=' -f 1)
  value=$(echo "$line" | cut -d '=' -f 2)

  # Export the key-value pair as an environment variable

  #echo "key: ${key}"
  #echo "value: ${value}"

  export "$key=$value"
done < "$1"

# Execute the command

exec "$@"
