#!/bin/bash

count=1
topic="mytopic"
broker="localhost:9092"

while :
do
    printf -v date '%(%Y/%m/%d-%H:%M:%S)T' -1 
    printf -v message "message n°%s - %s" "$count" "$date"
    printf "topic '%s'\n'%s'\n" "$topic" "$message"
    printf "%s" "$message" | kcat -b $broker -P -t $topic
    count=$((count + 1))
    sleep 10
done

