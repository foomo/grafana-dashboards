-include .makerc
.DEFAULT_GOAL:=help

settings.yaml:
	@echo "NOTE: Please add a settings yaml"
	@exit 1

# --- Targets -----------------------------------------------------------------

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

## === Utils ===

.PHONY: help
## Show help text
help:
	@awk '{ \
		if ($$0 ~ /^.PHONY: [a-zA-Z\-\_0-9]+$$/) { \
			helpCommand = substr($$0, index($$0, ":") + 2); \
			if (helpMessage) { \
				printf "\033[36m%-23s\033[0m %s\n", \
					helpCommand, helpMessage; \
				helpMessage = ""; \
			} \
		} else if ($$0 ~ /^[a-zA-Z\-\_0-9.]+:/) { \
			helpCommand = substr($$0, 0, index($$0, ":")); \
			if (helpMessage) { \
				printf "\033[36m%-23s\033[0m %s\n", \
					helpCommand, helpMessage"\n"; \
				helpMessage = ""; \
			} \
		} else if ($$0 ~ /^##/) { \
			if (helpMessage) { \
				helpMessage = helpMessage"\n                        "substr($$0, 3); \
			} else { \
				helpMessage = substr($$0, 3); \
			} \
		} else { \
			if (helpMessage) { \
				print "\n                        "helpMessage"\n" \
			} \
			helpMessage = ""; \
		} \
	}' \
	$(MAKEFILE_LIST)
