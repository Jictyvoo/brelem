package validators

import (
	"errors"
	"github.com/jictyvoo/brelem/utils"
)

func ValidateDetermineCPFCNPJ(element string) (ElementType, error) {
	if element == "" {
		return TypeUNKNOWN, utils.ErrEmptyString
	}

	var (
		cpfChan  = make(chan rune, LengthCPF)
		cnpjChan = make(chan rune, LengthCNPJ)
	)

	cpfResult := validateAsyncCPF(cpfChan)
	cnpjResult := validateAsyncCNPJ(cnpjChan)

	for _, value := range element {
		if value >= '0' && value <= '9' {
			cpfChan <- value
			cnpjChan <- value
		}
	}
	{
		close(cpfChan)
		close(cnpjChan)
	}

	// Check if the result channels close or has an error

	cpfErr, cnpjErr := <-cpfResult, <-cnpjResult
	if cpfErr != nil && cnpjErr != nil {
		return TypeUNKNOWN, errors.New("given value isn't a valid CPF/CNPJ")
	}

	if cpfErr == nil {
		return TypeCPF, nil
	}
	return TypeCNPJ, nil
}

func CPFCNPJ(element string) error {
	_, err := ValidateDetermineCPFCNPJ(element)
	return err
}
