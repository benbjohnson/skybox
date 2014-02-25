PKG=./...
TEST=.
BENCH=.
COVERPROFILE=/tmp/c.out

assets:
	cd server && go-bindata -pkg=server -prefix=assets/ assets

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
	@go fmt $(PKG)

run: assets templates
	@go run ./cmd/skybox/main.go

templates:
	@ego server/template

test: assets templates fmt
	@go test -v -test.run=$(TEST) $(PKG)

.PHONY: assets bench cloc cover fmt run templates test
