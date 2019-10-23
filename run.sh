#!/bin/bash

docker run --name url2img --rm --net=host -it -e DISPLAY=:0 yale8848/cutycapt-docker:v5_5