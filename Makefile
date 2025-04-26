# Makefile


user:
	cd ./internal/iam/infra/grpc/user/ && protoc -I. -I/../.. \
    -I./../../../../../proto \
    -I/home/user/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v1.2.1 \
    ./user.proto \
    --go_out=./ --go_opt=paths=source_relative \
    --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=./ --grpc-gateway_opt=paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    --validate_out=lang=go:./


