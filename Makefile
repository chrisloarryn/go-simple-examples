# run node: cd node-ts && yarn start:dev
# run go: cd go && go run main.go

# Path: node-ts/src/index.ts
# Path: go/main.go	

clean:
	rm -rf node-ts/node_modules
	rm -rf node-ts/dist
	rm -rf go/go.sum

build-node:
	cd node-ts && yarn build

build-go:
	cd go && go build -o main

# install nodejs packages using yarn
install:
	cd node-ts && yarn install

# run nodejs server
run-node:
	cd node-ts && yarn start:dev

# install go packages
install-go:
	cd go && go mod download

# run go server
run-go:
	cd go && go run main.go

help:
	@echo "install: install nodejs packages using yarn"
	@echo "run-node: run nodejs server"
	@echo "install-go: install go packages"
	@echo "run-go: run go server"
	@echo "help: show this help message"