
# Install/Update to the latest CLI tool.
.PHONY: cli
cli:
	@set -e; \
	go work sync
	echo "App is up-to-date";


# Check and install CLI tool.
.PHONY: cli.install
cli.install:
	@set -e; \
	make cli;