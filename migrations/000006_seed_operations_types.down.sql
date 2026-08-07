DELETE FROM operations_types
WHERE description IN (
	'Normal Purchase',
	'Purchase with installments',
	'Withdrawal',
	'Credit Voucher'
);

ALTER TABLE operations_types
	DROP CONSTRAINT operations_types_description_key;
