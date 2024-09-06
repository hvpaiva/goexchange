package money

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewAmount(t *testing.T) {
	tests := map[string]struct {
		quantity Decimal
		currency Currency
		want     Amount
		err      error
	}{
		"valid amount": {
			quantity: Decimal{subunits: 12345, precision: 2},
			currency: Currency{code: "USD", precision: 2},
			want: Amount{
				quantity: Decimal{subunits: 12345, precision: 2},
				currency: Currency{code: "USD", precision: 2},
			},
		},
		"invalid amount": {
			quantity: Decimal{subunits: 12345, precision: 2},
			currency: Currency{code: "USD", precision: 3},
			want:     Amount{},
			err:      ErrWrongPrecisionAmount,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewAmount(test.quantity, test.currency)

			if !errors.Is(err, test.err) {
				t.Errorf("want error %v, got %v", test.err, err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestAmount_IsValid(t *testing.T) {
	tests := map[string]struct {
		amount Amount
		want   bool
	}{
		"valid amount": {
			amount: Amount{
				quantity: Decimal{subunits: 12345, precision: 2},
				currency: Currency{code: "USD", precision: 2},
			},
			want: true,
		},
		"invalid amount": {
			amount: Amount{},
			want:   false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.amount.IsValid()

			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}
