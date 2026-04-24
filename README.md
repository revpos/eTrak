# Event Tracker

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.25.0%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)


**An event tracking system built in Go.**  
Track, process, and analyze events.

</div>


## 📖 Table of Contents

- [Overview](#overview)
- [Features](#-features-current)
- [Architecture](#️-architecture)
- [Prerequisites](#-prerequisites)
- [Getting Started](#getting-started)
  - [Manual](#1-manual)
  - [Git](#2-git)
    - [Clone Repo](#21-clone-repo)
    - [Setup PostgreSQL](#22-setup-postgresql)
    - [Run Migrations](#23-run-migrations)
    - [Run Server](#24-run-server)
- [API Protocols - HTTP](#-api-protocols---http)
- [Example Request](#-example-request)
- [Contributions](#-contributions)
- [License](#-license)



## 🧠 Overview

**Event Tracker** is a lightweight yet an effective event tracking service written in Go, like a mini analytics pipeline. Its designed to be modular for quick scalability and clear project layout.

> **Why Event Tracker?**  
> Built with a "zero-dependency philosophy" in mind — no heavy frameworks, idiomatic Go throughout, raw SQL for migrations and deployable as a single binary.

This project focuses on 
- building a systems from scratch in a pragmatic approach (no frameworks)
- understanding performance tradeoffs, design some relevant solutions
- simulating & debugging real-world failures  

## 🚀 Features (current)

- Event ingestion API (`/track`)
- PostgreSQL DBMS for event sourcing
- Structured project architecture (handler → service → repository)
- JSON-based flexible event schema


## 🏗️ Architecture

```
Client → HTTP API → Service Layer → Repository → PostgreSQL
```

Future evolution:
```
Client → API → Queue → Workers → DB
↓
Redis Cache
```

## 🔧 Prerequisites
 
| Requirement | Version   | Notes                          |
|-------------|-----------|--------------------------------|
| Go          | 1.25.0+   | [Download](https://go.dev)     |
| PostgreSQL  | 18.3+     | Primary event store            |
| Redis       | 8.6.2+    | Caching & pub/sub              |
| Docker      | 29.4.1+   | For local dev environment      |
| Make        | 4.4.1+    | Build automation               |
 

## ⚙️ Getting Started

### 1. Manual Way:

Here is a reference image for the manual project initialization.

![Init](https://github.com/revpos/event-tracker/blob/main/images/init.png)

For a quick setup, run the below commands:

```sh
$ chmod +x scripts/scaffold-go-project.yml

$ ./scripts/scaffold-go-project your-name module-name
```

### 2. Git Way: 

#### 2.1 Clone repo
```bash
git clone https://github.com/your-username/event-tracker.git
cd event-tracker
```
#### 2.2 Setup PostgreSQL

2.2.1 Create DB:

```SQL
CREATE DATABASE events;
```

2.2.2 Update connection string in environment:

```
export DB_URL="postgres://user:pass@localhost:5432/events?sslmode=disable"
```

#### 2.3 Run migrations

```sh
$ psql -d events -f migrations/001_create_events.sql
```

#### 2.4 Run server

```sh
$ go run cmd/server/main.go
```

Server runs on: http://localhost:8080


## 📡 API Protocols - HTTP

HTTP Request:
```
POST /track
{
  "user_id": "123",
  "event_type": "login",
  "properties": {
    "device": "mobile"
  }
}
```

HTTP Response: `202 Accepted`


### 🧪 Example Request

```sh
$ curl -X POST http://localhost:8080/track \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "123",
    "event_type": "login",
    "properties": {"device": "mobile"}
  }'
```


## 🛣️ Roadmap
- Async ingestion (worker queue)
- Batch processing
- Redis caching
- Observability (metrics + logs)


## 🤝 Contributions
PRs are welcome.

### Start with:
- improving error handling
- adding validation
- writing tests


## 📄 License

MIT License

Copyright (c) 2026 Revanth Madupoju


---
<div align="center">

Made with ❤️ for Go &nbsp;|&nbsp; [⭐ Star on GitHub](https://github.com/revpos/event-tracker)

</div>
