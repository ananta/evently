CREATE TABLE IF NOT EXISTS transactions (
  transaction_id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  account_id bigint NOT NULL REFERENCES accounts,
  operation_type_id bigint NOT NULL REFERENCES operations_types,
  amount DECIMAL(15,2) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
