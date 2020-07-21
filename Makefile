.PHONY: test cover check

test:
	@go test -v ./downloader

check:
	@revive

cover:
	@go test -coverprofile=coverage.out
	@go tool cover -html=coverage.out