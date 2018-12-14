# krtest

## Deploy server

    kubectl apply -f https://github.com/roffe/krtest/blob/master/templates/deploy.yml

## Run test client

    kubectl run --rm -ti ping-client --image roffe/ws-ping-pong:latest --command -- /go/bin/client -addr krtest.krtest.svc.cluster.local.