all: linux windows
prepare:
	sudo apt-get update
	sudo apt-get install -y golang
linux:
	go build main.go
windows:
	GOOS=windows GOARCH=amd64 go build main.go
