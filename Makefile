.PHONY: server-test server client all

all: server client

client:
	cd client && npm run dev

server:
	cd server && go run -race cmd/api/main.go

server-game-test:
	cd server/internal/game && go test

setup:
	cd client && npm install
	go mod tidy

clean:
	cd client && npm run clean
	go clean
