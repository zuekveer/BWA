.PHONY: build
build:
	@docker compose build

.PHONY: run
run:
	@docker compose up -d

.PHONY: stop
stop:
	@docker compose down

.PHONY: clean
clean:
	@docker compose down --rmi all

.PHONY: logs-app
logs-app:
	@docker logs app-calendar

.PHONY: logs-pg
logs-pg:
	@docker logs pg-16
