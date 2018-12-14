#!/usr/bin/env bash
K=kubectl
set -ex
$K run --rm -ti ping-client --image roffe/ws-ping-pong:latest --command -- /go/bin/client -addr ws-server