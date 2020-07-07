.PHONY: docs
docs:
	@docker run -v $$PWD/:/docs pandoc/latex -f markdown /docs/README.md -o /docs/README.pdf

tests:
	docker-compose down
	docker-compose down
	docker-compose up --build

test-locally:
	go clean -testcache . && ACCOUNT_API_HOST="localhost" ACCOUNT_API_PORT="8080" CGO_ENABLED=0 go test --cover ./...

coverage:
	go test -coverprofile accountservice.out ./... && go tool cover -html=accountservice.out -o accountservicecoverage.html