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
query='[.[] | select(.Package == "getting-started-with-go-unit-testing/step3" and .Test == "TestCalculateAreaNegativeLength")] | length'
result=$(jq "$query" <<< "$json")

if [ "$result" -eq 0 ]; then
    echo "<failed> There is no function called 'TestCalculateAreaNegativeLength'"
    exit 1
fi

# Was it called with the correct parameters?
query='[.[] | select(.Package == "getting-started-with-go-unit-testing/step3" and .Test == "TestCalculateAreaNegativeLength" and .Action == "output" and (.Output | startswith("CalculateArea called with parameters 5 and -7")))] | length'
result=$(jq "$query" <<< "$json")

if [ "$result" -eq 0 ]; then
    echo "<failed> You did not pass the correct parameters"
    exit 1
fi

# Does the test pass?
query='[.[] | select(.Package == "getting-started-with-go-unit-testing/step3" and .Test == "TestCalculateAreaNegativeLength" and .Action == "pass")] | length'
result=$(jq "$query" <<< "$json")

if [ "$result" -eq 1 ]; then
    echo "<failed> The test passed, but it should have failed; did you write the correct assertions?"
    exit 1
fi

exit 0
