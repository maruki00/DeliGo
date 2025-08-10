# Makefile
# Proto files.
#
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

role:
	cd ./internal/iam/infra/grpc/role/ && protoc -I. -I/../.. \
    -I./../../../../../proto \
    -I/home/user/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v1.2.1 \
    ./role.proto \
    --go_out=./ --go_opt=paths=source_relative \
    --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=./ --grpc-gateway_opt=paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    --validate_out=lang=go:./

profile:
	cd ./internal/profile/infra/grpc/profile/ && protoc -I. -I/../.. \
    -I./../../../../../proto \
    -I/home/user/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v1.2.1 \
    ./profile.proto \
    --go_out=./ --go_opt=paths=source_relative \
    --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=./ --grpc-gateway_opt=paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    --validate_out=lang=go:./

product:
	cd ./internal/product/infra/grpc/product/ && protoc -I. -I/../.. \
    -I./../../../../../proto \
    -I/home/user/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v1.2.1 \
    ./product.proto \
    --go_out=./ --go_opt=paths=source_relative \
    --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=./ --grpc-gateway_opt=paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    --validate_out=lang=go:./

shop:
	cd ./internal/shop/infra/grpc/shop/ && protoc -I. -I/../.. \
    -I./../../../../../proto \
    -I/home/user/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v1.2.1 \
    ./shop.proto \
    --go_out=./ --go_opt=paths=source_relative \
    --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=./ --grpc-gateway_opt=paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    --validate_out=lang=go:./

menu:
	cd ./internal/shop/infra/grpc/shop/ && protoc -I. -I/../.. \
    -I./../../../../../proto \
    -I/home/user/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v1.2.1 \
    ./shop.proto \
    --go_out=./ --go_opt=paths=source_relative \
    --go-grpc_out=./ --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=./ --grpc-gateway_opt=paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    --validate_out=lang=go:./

show:
	@echo "make [target]"
	@echo "targets:"
	@echo " user"
	@echo " profile"
	@echo " product"
	@echo " shop"
	@echo " role"
	@echo " menu"

help:
	@echo "make [target]"
	@echo "targets:"
	@echo " user"
	@echo " profile"
	@echo " product"
	@echo " shop"
	@echo " role"
	@echo " menu"
