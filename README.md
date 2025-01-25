# NavalWarfare

A P2P game based on the classic "Battleship" game, implemented to allow players enjoy the game across the internet

## Game Demo

![]()
![]()

### Technologies used

Golang, Typescript, React

## Usage

### Prerequisites

Make sure you have Golang installed and a recent version of Node installed (18+)
Clone the repository
Setup you environment variable to enable a connection to the database

### Steps

- Open a terminal and run the following commands

- Install packages for both server and client with the following command

```bash
make setup
```

- Use the following command to ensure all tests pass

```bash
make server-game-test
```

- Open a terminal and use the following command to start the server

```bash
make server
```

- Open a terminal and use the following command to start the client

```bash
make client
```
