#!/bin/bash
tempout=$(mktemp)
(docker-entrypoint.sh) > ${tempout} &

(</dev/tcp/localstack/4566) > /dev/null 2>&1
while [ $? -ne 0 ]; do
    echo "Waiting for localstack to start listening on port 4566"
    sleep 5
    (</dev/tcp/localstack/4566) > /dev/null 2>&1
done

while true; do
    # look for a line with just the content "Ready."
    readyline=$(cat $tempout | grep "^Ready\.\$")
    case "$readyline" in
    *Ready\.*)
        echo "Localstack is ready."
        awslocal sqs --create-queue default > /dev/null 2>&1
        echo "Queue has been created!"
        break
        ;;
    *)
        echo "Waiting for localstack to be ready"
        sleep 1
        ;;
    esac
done

# never exit, else the docker container would exit
while true; do
    sleep 5
done