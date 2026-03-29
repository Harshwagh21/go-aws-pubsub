# Pub-Sub Demo using aws sns/sqs

A CLI demo of AWS SNS → SQS pub/sub flow using Docker LocalStack, Go, Lipgloss Terminal UI no real AWS account needed.

## What This Does

Spins up a full pub/sub pipeline locally:

```
You type a message → SNS Topic → SQS Queue → Received & displayed
```

Everything is provisioned by the Go code itself — no manual setup, no AWS console.

## Demo Image

![image](https://github.com/user-attachments/assets/104f3381-ebaa-4e65-9383-ad163c131d2c)

## Tech Stack

| Tool        | Role                  |
| ----------- | --------------------- |
| Go          | Core application      |
| AWS SDK v2  | SNS + SQS clients     |
| LocalStack  | AWS emulation locally |
| Docker      | Runs LocalStack       |
| lipgloss    | Terminal UI styling   |
| progressbar | CLI progress bars     |

## Project Structure

```
go-aws-pubsub/
├── cmd/main.go              # Orchestrates the full flow
├── internal/
│   ├── aws/client.go        # LocalStack AWS config
│   ├── sns/publisher.go     # Create topic + publish
│   ├── sns/subscribe.go     # Subscribe SQS to SNS
│   ├── sqs/consumer.go      # Create queue + receive
│   └── ui/ui.go             # All terminal UI
├── docker-compose.yml
└── Makefile
```

## Prerequisites

- [Go 1.21+](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- Make

## Run

```bash
# 1. Start LocalStack (make sure docker is running)
make up

# 2. Run the app
make run

# 3. Stop LocalStack when done
make down
```

## Flow

1. Loads LocalStack AWS config pointing to `localhost:4566`
2. Creates SNS Topic — `Publisher-Sns-Topic`
3. Creates SQS Queue — `Consumer-Sqs-Queue`
4. Subscribes the queue to the topic
5. You type a message
6. Message published to SNS → routed to SQS
7. Message received from SQS and displayed

## Available Commands

```bash
make up      # Start LocalStack
make down    # Stop LocalStack
make run     # Run the app
make logs    # View LocalStack logs
make clean   # Stop + clean Go build cache
```
