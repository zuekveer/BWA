-include .env
current_dir := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))

# Tools.
export TOOLS=$(current_dir)/tools
export TOOLS_BIN=$(TOOLS)/bin
export PATH := $(TOOLS_BIN):$(PATH)

.PHONY: build
build:
	@docker compose build

.PHONY: rebuild
rebuild:
	@docker compose down && docker compose up -d --build

.PHONY: run
run:
	@docker compose up -d

.PHONY: stop
stop:
	@docker compose down

.PHONY: restart
restart:
	docker compose down && docker compose up -d

.PHONY: app
app:
	docker compose exec app sh

.PHONY: clean
clean:
	@docker compose down --rmi all

.PHONY: logs-app
logs-app:
	@docker logs bwa-app-1

.PHONY: logs-db
logs-db:
	@docker logs bwa-db-1

.PHONY:
lint:
	$(TOOLS_BIN)/golangci-lint run

.PHONY:
migrate: install-tools
#	@envsubst < $(TOOLS_BIN)/goose -dir ./migrations postgres "$(DB_URI)" up -v
	$(TOOLS_BIN)/goose -dir ./migrations postgres "$(DB_URI)" up -v
.PHONY:
install-tools: export GOBIN=$(TOOLS_BIN)
install-tools:
	@mkdir -p $(TOOLS_BIN)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2
	go install github.com/pressly/goose/v3/cmd/goose@v3.10.0