package kennitalaerror

import "errors"

var (
	ErrInvalidKennitalaType        = errors.New("invalid argument")
	ErrInvalidKennitalaLength      = errors.New("invalid length")
	ErrInvalidKennitalaCentury     = errors.New("invalid century")
	ErrInvalidKennitalaFirstLetter = errors.New("invalid first letter")
	ErrInvalidKennitalaCheckDigit  = errors.New("invalid check digit")
)
