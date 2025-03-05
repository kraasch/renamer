
test:
	go clean -testcache
	@#go test -tags testing ./pkg/... -run=.*/full-test_pipe-test_00
	go test -tags testing ./pkg/...
	go test -tags testing ./cmd/...
	@#go test -tags testing ./pkg/... -v

test_short:
	make test | grep -E '^(FAIL|ok)'

endtoend:
	go test ./e2e/... -v

run:
	go run ./cmd/renamer.go

.PHONY: build
build:
	rm -rf ./build/
	mkdir -p ./build/
	go build \
		-o ./build/renamer \
		-gcflags -m=2 \
		./cmd/ 

hub_update:
	@hub_ctrl ${HUB_MODE} ln "$(realpath ./build/renamer)"
