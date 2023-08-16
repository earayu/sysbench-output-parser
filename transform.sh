#!/bin/bash

path=$1

# Check if the given path exists
if [ ! -d "$path" ]; then
  echo "Path does not exist"
  exit 1
fi

# Iterate through all txt files in the path
for file in "$path"/*.txt; do
  # Check if the txt file exists
  if [ -e "$file" ]; then
    # Get the file name (without path and extension)
    filename=$(basename "$file" .txt)
    # Call the program to process the txt file and output to a csv file
    ./sysparser --file="$file" > "$path/$filename.csv"
    echo "Processing completed: $file"
  fi
done