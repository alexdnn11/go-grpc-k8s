# go-grpc-k8s

1) protoc -I pb/ pb/gcd.proto --go_out=plugins=grpc:pb
2) docker build -t local/gcd -f Dockerfile.gcd .
3) docker build -t local/api -f Dockerfile.api .
4) sudo kubectl create -f api.yaml
5) sudo kubectl create -f gcd.yaml
6) sudo minikube service api-service --url
7) curl http://[ip_addr]:[port]/gcd/294/462