.PHONY: build clean

build:
		env GOOS=linux go build -ldflags="-s -w" -o bin/qbprojects cmd/main.go

clean:
		rm -rf ./bin ./vendor Gopkg.lock