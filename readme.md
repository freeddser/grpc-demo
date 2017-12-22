git clone https://bitbucket.org/scpman/grpc-customer.git
docker-compose up -d
docker exec -it $name /bin/bash
apt-get update && apt-get install git -y
cd /mnt/app/src
go run customerServer/server.go 

go get github.com/golang/protobuf/proto
go get golang.org/x/net/context
go get google.golang.org/grpc


go run customerServer/server.go -port=50051

echo "" | openssl s_client  -connect 54.255.15.251:50051

go run client.go -tls=true -server_addr=54.255.15.251:50051



