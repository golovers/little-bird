GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
NAME=chimchip.bin
build:
	$(GO_BUILD_ENV) go build -v -o $(NAME) .

clean:
	rm -rf $(NAME)

heroku-config:
	heroku config:set $(cat .env | sed '/^$/d; /#[[:print:]]*$/d')
local-config:
	export $(grep -v '^#' .env | xargs)
