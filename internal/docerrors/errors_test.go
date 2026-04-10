package docerrors

import (
	"errors"
	"testing"
)

func TestValidationError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      ValidationError
		expected string
	}{
		{
			name:     "cpf empty",
			err:      New(ReasonEmpty, "cpf"),
			expected: "cpf: empty input",
		},
		{
			name:     "cnpj incorrect length",
			err:      New(ReasonIncorrectLength, "cnpj"),
			expected: "cnpj: incorrect length",
		},
		{
			name:     "cpf all same digit",
			err:      New(ReasonAllSameDigit, "cpf"),
			expected: "cpf: all digits are the same",
		},
		{
			name:     "cnh invalid digits",
			err:      New(ReasonInvalidDigits, "cnh"),
			expected: "cnh: invalid check digits",
		},
		{
			name:     "no context expired",
			err:      New(ReasonExpired, ""),
			expected: "document is expired",
		},
		{
			name:     "no context unknown reason",
			err:      ValidationError{Reason: ReasonNoError, Context: ""},
			expected: "unknown error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.err.Error()
			if got != tc.expected {
				t.Errorf("got %q, want %q", got, tc.expected)
			}
		})
	}
}

func TestErrorsAs(t *testing.T) {
	var err error = New(ReasonInvalidDigits, "cpf")
	var ve ValidationError
	if !errors.As(err, &ve) {
		t.Fatal("errors.As failed to match ValidationError")
	}
	if ve.Reason != ReasonInvalidDigits {
		t.Errorf("got reason %d, want %d", ve.Reason, ReasonInvalidDigits)
	}
	if ve.Context != "cpf" {
		t.Errorf("got context %q, want %q", ve.Context, "cpf")
	}
}

func TestReasonHelpers(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		check  func(error) bool
		expect bool
	}{
		{"IsEmpty true", New(ReasonEmpty, "cpf"), IsEmpty, true},
		{"IsEmpty false", New(ReasonInvalidDigits, "cpf"), IsEmpty, false},
		{"IsEmpty nil", nil, IsEmpty, false},
		{"IsIncorrectLength true", New(ReasonIncorrectLength, "cnpj"), IsIncorrectLength, true},
		{"IsIncorrectLength false", New(ReasonEmpty, "cnpj"), IsIncorrectLength, false},
		{"IsAllSameDigit true", New(ReasonAllSameDigit, "cpf"), IsAllSameDigit, true},
		{"IsAllSameDigit false", New(ReasonEmpty, "cpf"), IsAllSameDigit, false},
		{"IsInvalidDigits true", New(ReasonInvalidDigits, "cnh"), IsInvalidDigits, true},
		{"IsInvalidDigits false", New(ReasonEmpty, "cnh"), IsInvalidDigits, false},
		{"IsExpired true", New(ReasonExpired, "cnh"), IsExpired, true},
		{"IsExpired false", New(ReasonEmpty, "cnh"), IsExpired, false},
		{"non-ValidationError", errors.New("other"), IsEmpty, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.check(tc.err); got != tc.expect {
				t.Errorf("got %v, want %v", got, tc.expect)
			}
		})
	}
}
