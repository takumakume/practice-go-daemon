INFO_COLOR=\033[1;34m
RESET=\033[0m
BOLD=\033[1m

deps: ## Install dependencies
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Installing Dependencies$(RESET)"
	go get -u github.com/golang/dep/...
	dep ensure

build: deps
	GOOS=linux GOARCH=amd64 go build -o takumakumed
