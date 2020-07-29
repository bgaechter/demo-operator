# Demo Operator

This a example of an operator implemented with the [Operator SDK](https://github.com/operator-framework/getting-started/blob/master/README.md)

## Build
```bash
$ operator-sdk build docker.io/<user>/demo-operator:0.0.1
$ docker push docker.io/<user>/demo-operator:0.0.1
```

## Run locally
```bash
$ operator-sdk run --local --namespace=default
```