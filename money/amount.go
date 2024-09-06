package money

const (
	// ErrWrongPrecisionAmount is the error returned when the precision of the amount is wrong.
	ErrWrongPrecisionAmount = Error("wrong precision amount")
)

// Amount represents a quantity of money in a specific Currency.
// A zero value Amount is not valid.
type Amount struct {
	quantity Decimal
	currency Currency
}

// NewAmount creates a new Amount with the given quantity and currency.
//
// It returns an ErrWrongPrecisionAmount error if the precision of the quantity is bigger than the precision of the
// currency.
func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		return Amount{}, ErrWrongPrecisionAmount
	}

	quantity.precision = currency.precision

	return Amount{quantity: quantity, currency: currency}, nil
}

// IsValid returns true if the amount is valid, false otherwise.
//
// An amount is valid if its currency is valid.
func (a Amount) IsValid() bool {
	return a.currency.IsValid()
}
