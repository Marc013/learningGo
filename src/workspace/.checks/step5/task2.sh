#!/bin/bash

TEST_FILE="step5/calculator_test.go"

if ! grep -q "func TestMain(m \*testing.M)" "$TEST_FILE"; then
    echo "<failed> There is no function called 'TestMain'"
    exit 1
fi

PATH=$PATH:~/workspace/.bin
chmod +x ./.bin/jq

output=$(go test -v -json -count=1 ./...)

if [ $? -gt 1 ]; then
    echo "<failed> Your code could not be executed. Perhaps there is a syntax error. See below for details:"
    exit 1
fi

json=$(jq -s '.' <<< "$output")

# Does some test pass?
query="[.[] | select(.Package == \"getting-started-with-go-unit-testing/step5\" and .Action == \"fail\")] | length"
result=$(jq "$query" <<< "$json")

if [ "$result" -eq 0 ]; then
    echo "<failed> Some tests should have started failing. Did you add 't.Parallel()' in all the tests?"
    exit 1
fi

exit 0
