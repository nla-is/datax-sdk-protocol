.PHONY: generate generate-go

generate: generate-go

generate-go: v1/datax-sdk-protocol.pb.go

v1/datax-sdk-protocol.pb.go: v1/datax-sdk-protocol.proto
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		v1/datax-sdk-protocol.proto

