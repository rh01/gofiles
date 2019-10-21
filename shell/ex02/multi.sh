#!/bin/bash

for i in {1..9}; do
    for j in $(seq 1 $i); do
        echo -e -n "${i}x${j}=$[$i*$j]\t" 
    done
    echo
done

