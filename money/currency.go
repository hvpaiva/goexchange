package money

import "unicode"

// ErrParseInvalidCurrencyCode is the error returned when a currency code is invalid.
const ErrParseInvalidCurrencyCode = Error("invalid currency code")

// Currency represents the code and precision of a currency.
//
// The code is a 3-letter string. And a zero value of Currency is invalid.
type Currency struct {
	code      string
	precision byte
}

// ParseCurrency parses a currency code and returns a Currency.
//
// The code must be a 3-letter string.
// If a code is not valid, it returns an ErrParseInvalidCurrencyCode error.
func ParseCurrency(code string) (Currency, error) {
	if len(code) != 3 {
		return Currency{}, ErrParseInvalidCurrencyCode
	}

	for _, char := range code {
		if !unicode.IsLetter(char) {
			return Currency{}, ErrParseInvalidCurrencyCode
		}
	}

	switch code {
	case "IRR":
		return Currency{code: code, precision: 0}, nil
	case "CNY", "VND":
		return Currency{code: code, precision: 1}, nil
	case "BHD", "IQD", "JOD", "KWD", "LYD", "OMR", "TND":
		return Currency{code: code, precision: 3}, nil
	default:
		return Currency{code: code, precision: 2}, nil
	}
}

// IsValid returns true if the currency is valid.
//
// A currency is valid if its code is not empty.
func (c Currency) IsValid() bool {
	return c.code != ""
}
