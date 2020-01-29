VERSION = $$(gobump show -r cmd)

export GO111MODULE=on

.PHONY: deps
deps:
	go get -u -d
	go mod tidy

.PHONY: devel-deps
devel-deps:
	go get -v \
	github.com/x-motemen/gobump/cmd/gobump \
	github.com/Songmu/ghch/cmd/ghch \
	github.com/Songmu/goxz/cmd/goxz \
	github.com/tcnksm/ghr \
	github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint: devel-deps
	golangci-lint run

.PHONY: build
build: deps
	go build -o tn

.PHONY: install
install: deps
	go install 

.PHONY: release
release: devel-deps
	echo ghr -n=v$(VERSION) -b="$($GOPATH/bin/ghch -F markdown --latest)" -draft v$(VERSION) ./dist/v$(VERSION)

.PHONY: crossbuild
crossbuild: devel-deps
	goxz -pv=$(VERSION) -os=linux -arch=amd64 -d=./dist/v$(VERSION)
