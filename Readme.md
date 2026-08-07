## Evently

Test app for interview
## Running it

Requires Docker. Nothing else — no Go toolchain needed.

```bash
./run
```

That builds the API, starts PostgreSQL, applies migrations and serves on
**http://localhost:4000**. Add `-d` to run in the background.

```bash
curl localhost:4000/healthcheck
```

To stop, and to discard the database volume:

```bash
docker compose down
docker compose down -v
```

## API

All requests and responses are JSON. Amounts are JSON numbers with up to two
decimal places.

### `POST /accounts`

```bash
curl -X POST localhost:4000/accounts \
  -d '{"document_number": "12345678900"}'
```

```json
{ "account_id": 1, "document_number": "12345678900" }
```

| Status | When |
| --- | --- |
| `201` | Created |
| `400` | Body is malformed, empty, oversized or has unknown fields |
| `422` | Document number is missing, not all digits, over 50 characters, or already registered |

### `GET /accounts/:accountId`

```bash
curl localhost:4000/accounts/1
```

```json
{ "account_id": 1, "document_number": "12345678900" }
```

| Status | When |
| --- | --- |
| `200` | Found |
| `404` | No such account, or the id is not a positive integer |

### `POST /transactions`

```bash
curl -X POST localhost:4000/transactions \
  -d '{"account_id": 1, "operation_type_id": 4, "amount": 123.45}'
```

```json
{
  "transaction_id": 1,
  "account_id": 1,
  "operation_type_id": 4,
  "amount": 123.45,
  "event_date": "2026-08-07T16:48:28.223821Z"
}
```

| Status | When |
| --- | --- |
| `201` | Created |
| `400` | Body is malformed, empty, oversized or has unknown fields |
| `422` | A field failed validation (see below) |

### Operation types

| ID | Description | Amount sign |
| --- | --- | --- |
| 1 | Normal Purchase | negative |
| 2 | Purchase with installments | negative |
| 3 | Withdrawal | negative |
| 4 | Credit Voucher | positive |

### Validation errors

A `422` names every field that failed, so one round trip reports all problems:

```bash
curl -X POST localhost:4000/transactions \
  -d '{"account_id": 0, "operation_type_id": 99, "amount": 0}'
```

```json
{
  "error": {
    "account_id": "must be a positive integer",
    "amount": "must not be zero",
    "operation_type_id": "must be a known operation type"
  }
}
```

## Testing

```bash
make test    # run the suite with the race detector
make check   # gofmt, go vet and the tests — run this before pushing
```

The suite requires no database; everything currently under test is pure.
## Developing without Docker

Requires Go 1.26. Start the database and apply migrations, then run the API on
the host:

```bash
make up
make migrate-up
make run
```

Other migration commands:

```bash
make create-migration name=add_something
make migrate-down
make migrate-version
make migrate-force VERSION=3
```

Connect to the database directly:

```bash
docker compose exec db psql -U evently_user -d evently_db
```

## Layout

```
cmd/api/          HTTP handlers, routing, middleware, validation
internal/data/    models and database access
migrations/       golang-migrate SQL migrations
```

## Dependencies

- [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter) — HTTP routing
- [lib/pq](https://github.com/lib/pq) — PostgreSQL driver
- [shopspring/decimal](https://github.com/shopspring/decimal) — exact decimal arithmetic for monetary amounts
