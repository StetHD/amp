#!/bin/bash

amp  -s localhost service ps pinger_pinger 2>/dev/null | grep -o -w -E -q '[[:alnum:]]{25}'
