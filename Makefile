
.PHONY: run-api, create-api, shutdown, test test-external-provider test-application

run-api:
	docker compose up -d mysql && sleep 4 && docker compose up -d currency-api

create-api:
	docker compose build --no-cache --pull

shutdown:
	docker compose down

# Regra para rodar todos os testes
test: test-external-provider test-application

# Regra para rodar os testes no pacote external
test-external-provider:
	go test ./pkg/external/... -v

# Regra para rodar os testes no pacote application
test-application:
	go test ./application/... -v