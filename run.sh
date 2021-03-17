#!/bin/bash
kill -9 $(lsof -i:8080 -t)
rm -rf .rpcf
go build -o .rpcf
/opt/app/.rpcf 