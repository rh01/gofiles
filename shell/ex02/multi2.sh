#!/bin/bash

# for i in {1..9}; do
#     for j in $(seq 1 $i); do
#         echo -e -n "${i}x${j}=$[$i*$j]\t" 
#     done
#     echo
# done


declare -i i=1
declare -i j=0

while [ $i -le 9 ]; do
    let j=1
    while [ $j -le $i ]; do
        echo -e -n "${i}x${j}=$[$i*$j]\t" 
        let j++
    done
    echo
    let i++
done

