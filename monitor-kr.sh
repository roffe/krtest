#!/usr/bin/env bash
K=kubectl
set -ex
SVCIP=$($K get svc ws-server -o=jsonpath='{.spec.clusterIP}')
WSSERVERIP=$($K get po -l app=ws-server -o=jsonpath='{.items[0].status.hostIP}')
POD=$($K get po -l app=ws-server -o=jsonpath='{range .items[*]}{.status.podIP}{"\n"}{end}')
KRPOD=$($K get po -n kube-system -l k8s-app=kube-router -o=jsonpath='{range .items[*]}{.metadata.name} {.status.hostIP}{"\n"}{end}' | grep "${WSSERVERIP}")
#$K exec -ti -n kube-system $(echo $KRPOD | cut -d' ' -f1) -- bash
$K exec -ti -n kube-system $(echo $KRPOD | cut -d' ' -f1) -- bash -c "watch -n 1 bash -c 'ipvsadm | grep -A 3 ${SVCIP}'"