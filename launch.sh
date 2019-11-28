#!/usr/bin/env sh

docker run -i --rm confita-bug-demo:v0.9.0 sh -c "PORT=1234 /demo-cmd --host=foo"