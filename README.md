# Rabbit Hole

Rabbithole is a simple program for load testing RabbitMQ. It allows you to send a custom number of messages of a configurable size to a queue.

## Usage

```
Usage of rabbithole:
      --count int      Number of messages to send to the queue (default 100)
      --host string    Full URL of the RabbitMQ server (default "amqp://guest:guest@localhost:5672/")
      --max-size int   Maximum size of messages in bytes (default 10000)
      --min-size int   Minimum size of messages in bytes (default 100)
      --queue string   Name of the queue to send messages to (default "rabbithole")
      --verbose        Enable verbose logging
      --version        Show version information
```