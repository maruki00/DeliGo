# Makefile

protoc:
	cd protos && protoc --go_out=./protogen --go_opt=paths=source_relative ./product.proto
