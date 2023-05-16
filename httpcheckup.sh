#!/bin/bash

curl -X POST -H  "acceptlication/json" -H "Content-Type: application/json" -H "using-sandwich-raw-packet: 123" "http://127.0.0.1:8088" -d '{"uri":"internal_command.CmdCheckup","raw":"e30=","passThrough":"123"}' --connect-timeout 5