#
# proxrm
#
# Makefile
#

include .env

GOLANG := go

GOCYCLO := gocyclo
GOFMT := gofmt
GOLANGCI_LINT := golangci-lint
GOLINT := golint
GOSEC := gosec
GOVULNCHECK := govulncheck
INEFFASSIGN := ineffassign
STATICCHECK := staticcheck

.PHONY: all install build test iterate get format vet tidy download verify lint staticcheck cyclo ineffassign vulncheck clean

all: build

iterate: install get format tidy download verify lint golangci-lint vet staticcheck cyclo ineffassign gosec vulncheck build test

install:
	@echo " 📦 staticcheck"
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@echo " 📦 cyclo"
	@go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	@echo " 📦 ineffassign"
	@go install github.com/gordonklaus/ineffassign@latest
	@echo " 📦 vulncheck"
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@echo " 📦 golint"
	@go install golang.org/x/lint/golint@latest
	@echo " 📦 gosec"
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	@echo " 📦 golintci-lint"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

build:
	@echo " 🚧 building proxrm"
	@$(GOLANG) build -v .
	@sha256sum proxrm > checksum.txt

test:
	@echo " 🚧 testing proxrm"
	@$(GOLANG) test -v .

get:
	@echo " 🔘 Go Get"
	@$(GOLANG) get github.com/Telmate/proxmox-api-go/@latest

format:
	@echo " 🔘 Go Format"
	@$(GOFMT) -s -w .

vet:
	@echo " 🔘 Go Vet"
	@$(GOLANG) vet ./...

tidy:
	@echo " 🔘 Go Mod Tidy"
	@$(GOLANG) mod tidy

download:
	@echo " 🔘 Go Mod Download"
	@$(GOLANG) mod download

verify:
	@echo " 🔘 Go Mod Verify"
	@$(GOLANG) mod verify

golangci-lint:
	@echo " 🔘 Run golangci-lint"
	@$(GOLANGCI_LINT) run ./...

lint:
	@echo " 🔘 Run golint"
	@$(GOLINT) ./...

staticcheck:
	@echo " 🔘 Run staticcheck"
	@$(STATICCHECK) ./...

cyclo:
	@echo " 🔘 Run gocyclo"
	@echo ' 🌀 ----- files -----'
	@$(GOCYCLO) . || exit 0
	@echo ' 🌀 ----- top 10 -----'
	@$(GOCYCLO) -top 10 . || exit 0
	@echo ' 🌀 ----- over 5 -----'
	@$(GOCYCLO) -over 5 . || exit 0
	@echo ' 🌀 ----- average -----'
	@$(GOCYCLO) -avg . || exit 0
	@echo ' 🌀 ----- over 2 average -----'
	@$(GOCYCLO) -over 2 -avg . || exit 0

ineffassign:
	@echo " 🔘 Run ineffassign"
	@$(INEFFASSIGN) ./...

gosec:
	@echo " 🔘 Run gosec"
	@$(GOSEC) ./...

vulncheck:
	@echo " 🔘 Run govulncheck"
	@$(GOVULNCHECK) -test -show verbose ./...

clean:
	@echo " 🗑️  cleaning ..."
	@$(GOLANG) clean
	@rm checksum.txt

update:
	@echo " 📦 updating ..."
	@$(GOLANG) get -u ./...
