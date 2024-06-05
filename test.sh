#!/bin/bash
fswatch -l 1 -o ./**/*.go --event=Updated | while read -r; do
  clear
  output=$(go test ./... 2>&1)
  if echo "$output" | grep -- '- FAIL' > /dev/null; then
    echo -e "\e[31m$(echo "$output" | awk '/--- FAIL/,/Test:/')\e[0m"
  elif [ -z "$output" ]; then
    echo -e "\e[31mBUILD FAILED\e[0m"
  else
    echo -e "\e[32mPASS\e[0m"
  fi
done