.PHONY: build clean deploy

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main cmd/*.go

docs: build
	./bin/main -gendocs
	(cd frontend && npm run gen)

clean:
	rm -rf ./bin ./vendor Gopkg.lock coverage.*

format: 
	gofmt -w ./...

test:
	go test ./...

integration:
	go test -v -tags integration -run TestIntegration ./...

cov:
	-go test -coverpkg=./... -coverprofile=coverage.txt -covermode count ./...
	-gocover-cobertura < coverage.txt > coverage.xml
	-go tool cover -html=coverage.txt -o coverage.html
	-go tool cover -func=coverage.txt

lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.44.2 golangci-lint run --enable gofmt,stylecheck,gosec ./...
