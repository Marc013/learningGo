#!/bin/bash

TEST_FILE="step4/calculator_test.go"

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

tests=()
tests+=("TestCalculator/Add" "Add" "2" "3")
tests+=("TestCalculator/Subtract" "Subtract" "5" "3")
tests+=("TestCalculator/Multiply" "Multiply" "2" "3")
tests+=("TestCalculator/DivideValid" "Divide" "6" "3")
tests+=("TestCalculator/DivideByZero" "Divide" "6" "0")

index=0
col_size=4
row_size=$(("${#tests[@]}/$col_size"))

for test in $(seq 1 $row_size)
do
  name=${tests[($index*$col_size)]}
  fn=${tests[($index*$col_size+1)]}
  a=${tests[($index*$col_size+2)]}
  b=${tests[($index*$col_size+3)]}

  ((index++))

  # Does the test exist?
  query="[.[] | select(.Package == \"getting-started-with-go-unit-testing/step4\" and .Test == \"$name\")] | length"
  result=$(jq "$query" <<< "$json")

  if [ "$result" -eq 0 ]; then
      echo "<failed> There is no subtest called '$name'"
      exit 1
  fi

  # Was it called with the correct parameters?
  query="[.[] | select(.Package == \"getting-started-with-go-unit-testing/step4\" and .Test == \"$name\" and .Action == \"output\" and (.Output | contains(\"$fn called with parameters $a and $b\")))] | length"
  result=$(jq "$query" <<< "$json")

  if [ "$result" -eq 0 ]; then
      echo "<failed> You did not pass the correct parameters to $fn"
      exit 1
  fi

  # Does the test pass?
  query="[.[] | select(.Package == \"getting-started-with-go-unit-testing/step4\" and .Test == \"$name\" and .Action == \"pass\")] | length"
  result=$(jq "$query" <<< "$json")

  if [ "$result" -eq 0 ]; then
      echo "<failed> The test '$test' failed"
      exit 1
  fi
done

exit 0
