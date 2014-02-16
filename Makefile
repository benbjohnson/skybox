PKG=./...
TEST=.
BENCH=.
COVERPROFILE=/tmp/c.out

bench:
	go test -v -test.bench=$(BENCH)

cover: fmt
	go test -coverprofile=$(COVERPROFILE) -test.run=$(TEST) $(PKG)
	go tool cover -html=$(COVERPROFILE)
	rm $(COVERPROFILE)

fmt:
	@go fmt ./...

test: fmt
	@go test -v -cover -test.run=$(TEST) ./...

.PHONY: bench cover fmt test
