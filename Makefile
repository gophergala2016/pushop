.PHONY: assets

assets:
	go-bindata -o assets.go assets/...