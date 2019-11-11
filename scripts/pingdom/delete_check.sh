#!/bin/bash
email=$1
password=$2
appKey=$3
checkid=$4

curl -X DELETE -u "${email}:${password}" -H "app-key:${appKey}" https://api.pingdom.com/api/2.0/checks/${checkid}