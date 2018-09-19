GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
NAME=bird.bin

all: fmt vet install

build:
	$(GO_BUILD_ENV) go build -v -o $(NAME) .
install:
	go install
vet:
	go vet .
fmt:
	gofmt -l -w .
clean:
	rm -rf $(NAME)

config-heroku:
	cd scripts && ./env-heroku.sh
run-local:
	cd scripts && ./env-local.sh
	go run main.go
dist: build
	mkdir -p dist/templates
	mkdir -p dist/static
	cp $(NAME) dist/
	cp -R -f templates/* dist/templates
	cp -R -f static/* dist/static