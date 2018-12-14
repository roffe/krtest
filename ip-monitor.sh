#!/usr/bin/env bash
K=kubectl
set -ex
KRPOD=$($K get po -n kube-system -l k8s-app=kube-router -o=jsonpath='{range .items[*]}{.metadata.name} {.status.hostIP}{"\n"}{end}' | grep "${WSSERVERIP}")
$K exec -ti -n kube-system $(echo $KRPOD | cut -d' ' -f1) -- ip monitor