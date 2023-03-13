#!/bin/bash
THE_MODULE="github.com/thanders/quiz-startup"
PROJECT="broker"

services=(broker)
echo "Creating  pb.go and pb.grpc.go files for the broker microservice"

create_proto () {
    for i in ${services[@]}; do
        echo "${i} service"
        protoc -I${PROJECT}/proto --go_opt=module=${THE_MODULE} --go_out=. --go-grpc_opt=module=${THE_MODULE} --go-grpc_out=. ${i}/proto/*.proto
    done
}

create_proto

echo "BuildProto finished"
