.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./count-visitors/count-visitors
	
build:
	GOOS=linux GOARCH=amd64 go build -o count-visitors/count-visitors ./count-visitors
