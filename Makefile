.PHONY: generate generate-go

generate: generate-go

generate-go: datax-sdk-protocol.pb.go

datax-sdk-protocol.pb.go: datax-sdk-protocol.proto
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		datax-sdk-protocol.proto

