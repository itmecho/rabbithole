# Rabbit Hole

Rabbithole is a simple program for load testing RabbitMQ. It allows you to send a custom number of messages of a configurable size to a queue.

## Installation
Download the latest release binary for your system from the [releases page](https://github.com/itmecho/rabbithole/releases)

## Building it yourself

### Requirements

* Go >=1.11

### Process

Clone the master branch of this repository and enter the directory. Use the provided `Makefile` to build and install the binary for your system
```
git clone https://github.com/itmecho/rabbithole
cd rabbithole
make install
```

This will install the binary to your `$GOPATH/bin` directory.

To build the binary in the current folder, just run:
```
make
```

## Usage
To see a list of configurable flags, run the following command:
```
rabbithole --help
```

## Local Testing
This repository provides a `docker-compose.yml` which will run a rabbitmq server locally for you to test with. It will be available on `localhost:5672` with the default username/password (guest:guest). The defaults for the rabbithole flags will enable you to connect to this server without needing to set them.

```
docker-compose up -d
```

## TODO
- [ ] Tests
- [ ] Add `--delay` flag to add a delay between each message