.PHONY: server all

all: server

client:
	cd client && npm run dev

server:
	cd server && go run -race cmd/api/main.go

server-game-test:
	cd server/internal/game && go test

setup:
	cd server && go mod tidy
	cd client && npm install

clean:
	cd client && npm run clean
	cd server && go clean
