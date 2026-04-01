-include .makerc
.DEFAULT_GOAL:=help

# --- Config ------------------------------------------------------------------

# Newline hack for error output
define br


endef

grafanactl.yaml:
	@echo "NOTE: Please add a settings yaml"
	@exit 1

# --- Targets -----------------------------------------------------------------

# This allows us to accept extra arguments
%: .mise .lefthook
	@:

.PHONY: .mise
# Install dependencies
.mise: msg := $(br)$(br)Please ensure you have 'mise' installed and activated!$(br)$(br)$$ brew update$(br)$$ brew install mise$(br)$(br)See the documentation: https://mise.jdx.dev/getting-started.html$(br)$(br)
.mise:
ifeq (, $(shell command -v mise))
	$(error ${msg})
endif
	@mise install

# Configure git hooks for lefthook
.lefthook:
	@lefthook install

### Tasks

.PHONY: check
## Run lint & tests
check: tidy lint test

.PHONY: tidy
## Run go mod tidy
tidy:
	@echo "〉go mod tidy"
	@go mod tidy

.PHONY: lint
## Run linter
lint:
	@echo "〉golangci-lint run"
	@golangci-lint run

.PHONY: lint.fix
## Fix lint violations
lint.fix:
	@echo "〉golangci-lint run fix"
	@golangci-lint run --fix

.PHONY: test
## Run tests
test:
	@echo "〉go test"
	@# see https://github.com/pterm/pterm/issues/482
	@GO_TEST_TAGS=-skip go test -tags=safe -coverprofile=coverage.out
	@#GO_TEST_TAGS=-skip go test -tags=safe -coverprofile=coverage.out -race

.PHONY: outdated
## Show outdated direct dependencies
outdated:
	@echo "〉go mod outdated"
	@go list -u -m -json all | go-mod-outdated -update -direct

.PHONY: install
## Install binary
install:
	@echo "〉installing ${GOPATH}/bin/squadron"
	@go build -tags=safe -o ${GOPATH}/bin/squadron cmd/main.go

.PHONY: build
## Build binary
build:
	@mkdir -p bin
	@echo "〉building bin/squadron"
	@go build -tags=safe -o bin/squadron cmd/main.go

.PHONY: .context
.context: grafanactl.yaml
ifndef CONTEXT
	$(error $(br)$(br)CONTEXT variable is required.$(br)Usage: make [task] CONTEXT=foo$(br)$(br))
endif
	@grafanactl --config grafanactl.yaml config use-context ${CONTEXT}

.PHONY: .resource
.resource:
ifndef RESOURCE
	$(error $(br)$(br)RESOURCE variable is required.$(br)Usage: make [task] RESOURCE=foo$(br)$(br))
endif

.PHONY: dashboards
## Serve through grizzly
serve: .context .resource
	@grafanactl --config grafanactl.yaml resources serve ./${RESOURCE}

.PHONY: dashboards
## Serve through grizzly
foo: .context .resource
	@grafanactl --config grafanactl.yaml resources serve --script "go run . generate --raw" --watch .

.PHONY: dashboards
## Serve through grizzly
list: .context
	@grafanactl --config grafanactl.yaml resources list

### Utils

.PHONY: help
## Show help text
help:
	@echo "\033[1;36mGrafana Dashboards\033[0m"
	@awk '{ \
		if($$0 ~ /^### /){ \
			if(help) printf "%-23s %s\n\n", cmd, help; help=""; \
			printf "\n%s:\n", substr($$0,5); \
		} else if($$0 ~ /^[a-zA-Z0-9._-]+:/){ \
			cmd = substr($$0, 1, index($$0, ":")-1); \
			if(help) printf "  %-23s %s\n", cmd, help; help=""; \
		} else if($$0 ~ /^##/){ \
			help = help ? help "\n                        " substr($$0,3) : substr($$0,3); \
		} else if(help){ \
			print "\n                        " help "\n"; help=""; \
		} \
	}' $(MAKEFILE_LIST)
