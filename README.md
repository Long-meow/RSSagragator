# Gator (RSS Aggregator CLI)

Gator is a multi-user CLI application written in Go that fetches, aggregates, and manages RSS feeds in real time using PostgreSQL for persistent storage.

---

## Prerequisites

Before installing and running Gator, ensure you have the following tools installed on your machine:

- **[Go](https://go.dev/doc/install)** (version 1.22 or higher)
- **[PostgreSQL](https://www.postgresql.org/download/)** (running locally or accessible remotely)

Ensure your PostgreSQL service is running and create a dedicated database for Gator:

```bash
createdb gator
