#!/bin/bash


GO_API_SERVER_ROOT="$(dirname $(cd $(dirname $0) >/dev/null 2>&1; pwd -P;))"
cd $GOLANG_API_SERVER_ROOT


# check go code format
formatted_files=$(go fmt ./...)

if [[ $formatted_files ]]; then
    echo -e "the following files have format issues"
    echo -e "${formatted_files// /\\n}"
    exit 1
else
    echo "no file format issue"
    exit 0
fi
