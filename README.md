install grpc 
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc


path setup

PATH="${PATH}:${HOME}/go/bin"
{HOME}/go/bin/protoc-gen-go


run

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/blog.proto

