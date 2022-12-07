#!/bin/bash


# check go code format
formatted_files=$(go fmt ./...)

if [[ $formatted_files ]]; then

    echo -e "the following files have format issues\n"
    echo -e "${formatted_files// /\\n}"
    exit 1

else
    echo "no file format issue"
    exit 0
fi
