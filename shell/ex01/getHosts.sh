#!/bin/bash
net='172.16.3'
uphosts=0
downhosts=0
for i in {1..20};  do
    ping -c 1 -w 1 ${net}.${i} &> /dev/null
    if [ $? -eq 0 ]; then
        echo "${net}.${i} is up".
        let uphosts++
    else
        echo "${net}.${i} is down".
        let downhosts++
    fi
done
echo "Up host: $uphosts"
echo "Down host: $downhosts"
