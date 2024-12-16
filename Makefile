BIN_FOLDER:=$(CURDIR)/bin

.PHONY: install-goose
install-goose:
	GOBIN=$(BIN_FOLDER) go install github.com/pressly/goose/v3/cmd/goose@v3.23.1

.PHONY: apply-migrations
apply-migrations:
	bin/goose -dir migrations postgres "postgresql://sso:sso@127.0.0.1:5432/sso?sslmode=disable" up

.PHONY: install-sqlc
install-sqlc:
	GOBIN=$(BIN_FOLDER) go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27.0