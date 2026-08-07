package main

import (
	"testing"

	"github.com/ananta/evently/internal/data"
	"github.com/shopspring/decimal"
)

func TestValidDocumentNumber(t *testing.T) {
	tests := []struct {
		name string
		doc  string
		want bool
	}{
		{"valid", "12345678900", true},
		{"empty", "", false},
		{"letters", "12E45678900", false},
		{"punctuation", "123.456.789-00", false},
		{"too long", "123456789012345678901234567890123456789012345678901", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validDocumentNumber(tt.doc) == ""; got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidAmount(t *testing.T) {
	tests := []struct {
		name   string
		opType int64
		amount string
		want   bool
	}{
		{"voucher positive", data.CreditVoucher, "123.45", true},
		{"voucher negative", data.CreditVoucher, "-123.45", false},
		{"purchase negative", data.NormalPurchase, "-50.00", true},
		{"purchase positive", data.NormalPurchase, "50.00", false},
		{"withdrawal negative", data.Withdrawal, "-20.00", true},
		{"zero", data.CreditVoucher, "0", false},
		{"unknown type defers to the type error", 99, "10.00", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount := decimal.RequireFromString(tt.amount)
			if got := validAmount(tt.opType, amount) == ""; got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
