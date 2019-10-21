#/bin/bash

declare -i i=1
declare -i users=0

while [ $i -le 10 ]; do
    if ! id user$1 &> /dev/null; then
        useradd use${i}
        let users++
    fi
    let i++
done

