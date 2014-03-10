PKG=./...
TEST=.
BENCH=.
COVERPROFILE=/tmp/c.out

default: build

assets: less
	@cd server && go-bindata -pkg=server -prefix=assets/ assets

bench:
	go test -v -test.bench=$(BENCH)

# required: http://dave.cheney.net/2012/09/08/an-introduction-to-cross-compilation-with-go
build: assets templates
	mkdir -p build
	cd cmd/skybox && goxc -c=.goxc.json -d ../../build

# http://cloc.sourceforge.net/
cloc:
	@cloc --not-match-f='Makefile|_test.go' .

cover: fmt
	go test -coverprofile=$(COVERPROFILE) -test.run=$(TEST) $(PKG)
	go tool cover -html=$(COVERPROFILE)
	rm $(COVERPROFILE)

fmt:
	@go fmt $(PKG)

less:
	@lessc server/assets/application.less > server/assets/application.css

run: assets templates
	go run ./cmd/skybox/main.go --data-dir=/tmp/skybox

templates:
	@ego server/template

test: assets templates fmt
	@go test -a -v -test.run=$(TEST) $(PKG)

.PHONY: assets bench cloc cover fmt less run templates test
