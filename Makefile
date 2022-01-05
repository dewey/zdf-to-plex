# This Makefile is just for my own convinience of deploying a new version, this can be ignored or adapted for your use
# case.

.PHONY: release
release:
	@echo "> building binary"
	GOOS=linux GOARCH=amd64 go build -o bin/zdf-to-plex-amd64-linux main.go
	@echo "> deploying to server"
	scp bin/zdf-to-plex-amd64-linux ubuntu@notmyhostna.me:~/services/zdf-to-plex/
