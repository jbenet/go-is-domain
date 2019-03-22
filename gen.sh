#! /bin/bash

# parses tlds-alpha-by-domain.txt to generate map contents 
# for a string -> bool golang mapping.

if [[ -f formatted_tlds.txt ]]; then
    rm formatted_tlds.txt
fi

while IFS= read -r line; do
    # ignore comments in file
    if [[ "$line" =~ ^\#.* ]]; then
        continue
    fi 
    FORMATTED="\"$line\": true,"
    echo "$FORMATTED" >> formatted_tlds.txt 
done < tlds-alpha-by-domain.txt