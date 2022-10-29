.PHONY: build

build:
	@docker build --no-cache -t web_algo:1.1 .

run:
	@echo "=========== App is startin up ===="
	@echo "CONFIG ==== 😁 Exporting environemnt variables"
	@echo "SUCCESS ===  ✔ Environment variables exported"
	@echo "INIT ======  ⚡ Running server"
	@go run main.go