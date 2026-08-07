package data

const (
	NormalPurchase           int64 = 1
	PurchaseWithInstallments int64 = 2
	Withdrawal               int64 = 3
	CreditVoucher            int64 = 4
)

// IsOperationType reports whether id is a known operation type.
func IsOperationType(id int64) bool {
	return id >= NormalPurchase && id <= CreditVoucher
}

// IsCredit reports whether transactions of this type are recorded as positive
// amounts. Everything else is recorded negative.
func IsCredit(id int64) bool {
	return id == CreditVoucher
}
