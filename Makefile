.PHONY: assets

build:
	go build

assets:
	go-bindata -o assets.go assets/...
