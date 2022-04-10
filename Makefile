.PHONY: run

run:
	go run . -config config.json

test:
	go test -v ./pkg/...

.PHONY: pb

pb:
	@set -e; \
	rm -rf temp gen; \
	buf generate api; \
	mv temp/go/github.com/acrazing/universal-ingress-controller/gen gen; \
	rm -rf temp; \
	echo "Done!"

pb-vendor-root:
	@set -e; \
	rm -rf third_party; \
	mkdir -p third_party; \
	ln -sf $$(find $(HOME)/.cache/buf/v1/module/data -type f -name 'buf.yaml' | xargs -I % bash -c 'find $$(dirname %)/* -maxdepth 0 -type d') $(PWD)/third_party
