#!/bin/bash
#
declare -i max=0
declare -i min=0
declare -i i=1

max=$RANDOM
min=$max
echo $max

while [ $i -le 9 ]; do
    rand=$RANDOM
    echo $rand
    if [ $rand -gt $max ]; then
        let max=$rand
    fi

    if [ $rand -lt $min ]; then
        let min=$rand
    fi
    let i++
done

echo "max value: $max"
echo "min value: $min"

