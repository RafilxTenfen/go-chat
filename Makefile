GO=go
GOBUILD=$(GO) build
BUILDENV=GOTRACEBACK=none CGO_ENABLED=0
GOENV=$(GO) env
FLAGS=-trimpath
LDFLAGS=-ldflags "-w -s"
CMD_PACKAGE=./cmd

# os/env information
ARCH=$(shell $(GOENV) | grep GOARCH | sed -E 's/GOARCH="(.*)"/\1/')
OS=$(shell $(GOENV) | grep GOOS | sed -E 's/GOOS="(.*)"/\1/')

# source files
SOURCES=go.mod $(shell find . -path ./cmd -prune -o -name "*.go" -print)

# platforms and targets
TARGETS=bot user chatapi
PLATFORMS=linux-amd64 linux-386 darwin-amd64 linux-arm7
PLATFORM_TARGETS = $(foreach platf,$(PLATFORMS),$(addprefix .build-cache/$(platf)/,$(TARGETS)))
DIST_TARGETS = $(foreach platf,$(PLATFORMS),.build-cache/rhizom-$(LASTTAG).$(platf).tar.gz)


all: build

dist: $(DIST_TARGETS)
	rm -rf dist
	mkdir -p dist
	mv .build-cache/*.tar.gz dist/

build: $(TARGETS)

%: .build-cache/$(OS)-$(ARCH)/%
	cp $< $@

.build-cache/linux-amd64/%: $(CMD_PACKAGE)/% $(SOURCES)
	env $(BUILDENV) GOARCH=amd64 GOOS=linux $(GOBUILD) $(FLAGS) $(LDFLAGS) -o $@ ./$<

.build-cache/linux-386/%: $(CMD_PACKAGE)/% $(SOURCES)
	env $(BUILDENV) GOARCH=386 GOOS=linux $(GOBUILD) $(FLAGS) $(LDFLAGS) -o $@ ./$<

.build-cache/darwin-amd64/%: $(CMD_PACKAGE)/% $(SOURCES)
	env $(BUILDENV) GOARCH=amd64 GOOS=darwin $(GOBUILD) $(FLAGS) $(LDFLAGS) -o $@ ./$<

.build-cache/linux-arm7/%: $(CMD_PACKAGE)/% $(SOURCES)
	env $(BUILDENV) GOARM=7 GOARCH=arm GOOS=linux $(GOBUILD) $(FLAGS) $(LDFLAGS) -o $@ ./$<


.build-cache/rhizom-$(LASTTAG).%.tar.gz: $(foreach target,$(TARGETS),.build-cache/%/$(target)) .build-cache/CHANGELOG.md
	cp -f README.md .build-cache/$*
	tar czf $@ -C .build-cache/$* .

check: .build-cache/testpkgs.list
	cat $< | xargs go test

.build-cache/testpkgs.list: $(SOURCES)
	mkdir -p .build-cache
	go list ./... > $@

db-up:
	@ docker-compose up --build -d

db-down:
	@ docker-compose down

clean:
	rm -rf $(TARGETS) $(PLATFORM_TARGETS)
	rm -rf dist build log

.PHONY: all dist build clean