BINARY_NAME=MarketWatcher.app
APP_NAME=MarketWatcher
VERSION=1.0.0
BUILD_NO=1

## build: build binary and package app
build:
	rm -rf ${BINARY_NAME}
	fyne package -appVersion ${VERSION} -appBuild ${BUILD_NO} -name ${APP_NAME} -release
	rm -f go-for-MarketWatcher

## run: builds and runs the app
run:
	env DB_PATH="./sql.db" go run .

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@echo "Done!"

## tests: runs all tests
test:
	go test -v ./...