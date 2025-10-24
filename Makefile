-include .makerc
.DEFAULT_GOAL:=help

# --- Config ------------------------------------------------------------------

# Newline hack for error output
define br


endef

settings.yaml:
	@echo "NOTE: Please add a settings yaml"
	@exit 1

# --- Targets -----------------------------------------------------------------

# This allows us to accept extra arguments
%: .mise
	@:

.PHONY: .mise .lefthook
# Install dependencies
.mise: msg := $(br)$(br)Please ensure you have 'mise' installed and activated!$(br)$(br)$$ brew update$(br)$$ brew install mise$(br)$(br)See the documentation: https://mise.jdx.dev/getting-started.html$(br)$(br)
.mise:
ifeq (, $(shell command -v mise))
	$(error ${msg})
endif
	@mise install

.PHONY: .lefthook
# Configure git hooks for lefthook
.lefthook:
	@lefthook install

### Tasks

.PHONY: foomo.gotsrpc
## Serve through grizzly
foomo.gotsrpc: settings.yaml
	@grr serve --disable-reporting --watch --open-browser --only-spec --kind Dashboard dashboards/foomo/gotsrpc.json

.PHONY: foomo.http-server
## Serve through grizzly
foomo.http-server: settings.yaml
	@grr serve --disable-reporting --watch --open-browser --only-spec --kind Dashboard dashboards/foomo/http-server.json

.PHONY: foomo.squadron-releases
## Serve through grizzly
foomo.squadron-releases: settings.yaml
	@grr serve --disable-reporting --watch --open-browser --only-spec --kind Dashboard dashboards/foomo/squadron-releases.json

.PHONY: nodejs.resources
## Serve through grizzly
nodejs.resources: settings.yaml
	@grr serve --disable-reporting --watch --open-browser --only-spec --kind Dashboard dashboards/nodejs/resources.json

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
