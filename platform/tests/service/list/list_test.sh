#!/bin/bash

SECONDS=0
amp="amp -s localhost"

test_stack_deploy() {
  $amp stack up -c platform/tests/service/list/global.service.yml global
  $amp stack up -c platform/tests/service/list/replicated.service.yml replicated
}

test_service_starting() {
  $amp service ls 2>/dev/null | pcregrep -q "\s*global\s*0/1\s*[0-9]\s*STARTING\s*subfuzion/pinger\s*latest\s*"
  $amp service ls 2>/dev/null | pcregrep -q "\s*replicated\s*0/2\s*[0-9]\s*STARTING\s*subfuzion/pinger\s*latest\s*"
}

test_service_global_running() {
  while true
  do
     if $amp service ls 2>/dev/null | pcregrep -q "\s*global_pinger\s*global\s*1/1\s*[0-9]\s*RUNNING\s*" || [ $SECONDS -eq 5 ]
     then
             break
     fi
     sleep 1
     SECONDS=$[$SECONDS+1]
  done
}

test_service_replicated_running() {
  while true
  do
     if $amp service ls 2>/dev/null | pcregrep -q "\s*replicated_pinger\s*replicated\s*2/2\s*[0-9]\s*RUNNING\s*" || [ $SECONDS -eq 5 ]
     then
             break
     fi
     sleep 1
     SECONDS=$[$SECONDS+1]
  done
}
