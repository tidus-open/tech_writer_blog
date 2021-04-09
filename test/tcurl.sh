#!/bin/bash
curl -H "Content-Type: application/json" -d '{"user_name":"delphi", "passwd":"12345"}'  --noproxy  "*" https://127.0.0.1:9090/v1/accounts -k

