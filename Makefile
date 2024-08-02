SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help

.PHONY: run db

run:
	@go run .

db:
	@docker compose up -d
