# go-grpc-k8s

1)  docker build -t local/gcd -f Dockerfile.gcd .
2)  docker build -t local/api -f Dockerfile.api .
3) sudo kubectl create -f api.yaml
4) sudo kubectl create -f gcd.yaml
5) sudo minikube service api-service --url
6) curl http://[ip_addr]:[port]/gcd/294/462