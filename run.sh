#!/usr/bin/env bash

DIRS=`find . -mindepth 1 ! -path '*/.*' -type d`

for d in $DIRS; do
    pushd . > /dev/null
    cd $d
    for f in *; do 
        if [[ $f == *.go ]]; then
            echo -e "\n[+] Running $f"
            go run $f < input.txt
        fi
    done
    popd > /dev/null
done