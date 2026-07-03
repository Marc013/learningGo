#!/bin/bash

PATH=$PATH:~/workspace/.bin
chmod +x ./.bin/jq

output=$(go test -v -json -count=1 ./...)

if [ $? -gt 1 ]; then
    echo "<failed> Your code could not be executed. Perhaps there is a syntax error. See below for details:"
    exit 1
fi

json=$(jq -s '.' <<< "$output")

# Does the test exist?
query='[.[] | select(.Package == "getting-started-with-go-unit-testing/step2" and .Test == "TestCalculateArea")] | length'
result=$(jq "$query" <<< "$json")

if [ "$result" -eq 0 ]; then
    echo "<failed> There is no function called 'TestCalculateArea'"
    exit 1
fi

# Does the test pass?
query='[.[] | select(.Package == "getting-started-with-go-unit-testing/step2" and .Test == "TestCalculateArea" and .Action == "pass")] | length'
result=$(jq "$query" <<< "$json")

if [ "$result" -eq 0 ]; then
    echo "<failed> The test failed, but it should have passed; your test function should be empty'?"
    exit 1
fi

exit 0
