TLS

go run customerServer/server.go -tls=true --port=55555

go run customerClient/client.go -tls=true -server_addr=localhost:55555 -ca_file=/home/gavin/desktop/git_club/pumpkin/src/testdata/qraved_com.crt -server_host_override=www.qraved.com

Without TLS?
go run customerServer/server.go -tls=false --port=55555
go run customerClient/client.go -tls=false -server_addr=localhost:55555 