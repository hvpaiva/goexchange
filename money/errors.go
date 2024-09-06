package money

// Error defines the type of error in the money package.
type Error string

func (e Error) Error() string {
	return string(e)
}
