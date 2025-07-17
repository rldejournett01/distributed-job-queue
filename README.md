# Distributed Job Queue System
Small scale, simulated distribution system


## Features
- Task distibution to simulated worker nodes
- Priority and retry queues
- Sytem health monitoring (CPU, memory, job logs)
- Optional CLI or web dashboard
- Designed for scale and observability

## Technologies
-  [Language: Go / Python / C++]
-  REST/gRPC API (??)
-  System metrics (psutil, proc / Prometheus)

## Setup

'''bash
# Clone
git clone https://github.com/rldejourett01/distributed-job-queue.git
cd distributed-job-queue

# Install dependencies
[python] pip instll -r requirements.txt
[go] go mod tidy
