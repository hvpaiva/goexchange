package money

import (
	"errors"
	"testing"
)

func TestParseCurrency(t *testing.T) {
	tests := map[string]struct {
		code string
		want Currency
		err  error
	}{
		"valid code default": {
			code: "USD",
			want: Currency{code: "USD", precision: 2},
			err:  nil,
		},
		"valid code precision 0": {
			code: "IRR",
			want: Currency{code: "IRR", precision: 0},
			err:  nil,
		},
		"valid code precision 1": {
			code: "CNY",
			want: Currency{code: "CNY", precision: 1},
			err:  nil,
		},
		"valid code precision 3": {
			code: "BHD",
			want: Currency{code: "BHD", precision: 3},
			err:  nil,
		},
		"invalid code by len": {
			code: "US",
			want: Currency{},
			err:  ErrParseInvalidCurrencyCode,
		},
		"invalid code by len 2": {
			code: "USDA",
			want: Currency{},
			err:  ErrParseInvalidCurrencyCode,
		},
		"invalid code by char": {
			code: "US1",
			want: Currency{},
			err:  ErrParseInvalidCurrencyCode,
		},
		"invalid blank code": {
			code: "   ",
			want: Currency{},
			err:  ErrParseInvalidCurrencyCode,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ParseCurrency(test.code)
			if !errors.Is(err, test.err) {
				t.Errorf("want error %v, got %v", test.err, err)
			}
			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestCurrency_IsValid(t *testing.T) {
	tests := map[string]struct {
		currency Currency
		want     bool
	}{
		"valid currency": {
			currency: Currency{code: "USD", precision: 2},
			want:     true,
		},
		"invalid currency": {
			currency: Currency{},
			want:     false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.currency.IsValid()
			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}
