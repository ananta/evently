package data

import (
	"encoding/json"
	"testing"
)

func TestAmountParsesExactly(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"basic example", "123.45", "123.45"},
		{"negative purchase", "-50.0", "-50"},
		{"float32 rounded this up", "999999.99", "999999.99"},
		{"float32 lost 10 cents here", "12345678.90", "12345678.9"},
		{"column upper bound", "9999999999999.99", "9999999999999.99"},
		{"smallest unit", "0.01", "0.01"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var txn Transaction

			err := json.Unmarshal([]byte(`{"amount":`+tt.in+`}`), &txn)
			if err != nil {
				t.Fatalf("unmarshal %s: %v", tt.in, err)
			}

			if got := txn.Amount.String(); got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}

// malformed amount must be an error, never a silent 0.00.
func TestAmountRejectsInvalidJSON(t *testing.T) {
	for _, in := range []string{`"abc"`, `""`, `true`, `[]`} {
		t.Run(in, func(t *testing.T) {
			var txn Transaction

			err := json.Unmarshal([]byte(`{"amount":`+in+`}`), &txn)
			if err == nil {
				t.Errorf("want an error, got amount %s", txn.Amount)
			}
		})
	}
}
