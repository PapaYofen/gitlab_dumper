GO=go

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

VERSION := v2.1.0
BUILD := `git rev-parse --short HEAD`
TARGETS := gitlab_dumper
		
PROJECT=github.com/PapaYofen/gitlab_dumper
DOCKER_IMAGE_PRE=registry.example.com

all: check build

ifeq ($(GORACE), 1)
  BUILDFLAGS=-race
endif

build: $(TARGETS)

$(TARGETS): $(SRC)
	$(GO) build $(BUILDFLAGS) $(GOTAGS) $(PROJECT)/cmd/$@

image: $(TARGETS)
	docker build -t $(DOCKER_IMAGE_PRE)/gitlab_dumper:$(VERSION)-$(BUILD) .

.PHONY: clean all build check image

lint:
	@gometalinter --config=.gometalint ./...

packages = $(shell go list ./...|grep -v /vendor/ | grep -v /pb | grep -v /api | grep -v /test)
test: check
	$(GO) test ${packages}

cov: check
	gocov test $(packages) | gocov-html > coverage.html

check:
	@$(GO) tool vet ${SRC}

clean:
	rm -f $(TARGETS)