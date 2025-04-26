# Makefile


user:
	protoc -I. \                                           
    -I/../.. \
    -I$(go env GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v1.2.1 \ 
    ./internal/iam/infra/grpc/user/user.proto \
    --go_out=./internal/iam/infra/grpc/user --go_opt=paths=source_relative \
    --go-grpc_out=./internal/iam/infra/grpc/user --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=./internal/iam/infra/grpc/user --grpc-gateway_opt=paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    --validate_out=lang=go:./internal/iam/infra/grpc/user


