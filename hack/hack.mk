include ./hack/hack-cli.mk

.PHONY: start
start: cli.install
	@set -e; \
	if [ "$(filter gateway,$(MAKECMDGOALS))" ]; then \
		$(MAKE) start-gateway; \
	elif [ "$(filter service,$(MAKECMDGOALS))" ]; then \
		$(MAKE) start-service; \
	else \
		echo "Unknown service command. Use 'make help' for available commands."; \
	fi
doc:
	swag init --parseDependency -g ./internal/api_gateway/main.go

.PHONY: start-gateway
start-gateway:
	@set -e; \
	go run ./internal/api_gateway/main.go start

.PHONY: run-cmd
run-cmd:
	@set -e; \
	go run ./internal/$(name)_service/main.go $(cmd)

.PHONY: start-service
start-service:
	@set -e; \
	go run ./internal/$(name)_service/main.go start

# Build docker image.
.PHONY: image
image: cli.install
	$(eval _TAG  = $(shell git describe --dirty --always --tags --abbrev=8 --match 'v*' | sed 's/-/./2' | sed 's/-/./2'))
ifneq (, $(shell git status --porcelain 2>/dev/null))
	$(eval _TAG  = $(_TAG).dirty)
endif
	$(eval _TAG  = $(if ${TAG},  ${TAG}, $(_TAG)))
	$(eval _PUSH = $(if ${PUSH}, ${PUSH}, ))
	@docker ${_PUSH} -tn $(DOCKER_NAME):${_TAG};


# Build docker image and automatically push to docker repo.
.PHONY: image.push
image.push:
	@make image PUSH=-p;


# Deploy image and yaml to current kubectl environment.
.PHONY: deploy
deploy:
	$(eval _TAG = $(if ${TAG},  ${TAG}, develop))

	@set -e; \
	mkdir -p $(ROOT_DIR)/temp/kustomize;\
	cd $(ROOT_DIR)/manifest/deploy/kustomize/overlays/${_ENV};\
	kustomize build > $(ROOT_DIR)/temp/kustomize.yaml;\
	kubectl   apply -f $(ROOT_DIR)/temp/kustomize.yaml; \
	if [ $(DEPLOY_NAME) != "" ]; then \
		kubectl   patch -n $(NAMESPACE) deployment/$(DEPLOY_NAME) -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"$(shell date +%s)\"}}}}}"; \
	fi;

.PHONY: db
db:
	@if [ "$(filter create:migrate,$(MAKECMDGOALS))" ]; then \
		$(MAKE) db-create-migrate; \
	elif [ "$(filter run:migrate,$(MAKECMDGOALS))" ]; then \
		$(MAKE) db-run-migrate; \
	elif [ "$(filter run:seed,$(MAKECMDGOALS))" ]; then \
		$(MAKE) db-run-seed; \
	elif [ "$(filter run:init,$(MAKECMDGOALS))" ]; then \
		$(MAKE) db-run-init; \
	elif [ "$(filter run:rollback,$(MAKECMDGOALS))" ]; then \
		$(MAKE) db-run-rollback; \
	elif [ "$(filter run:status,$(MAKECMDGOALS))" ]; then \
		$(MAKE) db-run-status; \
	else \
		echo "Unknown db command. Use 'make help' for available commands."; \
	fi

db-create-migrate:
	@if [ -z "$(name)" ]; then \
		read -p "Enter migration name: " name; \
	fi; \
	mkdir -p ./internal/$(service)_service/infra/db/migrations/sql; \
	MIGRATION_DIR=./internal/$(service)_service/infra/db/migrations/sql; \
	UP_FILE=$$MIGRATION_DIR/$(TIMESTAMP)-$$name.up.sql; \
	DOWN_FILE=$$MIGRATION_DIR/$(TIMESTAMP)-$$name.down.sql; \
	echo "Creating migration files..."; \
	touch $$UP_FILE $$DOWN_FILE; \
	echo "-- Migration up script for $$name" > $$UP_FILE; \
	echo "-- Migration down script for $$name" > $$DOWN_FILE; \
	echo "Created $$UP_FILE and $$DOWN_FILE"
db-run-migrate:
	@set -e; \
	go run ./internal/$(service)_service/main.go db migrate
db-run-seed:
	@set -e; \
	go run ./internal/$(service)_service/main.go db seed
db-run-init:
	@set -e; \
	go run ./internal/$(service)_service/main.go db init
db-run-rollback:
	@set -e; \
	go run ./internal/$(service)_service/main.go db rollback
db-run-status:
	@set -e; \
	go run ./internal/$(service)_service/main.go db status