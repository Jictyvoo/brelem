package docerrors

import "errors"

// ErrorReason describes why a document validation failed.
type ErrorReason uint8

const (
	ReasonNoError ErrorReason = iota
	ReasonEmpty
	ReasonIncorrectLength
	ReasonAllSameDigit
	ReasonInvalidDigits
	ReasonExpired
)

// reasonMessages maps each ErrorReason to a pre-allocated message suffix.
var reasonMessages = [...]string{
	0:                     "unknown error",
	ReasonEmpty:           "empty input",
	ReasonIncorrectLength: "incorrect length",
	ReasonAllSameDigit:    "all digits are the same",
	ReasonInvalidDigits:   "invalid check digits",
	ReasonExpired:         "document is expired",
}

// ValidationError is a generic document validation error.
// It carries a Reason and a Context string set by the caller (e.g. "cpf", "cnpj").
type ValidationError struct {
	Reason  ErrorReason
	Context string
}

// Error returns a human-readable message like "cpf: invalid check digits".
func (e ValidationError) Error() string {
	msg := "unknown error"
	if int(e.Reason) < len(reasonMessages) {
		msg = reasonMessages[e.Reason]
	}
	if e.Context != "" {
		return e.Context + ": " + msg
	}
	return msg
}

// New creates a ValidationError with the given reason and context.
func New(reason ErrorReason, context string) ValidationError {
	return ValidationError{Reason: reason, Context: context}
}

func checkErrorReason(err error, reason ErrorReason) bool {
	var ve ValidationError
	if errors.As(err, &ve) {
		return ve.Reason == reason
	}
	return false
}

// IsEmpty reports whether err is a ValidationError with ReasonEmpty.
func IsEmpty(err error) bool {
	return checkErrorReason(err, ReasonEmpty)
}

// IsIncorrectLength reports whether err is a ValidationError with ReasonIncorrectLength.
func IsIncorrectLength(err error) bool {
	return checkErrorReason(err, ReasonIncorrectLength)
}

// IsAllSameDigit reports whether err is a ValidationError with ReasonAllSameDigit.
func IsAllSameDigit(err error) bool {
	return checkErrorReason(err, ReasonAllSameDigit)
}

// IsInvalidDigits reports whether err is a ValidationError with ReasonInvalidDigits.
func IsInvalidDigits(err error) bool {
	return checkErrorReason(err, ReasonInvalidDigits)
}

// IsExpired reports whether err is a ValidationError with ReasonExpired.
func IsExpired(err error) bool {
	return checkErrorReason(err, ReasonExpired)
}
