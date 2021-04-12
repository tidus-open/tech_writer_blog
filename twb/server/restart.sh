#!/bin/bash

killall -9 server
nohup ./server  > twb.log 2>&1 &
