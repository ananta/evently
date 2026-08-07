ALTER TABLE operations_types
	ADD CONSTRAINT operations_types_description_key UNIQUE (description);

INSERT INTO operations_types (description)
VALUES
	('Normal Purchase'),
	('Purchase with installments'),
	('Withdrawal'),
	('Credit Voucher')
ON CONFLICT (description) DO NOTHING;
