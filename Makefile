
test:
	go clean -testcache
	go test ./...
	@#go test -v ./...

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

