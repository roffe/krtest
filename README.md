# krtest

## Deploy server

    kubectl apply -f https://github.com/roffe/krtest/blob/master/templates/deploy.yml

## Run test client

    kubectl run --rm -ti krtest-client --image roffe/krtest:latest --command -- /go/bin/client -addr krtest.krtest.svc.cluster.local.