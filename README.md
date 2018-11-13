

go run customerServer/server.go 


go get github.com/golang/protobuf/proto

go get golang.org/x/net/context

go get google.golang.org/grpc

go get -u github.com/golang/protobuf/protoc-gen-go

cp $GOPATH/bin/protoc-gen-go /usr/local/bin/

./protoc --version

libprotoc 3.5.0

protoc -I customer/ customer/customer.proto --go_out=plugins=grpc:customer



go run customerServer/server.go -port=50051


echo "" | openssl s_client  -connect 54.255.15.251:50051


go run client.go -tls=true -server_addr=54.255.15.251:50051


#updated at 2018-09-11
#go run server.go

#go run client.go -tls=true -server_addr=localhost:50055



create private ssl test file:
openssl req -new -newkey rsa:2048 -nodes -keyout www.lijiuyang.com.key -out www.lijiuyang.com.csr
openssl x509 -req -days 365 -in www.lijiuyang.com.csr  -signkey www.lijiuyang.com.key -out www.lijiuyang.com.crt