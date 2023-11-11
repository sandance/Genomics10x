#!/bin/bash

while true
do
    echo "Testing endpoint 1: Filter by weather = rain"
    curl -i weatherapp:8080/query?weather=rain

    echo "Testing endpoint 2: Filter by weather = rain with limit 5"
    curl -i weatherapp:8080/query?weather=rain&limit=5

    echo "Testing endpoint 3: Filter by date = 2012-06-04"
    curl -i weatherapp:8080/query?date=2012-06-04

    sleep 5  # Add a delay between tests
done
