PROTOFILES=$(wildcard *.proto)
GOPROTOFILES=$(patsubst %.proto,%.pb.go,$(PROTOFILES))
GOGRPCPROTOFILES=$(patsubst %.proto,%_grpc.pb.go,$(PROTOFILES))

$(GOPROTOFILES) $(GOGRPCPROTOFILES) : $(PROTOFILES)
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $^