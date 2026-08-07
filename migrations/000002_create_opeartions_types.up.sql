CREATE TABLE IF NOT EXISTS operations_types (
  operation_type_id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  description text NOT NULL
);


