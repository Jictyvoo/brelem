package validators

import "errors"

type ElementType uint8

const (
	TypeUNKNOWN ElementType = iota
	TypeCPF
	TypeCNPJ
)

var (
	ErrValidateCPFCNPJ = errors.New("given value isn't a valid CPF/CNPJ")
)

func _modVerifierDigit(toVerify rune) rune {
	toVerify = 11 - (toVerify % 11)
	if toVerify >= 10 {
		toVerify = 0
	}
	return toVerify
}
