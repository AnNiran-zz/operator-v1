#!/bin/bash
crd=$1

# delete clientset
rm -rf pkg/client/${crd}
delcli=$?
echo $delcli
