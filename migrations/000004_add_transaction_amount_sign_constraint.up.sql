ALTER TABLE transactions ADD CONSTRAINT amount_sign CHECK (
    (operation_type_id IN (1,2,3) AND amount < 0) OR
    (operation_type_id = 4 AND amount > 0)
);
