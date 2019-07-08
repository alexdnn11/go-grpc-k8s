# go-grpc-k8s

1. Install [protobuf compiler](https://github.com/google/protobuf/blob/master/README.md#protocol-compiler-installation)

2. Install the protoc Go plugin

   ```
   $ go get -u github.com/golang/protobuf/protoc-gen-go

3. make generate

4. make up

5. sudo kubectl create -f api.yaml
6. sudo kubectl create -f gcd.yaml
7. sudo minikube service api-service --url
8. curl http://[ip_addr]:[port]/gcd/a/b

9. Postman collections: https://www.getpostman.com/collections/a69e637db692ce39cbd2