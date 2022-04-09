.PHONY: run

run:
	go run . -config config.json

test:
	go test -v ./pkg/...

.PHONY: deps-install

deps-install:
	@mkdir -p third_party && \
	cd third_party && \
	git clone https://github.com/envoyproxy/protoc-gen-validate --depth 1 && \
	git clone https://github.com/googleapis/googleapis.git --depth 1 && \
	echo "Done!"

.PHONY: deps-update

deps-update: ./third_party/*
	@for dir in $^; do \
		git -C $${dir} pull; \
	done

PROTO_FILE_LIST := $(shell find proto -type f -name '*.proto')

.PHONY: pb

pb:$(PROTO_FILE_LIST)
	@rm -rf gen temp/go
	@mkdir -p temp/go
	@protoc -Iproto -Ithird_party --go_out=plugins=grpc:temp/go --validate_out=lang=go:temp/go $^
	@mv temp/go/github.com/acrazing/universal-ingress-controller/gen/ gen
	@echo "Done!"
