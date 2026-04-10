package brdoc

import (
	"github.com/jictyvoo/brelem/internal/cnpjcalc"
	"github.com/jictyvoo/brelem/internal/cpfcalc"
	"github.com/jictyvoo/brelem/internal/docerrors"
)

// DocumentType indicates the type of Brazilian document.
type DocumentType uint8

const (
	TypeUnknown DocumentType = iota
	TypeCPF
	TypeCNPJ
)

// DetermineResult holds the outcome of DetermineCPFCNPJ.
// Check Type to know which field is populated.
type DetermineResult struct {
	Type DocumentType
	CPF  CPF
	CNPJ CNPJ
}

// DetermineCPFCNPJ validates the input as both CPF and CNPJ in a single pass
// and returns which one it matched inside a DetermineResult.
func DetermineCPFCNPJ(raw string) (DetermineResult, error) {
	var (
		cpfV   = cpfcalc.New()
		cnpjV  = cnpjcalc.New()
		result DetermineResult
		count  uint
	)

	for _, r := range raw {
		if r >= '0' && r <= '9' {
			d := byte(r - '0')
			if count < LengthCPF {
				result.CPF.digits[count] = d
			}
			if count < LengthCNPJ {
				result.CNPJ.digits[count] = d
			}
			cpfV.Feed(r)
			cnpjV.Feed(r)
			count++
		}
	}

	if count == 0 {
		return DetermineResult{}, docerrors.New(docerrors.ReasonEmpty, "cpf/cnpj")
	}

	_, cpfReason := cpfV.Finish()
	_, cnpjReason := cnpjV.Finish()

	if cpfReason == docerrors.ReasonNoError {
		result.Type = TypeCPF
		return result, nil
	}
	if cnpjReason == docerrors.ReasonNoError {
		result.Type = TypeCNPJ
		return result, nil
	}

	return DetermineResult{}, docerrors.New(docerrors.ReasonInvalidDigits, "cpf/cnpj")
}
