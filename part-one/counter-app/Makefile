.PHONY: all build-app build-container push

all: build-app build-container push

build-app:
	GOOS=linux GOARCH=amd64 go build -o counter-app .

build-container:
	docker build -t bmcrae/counter-app:latest .

push:
	docker push bmcrae/counter-app:latest
