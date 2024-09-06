package money

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	maxDecimal = 12

	// ErrParseInvalidDecimal is the error returned when a decimal is invalid.
	ErrParseInvalidDecimal = Error("invalid decimal")

	// ErrParseTooLargeDecimal is the error returned when a decimal is too large.
	ErrParseTooLargeDecimal = Error("decimal over 10ˆ12 is too large")
)

// Decimal represents a floating-point number with a fixed number of decimal places.
type Decimal struct {
	subunits  int64 // The amount of subunits in the decimal. Multiply by 10^precision to get the amount of units.
	precision byte  // Number of subunits in a unit. Expressed as a power of 10.
}

func ParseDecimal(s string) (Decimal, error) {
	if s == "" || s == "." {
		return Decimal{}, ErrParseInvalidDecimal
	}

	rawIntPart, rawFracPart, _ := strings.Cut(s, ".")

	intPart := strings.TrimLeft(rawIntPart, "0")
	fracPart := strings.TrimRight(rawFracPart, "0")

	if len(intPart)+len(fracPart) == 0 {
		return Decimal{}, nil
	}

	if len(intPart)+len(fracPart) > maxDecimal {
		return Decimal{}, ErrParseTooLargeDecimal
	}

	subunits, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrParseInvalidDecimal, err.Error())
	}

	precision := byte(len(fracPart))

	return Decimal{subunits: subunits, precision: precision}, nil
}
