#!/bin/bash
for i in `seq 10`; do
  echo EOF | bm_CLI
  if [ "$?" == "0" ]; then
    echo "bm_CLI is ready!"
    exit 0
  fi
  sleep 1
done
echo "bm_CLI is not ready"
exit 1
