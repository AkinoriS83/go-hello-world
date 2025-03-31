#!/bin/bash
cd /home/ubuntu/myapp
nohup ./myapp > app.log 2>&1 &
