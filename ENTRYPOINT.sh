#!/bin/bash

export DISPLAY=:0
Xvfb :0 -screen 0 1920x1080x24 &
/bin/app