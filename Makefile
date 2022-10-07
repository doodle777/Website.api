build-img:
	docker build -t fc-go-runtime  -f build-image/Dockerfile build-image

build: build-img
	docker run --rm -it -v $$(pwd):/tmp fc-go-runtime bash -c "GOARCH=amd64 go build -o /tmp/code//bootstrap /tmp/code/main.go"
	chmod +x code/bootstrap

deploy: build
	s deploy -y