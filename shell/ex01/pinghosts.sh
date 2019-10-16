#!/bin/bash
#
declare -i i=1
declare -i uphosts=0
declare -i downhosts=0
net='172.16.3'

while [ $i -le 254 ]; do
    ping -c 1 -w 1 $net.$i &> /dev/null
    if [ $? -eq 0 ]; then
        echo "${net}.${i} is up".
        let uphosts++
    else
        echo "${net}.${i} is down".
        let downhosts++
    fi
    let i++
done

echo "Host up: $uphosts"
echo "Host down: $downhosts"
