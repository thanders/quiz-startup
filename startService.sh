#!/bin/bash
THE_MODULE="github.com/thanders/quiz-startup"
PROJECT="broker"

echo "Start Service will start microservices from their binaries"

case $1 in

  brokerService)
    echo "Starting broker service"
    echo "Binary path - /broker/server/bin/server"
    ./broker/server/bin/server
    ;;

  brokerClient)
    echo "Starting broker client"
    echo "Binary path - /broker/client/bin/client"
    ./broker/client/bin/client
    ;;

  help)
    echo && cat servicesList.csv | column -t -s ","
    ;;

  *)
    STATEMENTS
    ;;
esac