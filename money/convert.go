package money

const (
	// ErrConvertToInvalidCurrency is the error returned when the currency to convert to is invalid.
	ErrConvertToInvalidCurrency = Error("invalid currency to convert")

	// ErrConvertFromInvalidAmount is the error returned when the amount to be converted is invalid.
	ErrConvertFromInvalidAmount = Error("invalid amount to be converted")
)

// Convert applies a conversion rate to an Amount and returns a new Amount in the specified Currency.
func Convert(amount Amount, to Currency) (Amount, error) {
	if !amount.IsValid() {
		return Amount{}, ErrConvertFromInvalidAmount
	}

	if !to.IsValid() {
		return Amount{}, ErrConvertToInvalidCurrency
	}

	return Amount{}, nil
}
