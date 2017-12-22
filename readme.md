git clone https://bitbucket.org/scpman/grpc-customer.git

docker-compose up -d

docker exec -it $name /bin/bash

apt-get update && apt-get install git -y

cd /mnt/app/src

go run customerServer/server.go 


go get github.com/golang/protobuf/proto

go get golang.org/x/net/context

go get google.golang.org/grpc

go get -u github.com/golang/protobuf/protoc-gen-go

cp $GOPATH/bin/protoc-gen-go /usr/local/bin/

./protoc --version
libprotoc 3.5.0
protoc -I customer/ customer/customer.proto --go_out=plugins=grpc:customer

https://app.yinxiang.com/Home.action#n=27dfd813-689c-47d1-9ab2-a798f8bc2500&s=s3&ses=4&sh=2&sds=5&

go run customerServer/server.go -port=50051


echo "" | openssl s_client  -connect 54.255.15.251:50051


go run client.go -tls=true -server_addr=54.255.15.251:50051



