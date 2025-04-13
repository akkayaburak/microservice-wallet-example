# Microservice Wallet Example

## Services

### Wallet Service
Manages crypto wallets.
- Create / Delete / List wallets
- Each wallet is uniquely identified by `address + network`
- DB: SQLite
- Some unit tests included

### Asset Service
Manages assets.
- Create / Update / List assets.
- Some amount checks
- DB: SQLite
- Some unit tests included

### Scheduled Transactions
Executes transactions at a specified time.
- Uses `robfig/cron`
- Marks executed transactions as `executed` in DB
- Handles time-based queries safely

---

## Tech Stack

- **Language:** Go
- **Database:** SQLite (can be swapped with Postgres)
- **Architecture:** Microservices
- **Scheduling:** `github.com/robfig/cron`
- **Testing:** Standard Go testing with memory DB

## Todo

- Transfer data between services
- Integration tests (and better unit tests too)
- User implementation
- Swagger

## Running Tests
```bash
go test ./...
```

## Getting Started
```bash
git clone https://github.com/akkayaburak/microservice-wallet-example.git
cd microservice-wallet-example

# Build and run all services with Docker Compose
docker-compose up --build

```

# API Endpoints

### Wallet Service

Base URL: `http://localhost:8081`

---

#### Create Wallet

- **POST** `/wallet`
- **Request Body:**
```json
{
  "address": "1Lbcfr7sAHTD9CgdQo3HTMTkV8LK4ZnX71",
  "network": "Bitcoin"
}
```

#### Get Wallet

- **GET** `/wallet{address}/{network}`

#### Delete Wallet

- **DELETE** `/wallet`
- **Request Body:**
```json
{
  "address": "1Lbcfr7sAHTD9CgdQo3HTMTkV8LK4ZnX71",
  "network": "Bitcoin"
}
```


### Asset Service

Base URL: `http://localhost:8082`

---
#### Create Asset

- **POST** `/assets`
- **Request Body:**
```json
{
  "wallet_id": 3,
  "symbol": "BTC",
  "amount: 10,
  "network": "Bitcoin"
}
```

#### List Assets

- **GET** `/assets/list`

#### Create Asset

- **POST** `/assets`
- **Request Body:**
```json
{
  "id": 2,
  "amount: 10,
}
```

### Scheduled Service

Base URL: `http://localhost:8083`

---

#### Create Scheduled Transaction

- **POST** `/schedule`
- **Request Body:**
```json
{
  "id": 1,
  "wallet_id": 42,
  "to_address": "3FZbgi29cpjq2GjdwV8eyHuJJnkLtktZc5",
  "amount": 0.5,
  "symbol": "BTC",
  "network": "Bitcoin",
  "scheduled_at": "2025-04-13T15:00:00Z",
  "status": "pending"
}
```

#### List Scheduled Transactions

- **GET** `/schedules`




