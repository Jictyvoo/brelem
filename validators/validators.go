package validators

import "github.com/jictyvoo/brelem/utils"

func DetermineCPFCNPJ(element string) (ElementType, error) {
	if element == "" {
		return TypeUNKNOWN, utils.ErrEmptyString
	}

	// Check if the result channels close or has an error
	cpfErr, cnpjErr := CPF(element), CNPJ(element)
	if cpfErr != nil && cnpjErr != nil {
		return TypeUNKNOWN, ErrValidateCPFCNPJ
	}

	if cpfErr == nil {
		return TypeCPF, nil
	}
	return TypeCNPJ, nil
}

func CPFCNPJ(element string) error {
	_, err := DetermineCPFCNPJ(element)
	return err
}
