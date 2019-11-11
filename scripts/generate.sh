#!/bin/bash
crd=$1

current_dir=$PWD

cd $GOPATH/src/k8s.io/code-generator
./generate-groups.sh all operator-v1/pkg/client/${crd} operator-v1/pkg/apis ${crd}:v1
res=$?
if [ $res -ne 0 ]; then
    cd $current_dir
    exit 1
fi

cd $current_dir