#!/bin/bash
kill -9 $(lsof -i:8080 -t)
/usr/local/go/bin/go build -o .rpcf_local
/opt/app/.rpcf_local 
