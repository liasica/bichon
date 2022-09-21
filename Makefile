define gendoc
	swag init -g ./router/document.go -d ./app -o ./assets/docs --md ./wiki --parseDependency --parseDepth 100
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
	@$(call gendoc)


KEYCODE ?= $(shell bash -c 'read -p "Key File: " path; echo `cat $$path`')

define build
	@$(call gendoc); \
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -trimpath -ldflags "-X github.com/chatpuppy/puppychat/cmd/puppychat/command.NodeSecret=$(1)" -o build/release/puppychat cmd/puppychat/main.go
endef

.PHONY: build
build:
	$(call build,$(KEYCODE))


.PHONY: next
next:
	@$(call build,$(KEYCODE))
	@sh next.sh

