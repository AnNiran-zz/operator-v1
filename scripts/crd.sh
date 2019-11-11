#!/bin/bash
command=$1
file=$2

# locate kubectl
which kubectl
res=$?
if [ $res -ne 0 ]; then
    exit 1
fi

# execute command on file
# kubectl create -f [filename]
kubectl ${command} -f ${file}
rescmd=$?
if [ $rescmd -ne 0 ]; then
    exit 1
fi