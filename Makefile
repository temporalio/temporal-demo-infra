.PHONY: bff ui
GITHUB_REPOSITORY = temporalio/temporal-demo-infra

deps:
	@go mod tidy
	@cd bff && go mod tidy
	@cd ui && npm install
	@cd provisioning_aws && npm install

domain:
	go run main.go
api:
	cd bff && go run cmd/bff/main.go
aws:
	cd provisioning_aws && npm start
ui:
	cd ui && npm run dev