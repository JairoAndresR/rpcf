#!/bin/bash
kill -9 $(lsof -i:8080 -t)
cd /opt/app
/usr/local/go/bin/go build -o .rpcf_dev
/opt/app/.rpcf_dev 
