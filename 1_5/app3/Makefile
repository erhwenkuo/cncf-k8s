
## usage: make build repos=xxxx tag=v0.1
# default
repos=docker.io/witlab/ingress-training-app3
tag=v0.1

all: push

build:
	GOOS=linux GOARCH=amd64 go build -o app3 main.go

buildimage: build 
	docker build -t $(repos):$(tag) .

push: buildimage
	docker push $(repos):$(tag)

pull:
	docker pull $(repos):$(tag)

test:
	go test -v ./...


