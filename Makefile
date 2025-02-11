PROTO_SRC=./proto
PROTO_GEN=./app/book/client/book

.PHONY: proto
proto: $(PROTO_SRC)/*.proto
	#mkdir ${PROTO_GEN}
	protoc --proto_path=$(PROTO_SRC) \
	       --go_out=${PROTO_GEN} --go_opt=paths=source_relative \
	       --go-grpc_out=${PROTO_GEN} --go-grpc_opt=paths=source_relative \
	       $(PROTO_SRC)/*.proto
tidy:
	go mod tidy

run:
	go run .