package validators

import "github.com/jictyvoo/brelem/utils"

func DetermineCPFCNPJ(element string) (ElementType, error) {
	if element == "" {
		return TypeUNKNOWN, utils.ErrEmptyString
	}

	var (
		cpfVld  = newCpfValidator()
		cnpjVld = newCnpjValidator()
	)

	for _, value := range element {
		if value >= '0' && value <= '9' {
			cpfVld.iterateRune(value)
			cnpjVld.iterateRune(value)
		}
	}

	// Check if the result channels close or has an error
	cpfErr, cnpjErr := cpfVld.finishValidation(), cnpjVld.finishValidation()
	if cpfErr != nil && cnpjErr != nil {
		return TypeUNKNOWN, ErrValidateCPFCNPJ
	}

	if cpfErr == nil {
		return TypeCPF, nil
	}
	return TypeCNPJ, nil
}

func AsyncDetermineCPFCNPJ(element string) (ElementType, error) {
	if element == "" {
		return TypeUNKNOWN, utils.ErrEmptyString
	}

	var (
		cpfChan  = make(chan rune, LengthCPF)
		cnpjChan = make(chan rune, LengthCNPJ)
	)

	cpfResult := AsyncCPF(cpfChan)
	cnpjResult := AsyncCNPJ(cnpjChan)

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

func AsyncCPFCNPJ(element string) error {
	_, err := AsyncDetermineCPFCNPJ(element)
	return err
}
