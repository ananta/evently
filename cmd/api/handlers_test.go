package main

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ananta/evently/internal/data"
)

// These tests cover handler wiring: that validation runs, that store errors
// map to the right status, and that routing works. The rules themselves are
// covered by validation_test.go, so only one representative failure per
// handler is exercised here.

// stubAccounts knows about exactly one account, so handler behaviour can be
// tested without a database. Keying off the id means a handler that looks up
// the wrong account is caught.
type stubAccounts struct {
	existingID int64
	insertErr  error
}

func (s stubAccounts) Get(ctx context.Context, id int64) (*data.Account, error) {
	if id != s.existingID {
		return nil, data.ErrRecordNotFound
	}
	return &data.Account{AccountID: id, DocumentNumber: "12345678900"}, nil
}

func (s stubAccounts) Insert(ctx context.Context, account *data.Account) error {
	if s.insertErr != nil {
		return s.insertErr
	}
	account.AccountID = 1
	return nil
}

type stubTransactions struct {
	insertErr error
}

func (s stubTransactions) Insert(ctx context.Context, transaction *data.Transaction) error {
	if s.insertErr != nil {
		return s.insertErr
	}
	transaction.TransactionID = 1
	return nil
}

func newTestApp(accounts accountStore, transactions transactionStore) *application {
	return &application{
		logger:       slog.New(slog.NewJSONHandler(io.Discard, nil)),
		accounts:     accounts,
		transactions: transactions,
	}
}

// do sends a request through the real router and middleware stack.
func do(app *application, method, path, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	rr := httptest.NewRecorder()
	app.routes().ServeHTTP(rr, req)
	return rr
}

func TestCreateAccount(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		insertErr error
		wantCode  int
	}{
		{"valid", `{"document_number":"12345678900"}`, nil, http.StatusCreated},
		{"invalid document", `{"document_number":""}`, nil, http.StatusUnprocessableEntity},
		{"duplicate", `{"document_number":"12345678900"}`, data.ErrDuplicateDocument, http.StatusUnprocessableEntity},
		{"malformed json", `{"document_number":`, nil, http.StatusBadRequest},
		{"unknown field", `{"doc":"123"}`, nil, http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := newTestApp(stubAccounts{insertErr: tt.insertErr}, stubTransactions{})

			rr := do(app, http.MethodPost, "/accounts", tt.body)

			if rr.Code != tt.wantCode {
				t.Errorf("got %d, want %d (body: %s)", rr.Code, tt.wantCode, rr.Body)
			}
		})
	}
}

func TestGetAccount(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		wantCode int
	}{
		{"found", "/accounts/1", http.StatusOK},
		{"not found", "/accounts/9999", http.StatusNotFound},
		{"non-numeric id", "/accounts/abc", http.StatusNotFound},
		{"zero id", "/accounts/0", http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := newTestApp(stubAccounts{existingID: 1}, stubTransactions{})

			rr := do(app, http.MethodGet, tt.path, "")

			if rr.Code != tt.wantCode {
				t.Errorf("got %d, want %d (body: %s)", rr.Code, tt.wantCode, rr.Body)
			}
		})
	}
}

func TestCreateTransaction(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantCode int
	}{
		{"credit voucher", `{"account_id":1,"operation_type_id":4,"amount":123.45}`, http.StatusCreated},
		{"normal purchase", `{"account_id":1,"operation_type_id":1,"amount":-50.00}`, http.StatusCreated},
		{"invalid amount", `{"account_id":1,"operation_type_id":4,"amount":-5.00}`, http.StatusUnprocessableEntity},
		{"unknown account", `{"account_id":9999,"operation_type_id":4,"amount":5.00}`, http.StatusUnprocessableEntity},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := newTestApp(stubAccounts{existingID: 1}, stubTransactions{})

			rr := do(app, http.MethodPost, "/transactions", tt.body)

			if rr.Code != tt.wantCode {
				t.Errorf("got %d, want %d (body: %s)", rr.Code, tt.wantCode, rr.Body)
			}
		})
	}
}
