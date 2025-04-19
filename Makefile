# Makefile


user:
	protoc -I. -I./internal/iam/infra/grpc/user --go_out=./internal/iam/infra/grpc/user --go_opt=paths=source_relative \
	--go-grpc_out=./internal/iam/infra/grpc/user --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=./internal/iam/infra/grpc/user --grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	./internal/iam/infra/grpc/user/user.proto

profile:
	cd proto && protoc -I. -I./profile --go_out=../protogen/golang --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=../protogen/golang --grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	./profile/*.proto

	

auth:
	cd proto && protoc -Iproto --go_out=../protogen/golang --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=../protogen/golang --grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	./auth/*.proto