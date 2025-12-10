# Event Streaming Pipeline (Golang + Kafka)

A simple real-time event streaming pipeline built using **Kafka**, **Golang**, and **Docker**.

## Features
- Go Producer sends JSON events.
- Go Consumer reads messages in real time.
- Kafka + Zookeeper run via Docker Compose.

## Run Steps
1. Start Kafka: `docker-compose up -d`
2. Run consumer: `go run consumer/main.go`
3. Run producer: `go run producer/main.go`
