package money_test

import (
	"errors"
	"testing"

	"github.com/hvpaiva/goexchange/money"
)

func TestConvert(t *testing.T) {
	tcs := map[string]struct {
		amount   money.Amount
		currency money.Currency
		validate func(t *testing.T, got money.Amount, err error)
	}{
		"500 USD to BRL": {
			amount:   mustCreateAmount(t, "500", "USD"),
			currency: mustParseCurrency(t, "BRL"),
			validate: func(t *testing.T, got money.Amount, err error) {
				if err != nil {
					t.Errorf("expected no error, got %s", err)
				}

				want := money.Amount{}
				if got != want {
					t.Errorf("expected %v, got %v", want, got)
				}
			},
		},
		"zero value currency": {
			amount:   mustCreateAmount(t, "500", "USD"),
			currency: money.Currency{},
			validate: func(t *testing.T, got money.Amount, err error) {
				if err == nil {
					t.Error("expected error, got none")
				}

				want := money.Amount{}
				if got != want {
					t.Errorf("expected %v, got %v", want, got)
				}

				if !errors.Is(err, money.ErrConvertToInvalidCurrency) {
					t.Errorf("expected error %v, got %v", money.ErrConvertToInvalidCurrency, err)
				}
			},
		},
		"zero value amount": {
			amount:   money.Amount{},
			currency: mustParseCurrency(t, "BRL"),
			validate: func(t *testing.T, got money.Amount, err error) {
				if err == nil {
					t.Error("expected error, got none")
				}

				want := money.Amount{}
				if got != want {
					t.Errorf("expected %v, got %v", want, got)
				}

				if !errors.Is(err, money.ErrConvertFromInvalidAmount) {
					t.Errorf("expected error %v, got %v", money.ErrConvertFromInvalidAmount, err)
				}
			},
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.currency)
			tc.validate(t, got, err)
		})
	}
}

func mustParseCurrency(t *testing.T, code string) money.Currency {
	t.Helper()

	currency, err := money.ParseCurrency(code)
	if err != nil {
		t.Fatalf("cannot parse currency %s code: %s", code, err)
	}

	return currency
}

func mustParseDecimal(t *testing.T, input string) money.Decimal {
	t.Helper()

	decimal, err := money.ParseDecimal(input)
	if err != nil {
		t.Fatalf("cannot parse decimal %s: %s", input, err)
	}

	return decimal
}

func mustCreateAmount(t *testing.T, quantity string, code string) money.Amount {
	t.Helper()

	decimal := mustParseDecimal(t, quantity)

	currency := mustParseCurrency(t, code)

	amount, err := money.NewAmount(decimal, currency)
	if err != nil {
		t.Fatalf("cannot create amount: %s", err)
	}

	return amount
}
