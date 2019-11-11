#!/bin/bash
email=$1
password=$2
appKey=$3
object=$4
host=$5
objectUrl=$6

curl -X POST -u "${email}:${password}" -H "app-key:${appKey}" -d "name=${object}&type=http&port=8080&type=httpcustom&host=${host}&url=${objectUrl}" https://api.pingdom.com/api/2.0/checks

