#!/bin/bash

cd broker/server
echo -e "Creating broker server binary: \n$(pwd)"
rm -rf bin
go build -o ./bin/server

cd ..
cd client
echo -e "creating broker client binary: \n$(pwd)"
rm -rf bin
go build -o ./bin/client
