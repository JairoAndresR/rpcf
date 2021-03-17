#!/bin/bash
kill -9 $(lsof -i:8080 -t)
rm -rf .rpcf
/usr/local/go/bin/go build -o .rpcf
/opt/app/.rpcf 
