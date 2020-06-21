#!/bin/bash
set -x
docker build . -f Dockerfile -t hjames/simpler:latest
