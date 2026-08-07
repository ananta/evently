## Evently

Test app for interview

### Dependencies
## Testing

```bash
make test    # run the suite with the race detector
make check   # gofmt, go vet and the tests — run this before pushing
```

The suite requires no database; everything currently under test is pure.

## Dependencies

- [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter) — HTTP routing
- [lib/pq](https://github.com/lib/pq) — PostgreSQL driver
- [shopspring/decimal](https://github.com/shopspring/decimal) — exact decimal arithmetic for monetary amounts
