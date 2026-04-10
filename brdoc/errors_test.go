package brdoc

import (
	"errors"
	"testing"

	"github.com/jictyvoo/brelem/internal/docerrors"
)

func TestIsCPFError(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		expect bool
	}{
		{"cpf context", docerrors.New(docerrors.ReasonInvalidDigits, "cpf"), true},
		{"cnpj context", docerrors.New(docerrors.ReasonInvalidDigits, "cnpj"), false},
		{"cnh context", docerrors.New(docerrors.ReasonInvalidDigits, "cnh"), false},
		{"nil", nil, false},
		{"other error", errors.New("other"), false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsCPFError(tc.err); got != tc.expect {
				t.Errorf("IsCPFError() = %v, want %v", got, tc.expect)
			}
		})
	}
}

func TestIsCNPJError(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		expect bool
	}{
		{"cnpj context", docerrors.New(docerrors.ReasonInvalidDigits, "cnpj"), true},
		{"cpf context", docerrors.New(docerrors.ReasonInvalidDigits, "cpf"), false},
		{"nil", nil, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsCNPJError(tc.err); got != tc.expect {
				t.Errorf("IsCNPJError() = %v, want %v", got, tc.expect)
			}
		})
	}
}

func TestIsCNHError(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		expect bool
	}{
		{"cnh context", docerrors.New(docerrors.ReasonInvalidDigits, "cnh"), true},
		{"cpf context", docerrors.New(docerrors.ReasonInvalidDigits, "cpf"), false},
		{"nil", nil, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsCNHError(tc.err); got != tc.expect {
				t.Errorf("IsCNHError() = %v, want %v", got, tc.expect)
			}
		})
	}
}
