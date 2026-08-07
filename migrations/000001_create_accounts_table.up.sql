CREATE TABLE IF NOT EXISTS accounts (
  account_id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  document_number varchar(50) NOT NULL UNIQUE,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

