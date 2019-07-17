### go-grpc-k8s

* Install [protobuf compiler](https://github.com/google/protobuf/blob/master/README.md#protocol-compiler-installation)

* Install the protoc GO plugin

   ```
   $ go get -u github.com/golang/protobuf/protoc-gen-go

* make generate

* make up

* sudo kubectl create -f api.yaml
* sudo kubectl create -f gcd.yaml
* sudo minikube service api-service --url
* Postman collections: https://www.getpostman.com/collections/a69e637db692ce39cbd2
