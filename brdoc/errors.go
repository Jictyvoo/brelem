package brdoc

import (
	"errors"

	"github.com/jictyvoo/brelem/internal/docerrors"
)

// IsCPFError reports whether err is a validation error originating from CPF validation.
func IsCPFError(err error) bool {
	var ve docerrors.ValidationError
	if errors.As(err, &ve) {
		return ve.Context == "cpf"
	}
	return false
}

// IsCNPJError reports whether err is a validation error originating from CNPJ validation.
func IsCNPJError(err error) bool {
	var ve docerrors.ValidationError
	if errors.As(err, &ve) {
		return ve.Context == "cnpj"
	}
	return false
}

// IsCNHError reports whether err is a validation error originating from CNH validation.
func IsCNHError(err error) bool {
	var ve docerrors.ValidationError
	if errors.As(err, &ve) {
		return ve.Context == "cnh"
	}
	return false
}
