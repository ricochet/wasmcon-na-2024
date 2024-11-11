#!/bin/bash

mkdir -p /tmp/passwords/
curl -L -O --output-dir /tmp/passwords https://raw.githubusercontent.com/danielmiessler/SecLists/refs/heads/master/Passwords/Common-Credentials/500-worst-passwords.txt
