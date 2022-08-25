define gendoc
	@swag init -g ./router/document.go -d ./app -o ./assets/docs --md ./wiki --parseDependency --parseDepth 100
endef


.PHONY: proto
proto:
	@cd grpc; \
	rm -rf pb; \
 	mkdir pb; \
 	protoc --proto_path=./proto --go_out=module=github.com/chatpuppy/puppychat/grpc/pb:./pb --go-grpc_out=module=github.com/chatpuppy/puppychat/grpc/pb:./pb proto/*.proto; \
 	echo "Regenerate proto done."

.PHONY: doc
doc:
	$(call gendoc)