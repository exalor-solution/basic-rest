# Subscription REST API (Go)

A simple REST service written in Go to manage subscriptions.

- Base URL: `http://localhost:8080`
- Port: `8080`

## Model

```json
{
  "name": "string",
  "price": 0.0,
  "currency": "USD"
}
```

Fields:
- `name` (string) required, used as the unique key
- `price` (number) required
- `currency` (string) required (examples: `USD`, `CAD`, `EUR`)

---

## Endpoints

### 1) Add Subscription

Create a new subscription.

- Method: `POST`
- Path: `/add`
- Content-Type: `application/json`

Request JSON:
```json
{
  "name": "Netflix",
  "price": 16.99,
  "currency": "CAD"
}
```

cURL:
```bash
curl -X POST http://localhost:8080/add   -H "Content-Type: application/json"   -d '{"name":"Netflix","price":16.99,"currency":"CAD"}'
```

---

### 2) Update Subscription

Update an existing subscription by `name`.

- Method: `PUT`
- Path: `/put?name=Netflix`
- Content-Type: `application/json`

Example:
- `/put?name=Netflix`

Request JSON:
```json
{
  "name": "Netflix",
  "price": 18.99,
  "currency": "CAD"
}
```

cURL:
```bash
curl -X PUT http://localhost:8080/put?name=Netflix   -H "Content-Type: application/json"   -d '{"name":"Netflix","price":18.99,"currency":"CAD"}'
```

---

### 3) Find Subscription (by name)

Get one subscription by `name`.

- Method: `GET`
- Path: `/find?name=Netflix`

Example:
- `GET /find?name=Netflix`

cURL:
```bash
curl http://localhost:8080//find?name=Netflix
```

Response JSON:
```json
{
  "name": "Netflix",
  "price": 18.99,
  "currency": "CAD"
}
```

---

### 4) Delete Subscription (by name)

Delete one subscription by `name`.

- Method: `DELETE`
- Path: `/del?name=Netflix`

Example:
- `DELETE //del?name=Netflix`

cURL:
```bash
curl -X DELETE http://localhost:8080/del?name=Netflix
```

---

## Example JSON Payload (copy and paste)

```json
{
  "name": "Spotify",
  "price": 11.99,
  "currency": "USD"
}
```
