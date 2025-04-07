#!/bin/bash
pkill myapp || true
[ -f /home/ubuntu/myapp/myapp ] && rm /home/ubuntu/myapp/myapp