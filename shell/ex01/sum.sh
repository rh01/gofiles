#!/bin/bash
declare -i sum=0
declare -i i=0
while [ $i -le 100 ]; do
    let sum+=$i
    let i+=2
done
echo "sum of $i is $sum"