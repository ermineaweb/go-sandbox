#!/bin/bash

# read the last 3 messages and the next
kcat -C -b localhost:9092 -t mytopic 