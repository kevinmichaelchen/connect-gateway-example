update:
	@docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf mod update proto

format:
	@docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf format -w

generate:
	@docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf generate

lint:
	@docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf lint
	@docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf breaking --against 'https://github.com/johanbrandhorst/connect-gateway-example.git#branch=master'

.PHONY: update format generate lint
