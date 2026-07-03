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

tests=("TestDivideValid" "TestDivideByZero")
tests+=("TestDivideValid")

for test in "${tests[@]}"
do
  # Does the test exist?
  query="[.[] | select(.Package == \"getting-started-with-go-unit-testing/step5\" and .Test == \"$test\")] | length"
  result=$(jq "$query" <<< "$json")

  if [ "$result" -eq 0 ]; then
      echo "<failed> There is no function called '$test'"
      exit 1
  fi

  # Was it skipped?
  query="[.[] | select(.Package == \"getting-started-with-go-unit-testing/step5\" and .Test == \"$test\" and .Action == \"skip\")] | length"
  result=$(jq "$query" <<< "$json")

  if [ "$result" -eq 0 ]; then
      echo "<failed> You did not skip $test"
      exit 1
  fi
done

exit 0
