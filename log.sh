#!/usr/bin/env bash
K=kubectl
set -ex
SVCIP=$($K get svc ws-server -o=jsonpath='{.spec.clusterIP}')
WSSERVERIP=$($K get po -l app=ws-server -o=jsonpath='{.items[0].status.hostIP}')
POD=$($K get po -n kube-system -l k8s-app=kube-router -o=jsonpath='{range .items[*]}{.metadata.name} {.status.hostIP}{"\n"}{end}' | grep "${WSSERVERIP}")

$K logs -f -n kube-system $(echo $POD | cut -d' ' -f1) --since=30s