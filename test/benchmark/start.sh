#!/bin/bash

killall -9 twbclient

nohup ./twbclient >> twbtest.log 2>&1 &
