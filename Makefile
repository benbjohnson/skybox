PKG=./...
TEST=.
BENCH=.
COVERPROFILE=/tmp/c.out

bench:
	go test -v -test.bench=$(BENCH)

# http://cloc.sourceforge.net/
cloc:
	@cloc --not-match-f='Makefile|_test.go' .

cover: fmt
	go test -coverprofile=$(COVERPROFILE) -test.run=$(TEST) $(PKG)
	go tool cover -html=$(COVERPROFILE)
	rm $(COVERPROFILE)

fmt:
	@go fmt ./...

test: fmt
	@go test -v -test.run=$(TEST) ./...

.PHONY: bench cloc cover fmt test
