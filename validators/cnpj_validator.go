package validators

import "github.com/jictyvoo/brelem/utils"

// LengthCNPJ The complete cnpj length (only numbers)
const LengthCNPJ = 14

// / Check a cnpj String and returns a map containing the validation results
func validateAsyncCNPJ(cnpjString chan rune) (result chan error) {
	result = make(chan error, 1)
	go ValidateCNPJ(cnpjString, result)

	return
}

func ValidateCNPJ(cnpjString chan rune, result chan error) {
	defer close(result)
	var (
		verifierDigits      = [...]rune{0, 0}
		originalVerifier    = make([]rune, 0, 2)
		subscriptionWeights = [...]rune{5, 6}
		branchWeights       = [...]rune{9, 9}
		repeatedNumber      = [...]rune{-1, 0}
		iterations          = 0
	)

	for intChar := range cnpjString {
		if intChar >= '0' && intChar <= '9' {
			if iterations < LengthCNPJ-2 {
				if repeatedNumber[0] == -1 {
					// getting first number
					repeatedNumber[0] = intChar
				}

				if repeatedNumber[0] == intChar {
					// counting number repetitions
					repeatedNumber[1]++
				}
				if subscriptionWeights[0] >= 2 {
					verifierDigits[0] += (intChar - '0') * subscriptionWeights[0]
					subscriptionWeights[0]--
				} else {
					verifierDigits[0] += (intChar - '0') * branchWeights[0]
					branchWeights[0]--
				}
				if subscriptionWeights[1] >= 2 {
					verifierDigits[1] += (intChar - '0') * subscriptionWeights[1]
					subscriptionWeights[1]--
				} else {
					verifierDigits[1] += (intChar - '0') * branchWeights[1]
					branchWeights[1]--
				}
			} else {
				originalVerifier = append(originalVerifier, intChar-'0')
			}
			iterations++
		}
	}
	verifierDigits[0] = _modVerifierDigit(verifierDigits[0])
	verifierDigits[1] += verifierDigits[0] * branchWeights[1]
	verifierDigits[1] = _modVerifierDigit(verifierDigits[1])

	if repeatedNumber[1] >= LengthCNPJ-2 || iterations > LengthCNPJ || len(originalVerifier) != 2 {
		result <- utils.ErrElementIncorrectLength
	} else if originalVerifier[0] != verifierDigits[0] || originalVerifier[1] != verifierDigits[1] {
		result <- utils.ErrInvalidElement
	}
}
