package money

import (
	"errors"
	"testing"
)

func TestParseDecimal(t *testing.T) {
	tcs := map[string]struct {
		input string
		want  Decimal
		err   error
	}{
		"valid input": {
			input: "123.45",
			want:  Decimal{subunits: 12345, precision: 2},
			err:   nil,
		},
		"invalid input": {
			input: "123.45.67",
			want:  Decimal{},
			err:   ErrParseInvalidDecimal,
		},
		"negative input": {
			input: "-123.45",
			want:  Decimal{subunits: -12345, precision: 2},
			err:   nil,
		},
		"rounded input": {
			input: "1230000.456",
			want:  Decimal{subunits: 1230000456, precision: 3},
			err:   nil,
		},
		"empty input": {
			input: "",
			want:  Decimal{},
			err:   ErrParseInvalidDecimal,
		},
		"zero input": {
			input: "0",
			want:  Decimal{subunits: 0, precision: 0},
			err:   nil,
		},
		"zero decimal input": {
			input: "0.0",
			want:  Decimal{subunits: 0, precision: 0},
			err:   nil,
		},
		"crazy zero input": {
			input: "0.000000000",
			want:  Decimal{subunits: 0, precision: 0},
			err:   nil,
		},
		"long crazy zero input": {
			input: "00000000000000000.00000000000000000000000",
			want:  Decimal{subunits: 0, precision: 0},
			err:   nil,
		},
		"leading zero input": {
			input: "0.123",
			want:  Decimal{subunits: 123, precision: 3},
			err:   nil,
		},
		"trailing zero input": {
			input: "123.450",
			want:  Decimal{subunits: 12345, precision: 2},
			err:   nil,
		},
		"leading and trailing zero input": {
			input: "0.1230",
			want:  Decimal{subunits: 123, precision: 3},
			err:   nil,
		},
		"unnecessary zero input": {
			input: "0000001.123",
			want:  Decimal{subunits: 1123, precision: 3},
			err:   nil,
		},
		"long precision input": {
			input: "123.45678901234567890123456789012345667788990",
			want:  Decimal{},
			err:   ErrParseTooLargeDecimal,
		},
		"not a number input": {
			input: "abc",
			want:  Decimal{},
			err:   ErrParseInvalidDecimal,
		},
		"too long input": {
			input: "12345678901234567890123456789012345667788990",
			want:  Decimal{},
			err:   ErrParseTooLargeDecimal,
		},
		"blank input": {
			input: "   ",
			want:  Decimal{},
			err:   ErrParseInvalidDecimal,
		},
		"leading whitespace input": {
			input: "  123.45",
			want:  Decimal{},
			err:   ErrParseInvalidDecimal,
		},
		"trailing whitespace input": {
			input: "123.45  ",
			want:  Decimal{},
			err:   ErrParseInvalidDecimal,
		},
		"leading and trailing whitespace input": {
			input: "  123.45  ",
			want:  Decimal{},
			err:   ErrParseInvalidDecimal,
		},
		"only separator input": {
			input: ".",
			want:  Decimal{},
			err:   ErrParseInvalidDecimal,
		},
		"leading separator input": {
			input: ".123",
			want:  Decimal{subunits: 123, precision: 3},
			err:   nil,
		},
		"trailing separator input": {
			input: "123.",
			want:  Decimal{subunits: 123, precision: 0},
			err:   nil,
		},
		"leading and trailing separator input": {
			input: ".123.",
			want:  Decimal{},
			err:   ErrParseInvalidDecimal,
		},
		"multiple separators input": {
			input: "123..45",
			want:  Decimal{},
			err:   ErrParseInvalidDecimal,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			got, err := ParseDecimal(tc.input)
			if !errors.Is(err, tc.err) {
				t.Errorf("ParseDecimal(%s) error = %v, want %v", tc.input, err, tc.err)
			}
			if got != tc.want {
				t.Errorf("ParseDecimal(%s) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
