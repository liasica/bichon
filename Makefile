
.PHONY: proto
proto:
	@cd grpc; \
	rm -rf pb; \
 	mkdir pb; \
 	protoc --proto_path=./proto --go_out=module=github.com/chatpuppy/puppychat/grpc/pb:./pb --go-grpc_out=module=github.com/chatpuppy/puppychat/grpc/pb:./pb proto/*.proto; \
 	echo "Regenerate proto done."